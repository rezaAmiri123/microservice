package raft

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	"github.com/rezaAmiri123/microservice/proglog/internal/constats"
	"github.com/rezaAmiri123/microservice/proglog/internal/discovery"
	"github.com/rezaAmiri123/microservice/proglog/internal/distribution"
	"github.com/rezaAmiri123/microservice/proglog/internal/domain"
	"google.golang.org/protobuf/proto"
	"os"
	"path/filepath"
	"time"
)

var (
	SnapshotRetrain = 1
	MaxPool         = 5
	Timeout         = 10 * time.Second
)

var _ interface {
	discovery.Handler
	distribution.GetServers
} = (*DistributedLog)(nil)

type (
	Config struct {
		raft.Config
		BindAddr    string
		StreamLayer *StreamLayer
		Bootstrap   bool
	}
	DistributedLog struct {
		config    Config
		logConfig domain.Config
		log       domain.Log
		raft      *raft.Raft
	}
)

func NewDistributedLog(dataDir string, config Config, logConfig domain.Config) (*DistributedLog, error) {
	l := &DistributedLog{
		config:    config,
		logConfig: logConfig,
	}
	if err := l.setupLog(dataDir); err != nil {
		return nil, err
	}
	if err := l.setupRaft(dataDir); err != nil {
		return nil, err
	}
	return l, nil
}

func (l *DistributedLog) setupLog(dataDir string) error {
	logDir := filepath.Join(dataDir, "log")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	var err error
	l.log, err = domain.NewLog(logDir, l.logConfig)
	return err
}

func (l *DistributedLog) setupRaft(dataDir string) error {
	fsm := &fsm{log: l.log}

	logDir := filepath.Join(dataDir, "raft", "log")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}
	logConfig := l.logConfig
	logConfig.Segment.InitialOffset = 1
	logStore, err := newLogStore(logDir, logConfig)
	if err != nil {
		return err
	}

	stableStore, err := raftboltdb.NewBoltStore(
		filepath.Join(dataDir, "raft", "stable"),
	)
	if err != nil {
		return err
	}

	snapshotStore, err := raft.NewFileSnapshotStore(
		filepath.Join(dataDir, "raft"),
		SnapshotRetrain,
		os.Stderr,
	)
	if err != nil {
		return err
	}

	transport := raft.NewNetworkTransport(
		l.config.StreamLayer,
		MaxPool,
		Timeout,
		os.Stderr,
	)

	config := raft.DefaultConfig()
	config.LocalID = l.config.LocalID
	if l.config.HeartbeatTimeout != 0 {
		config.HeartbeatTimeout = l.config.HeartbeatTimeout
	}
	if l.config.ElectionTimeout != 0 {
		config.ElectionTimeout = l.config.ElectionTimeout
	}
	if l.config.LeaderLeaseTimeout != 0 {
		config.LeaderLeaseTimeout = l.config.LeaderLeaseTimeout
	}
	if l.config.CommitTimeout != 0 {
		config.CommitTimeout = l.config.CommitTimeout
	}

	l.raft, err = raft.NewRaft(
		config,
		fsm,
		logStore,
		stableStore,
		snapshotStore,
		transport,
	)
	if err != nil {
		return err
	}

	if l.config.Bootstrap {
		cfg := raft.Configuration{
			Servers: []raft.Server{{
				ID:      config.LocalID,
				Address: raft.ServerAddress(l.config.BindAddr),
			}},
		}
		err = l.raft.BootstrapCluster(cfg).Error()
	}
	return err
}

func (l *DistributedLog) Join(id, addr string) error {
	configFuture := l.raft.GetConfiguration()
	if err := configFuture.Error(); err != nil {
		return err
	}
	serverID := raft.ServerID(id)
	serverAddr := raft.ServerAddress(addr)
	for _, srv := range configFuture.Configuration().Servers {
		if srv.ID == serverID || srv.Address == serverAddr {
			// server has already joined
			return nil
		}
		// remove the existing server
		removeFuture := l.raft.RemoveServer(serverID, 0, 0)
		if err := removeFuture.Error(); err != nil {
			return err
		}
	}
	addFuture := l.raft.AddVoter(serverID, serverAddr, 0, 0)
	if err := addFuture.Error(); err != nil {
		return err
	}
	return nil
}
func (l *DistributedLog) Leave(id string) error {
	removeFuture := l.raft.RemoveServer(raft.ServerID(id), 0, 0)
	return removeFuture.Error()
}
func (l *DistributedLog) Append(record *domain.Record) (uint64, error) {
	res, err := l.apply(constats.AppendRequestType, record)
	if err != nil {
		return 0, err
	}
	return res.(*domain.Record).Offset, nil
}

func (l *DistributedLog) apply(reqType constats.RequestType, req proto.Message) (any, error) {
	var buf bytes.Buffer
	if _, err := buf.Write([]byte{byte(reqType)}); err != nil {
		return nil, err
	}
	data, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(data)
	if err != nil {
		return nil, err
	}

	future := l.raft.Apply(buf.Bytes(), Timeout)
	if future.Error() != nil {
		return nil, future.Error()
	}
	res := future.Response()
	if err, ok := res.(error); ok {
		return nil, err
	}
	return res, nil

}
func (l *DistributedLog) Read(offset uint64) (*domain.Record, error) {
	return l.log.Read(offset)
}

func (l *DistributedLog) WaitForLeader(timeout time.Duration) error {
	timeoutc := time.After(timeout)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeoutc:
			return fmt.Errorf("timed out")
		case <-ticker.C:
			if l, _ := l.raft.LeaderWithID(); l != "" {
				return nil
			}
		}
	}
}

func (l *DistributedLog) GetServers() ([]*distribution.Server, error) {
	future := l.raft.GetConfiguration()
	if err := future.Error(); err != nil {
		return nil, err
	}
	var servers []*distribution.Server
	for _, server := range future.Configuration().Servers {
		serverAddr, _ := l.raft.LeaderWithID()
		servers = append(servers, &distribution.Server{
			Id:       string(server.ID),
			RpcAddr:  string(server.Address),
			IsLeader: serverAddr == server.Address,
		})
	}
	return servers, nil
}
