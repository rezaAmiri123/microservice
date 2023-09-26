package agent

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/hashicorp/raft"
	"github.com/rezaAmiri123/microservice/proglog/internal/constants"
	pkgSerf "github.com/rezaAmiri123/microservice/proglog/internal/discovery/serf"
	pkgRaft "github.com/rezaAmiri123/microservice/proglog/internal/distribution/raft"
	"github.com/rezaAmiri123/microservice/proglog/internal/domain"
	pkgGrpc "github.com/rezaAmiri123/microservice/proglog/internal/ports/grpc"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"net"
	"sync"
	"time"
)

//type Config struct {
//	ServerTLSConfig *tls.Config
//	PeerTLSConfig   *tls.Config
//	// DataDir stores the log and raft data.
//	DataDir string
//	// BindAddr is the address serf runs on.
//	BindAddr string
//	// RPCPort is the port for client (and Raft) connections.
//	RPCPort int
//	// Raft server id.
//	NodeName string
//	// Bootstrap should be set to true when starting the first node of the cluster.
//	StartJoinAddrs []string
//	ACLModelFile   string
//	ACLPolicyFile  string
//	// START: config
//	Bootstrap bool
//	// END: config
//}
//
//func (c Config) RPCAddr() (string, error) {
//	host, _, err := net.SplitHostPort(c.BindAddr)
//	if err != nil {
//		return "", err
//	}
//	return fmt.Sprintf("%s:%d", host, c.RPCPort), nil
//}
//
//type Agent struct {
//	Config Config
//
//	mux        cmux.CMux
//	DLog       *pkgRaft.DistributedLog
//	server     *grpc.Server
//	membership *pkgSerf.MemberShip
//	logger     *zap.Logger
//
//	shutdown     bool
//	shutdowns    chan struct{}
//	shutdownLock sync.Mutex
//}
//
//func New(config Config) (*Agent, error) {
//	logger, err := zap.NewDevelopment()
//	if err != nil {
//		return nil, err
//	}
//	zap.ReplaceGlobals(logger)
//
//	a := &Agent{
//		Config:    config,
//		shutdowns: make(chan struct{}),
//		logger:    logger.Named("agent"),
//	}
//
//	setup := []func() error{
//		a.setupMux,
//		a.setupLog,
//		a.setupServer,
//		a.setupMembership,
//	}
//	for _, fn := range setup {
//		if err := fn(); err != nil {
//			return nil, err
//		}
//	}
//	go a.serve()
//
//	return a, nil
//}
//func (a *Agent) setupMux() error {
//	rpcAddr := fmt.Sprintf(":%d", a.Config.RPCPort)
//	ln, err := net.Listen("tcp", rpcAddr)
//	if err != nil {
//		return err
//	}
//	a.mux = cmux.New(ln)
//	a.logger.Info("mux is setted up")
//	return nil
//}
//
//func (a *Agent) setupLog() error {
//	raftLn := a.mux.Match(func(reader io.Reader) bool {
//		first := make([]byte, 1)
//		if _, err := reader.Read(first); err != nil {
//			return false
//		}
//		return bytes.Compare(first, []byte{byte(constants.RaftRPC)}) == 0
//	})
//
//	raftConfig := pkgRaft.Config{}
//	raftConfig.StreamLayer = pkgRaft.NewStreamLayer(raftLn, nil, nil)
//	rpcAddr, err := a.Config.RPCAddr()
//	if err != nil {
//		return err
//	}
//	raftConfig.BindAddr = rpcAddr
//	raftConfig.LocalID = raft.ServerID(a.Config.NodeName)
//	raftConfig.Bootstrap = a.Config.Bootstrap
//
//	a.DLog, err = pkgRaft.NewDistributedLog(a.Config.DataDir, raftConfig, domain.Config{})
//	if err != nil {
//		return err
//	}
//
//	if a.Config.Bootstrap {
//		if err := a.DLog.WaitForLeader(3 * time.Second); err != nil {
//			return err
//		}
//	}
//	a.logger.Info("log is setted up")
//	return nil
//}
//
//func (a *Agent) setupServer() error {
//	serverConfig := &pkgGrpc.Config{
//		Log:         a.DLog,
//		GetServerer: a.DLog,
//	}
//
//	var opts []grpc.ServerOption
//	var err error
//	a.server, err = pkgGrpc.NewGRPCServer(serverConfig, opts...)
//	if err != nil {
//		return err
//	}
//
//	grpcLn := a.mux.Match(cmux.Any())
//	go func() {
//		if err := a.server.Serve(grpcLn); err != nil {
//			_ = a.Shutdown()
//		}
//	}()
//	if err != nil {
//		return err
//	}
//	a.logger.Info("grpc server is setted up")
//	return nil
//}
//func (a *Agent) setupMembership() error {
//	rpcAddr, err := a.Config.RPCAddr()
//	if err != nil {
//		return err
//	}
//	a.membership, err = pkgSerf.New(a.DLog, pkgSerf.Config{
//		NodeName: a.Config.NodeName,
//		BindAddr: a.Config.BindAddr,
//		Tags: map[string]string{
//			constants.EventRPCAddress: rpcAddr,
//		},
//		StartJoinAddrs: a.Config.StartJoinAddrs,
//	})
//	if err != nil {
//		return err
//	}
//	a.logger.Info("membership is setted up")
//	return nil
//}
//
//func (a *Agent) serve() error {
//	if err := a.mux.Serve(); err != nil {
//		_ = a.Shutdown()
//		return err
//	}
//	return nil
//}
//
//func (a *Agent) Shutdown() error {
//	a.shutdownLock.Lock()
//	defer a.shutdownLock.Unlock()
//	if a.shutdown {
//		return nil
//	}
//	a.shutdown = true
//	close(a.shutdowns)
//
//	shutdown := []func() error{
//		a.membership.Leave,
//		func() error {
//			a.server.GracefulStop()
//			return nil
//		},
//		a.DLog.Close,
//	}
//	for _, fn := range shutdown {
//		if err := fn(); err != nil {
//			return err
//		}
//	}
//	a.logger.Info("shutdown is finished")
//	return nil
//}
//

// ###############################################
type Config struct {
	ServerTLSConfig *tls.Config
	PeerTLSConfig   *tls.Config
	// DataDir stores the log and raft data.
	DataDir string
	// BindAddr is the address serf runs on.
	BindAddr string
	// RPCPort is the port for client (and Raft) connections.
	RPCPort int
	// Raft server id.
	NodeName string
	// Bootstrap should be set to true when starting the first node of the cluster.
	StartJoinAddrs []string
	ACLModelFile   string
	ACLPolicyFile  string
	// START: config
	Bootstrap bool
	// END: config
}

func (c Config) RPCAddr() (string, error) {
	host, _, err := net.SplitHostPort(c.BindAddr)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%d", host, c.RPCPort), nil
}

// START: agent
type Agent struct {
	Config Config

	mux        cmux.CMux
	log        *pkgRaft.DistributedLog
	server     *grpc.Server
	membership *pkgSerf.MemberShip

	shutdown     bool
	shutdowns    chan struct{}
	shutdownLock sync.Mutex
}

// END: agent

func New(config Config) (*Agent, error) {
	a := &Agent{
		Config:    config,
		shutdowns: make(chan struct{}),
	}
	// START: add_setup_mux
	setup := []func() error{
		// START_HIGHLIGHT
		a.setupMux,
		// END_HIGHLIGHT
		a.setupLog,
		a.setupServer,
		a.setupMembership,
	}
	// END: add_setup_mux
	for _, fn := range setup {
		if err := fn(); err != nil {
			return nil, err
		}
	}
	// START: new_serve
	go a.serve()
	// END: new_serve
	return a, nil
}

// START: setup_mux
func (a *Agent) setupMux() error {
	rpcAddr := fmt.Sprintf(
		":%d",
		a.Config.RPCPort,
	)
	ln, err := net.Listen("tcp", rpcAddr)
	if err != nil {
		return err
	}
	a.mux = cmux.New(ln)
	return nil
}

// END: setup_mux

// START: setup_log_start
func (a *Agent) setupLog() error {
	raftLn := a.mux.Match(func(reader io.Reader) bool {
		b := make([]byte, 1)
		if _, err := reader.Read(b); err != nil {
			return false
		}
		return bytes.Compare(b, []byte{byte(constants.RaftRPC)}) == 0
	})
	// END: setup_log_start
	// START: setup_log_end
	logConfig := pkgRaft.Config{}
	logConfig.StreamLayer = pkgRaft.NewStreamLayer(
		raftLn,
		a.Config.ServerTLSConfig,
		a.Config.PeerTLSConfig,
	)
	logConfig.LocalID = raft.ServerID(a.Config.NodeName)
	logConfig.Bootstrap = a.Config.Bootstrap
	var err error
	a.log, err = pkgRaft.NewDistributedLog(
		a.Config.DataDir,
		logConfig,
		domain.Config{},
	)
	if err != nil {
		return err
	}
	if a.Config.Bootstrap {
		return a.log.WaitForLeader(3 * time.Second)
	}
	return nil
}

// END: setup_log_end

func (a *Agent) setupServer() error {
	//authorizer := auth.New(
	//	a.Config.ACLModelFile,
	//	a.Config.ACLPolicyFile,
	//)
	serverConfig := &pkgGrpc.Config{
		//CommitLog:   a.log,
		//Authorizer:  authorizer,
		Log:         a.log,
		GetServerer: a.log,
	}
	var opts []grpc.ServerOption
	if a.Config.ServerTLSConfig != nil {
		creds := credentials.NewTLS(a.Config.ServerTLSConfig)
		opts = append(opts, grpc.Creds(creds))
	}
	var err error
	a.server, err = pkgGrpc.NewGRPCServer(serverConfig, opts...)
	if err != nil {
		return err
	}
	// START: setup_server
	grpcLn := a.mux.Match(cmux.Any())
	go func() {
		if err := a.server.Serve(grpcLn); err != nil {
			_ = a.Shutdown()
		}
	}()
	return err
	// END: setup_server
}

// START: setup_membership
func (a *Agent) setupMembership() error {
	rpcAddr, err := a.Config.RPCAddr()
	if err != nil {
		return err
	}
	a.membership, err = pkgSerf.New(a.log, pkgSerf.Config{
		NodeName: a.Config.NodeName,
		BindAddr: a.Config.BindAddr,
		Tags: map[string]string{
			"rpc_addr": rpcAddr,
		},
		StartJoinAddrs: a.Config.StartJoinAddrs,
	})
	return err
}

// END: setup_membership

// START: serve
func (a *Agent) serve() error {
	if err := a.mux.Serve(); err != nil {
		_ = a.Shutdown()
		return err
	}
	return nil
}

// END: serve

func (a *Agent) Shutdown() error {
	a.shutdownLock.Lock()
	defer a.shutdownLock.Unlock()
	if a.shutdown {
		return nil
	}
	a.shutdown = true
	close(a.shutdowns)

	shutdown := []func() error{
		a.membership.Leave,
		func() error {
			a.server.GracefulStop()
			return nil
		},
		a.log.Close,
	}
	for _, fn := range shutdown {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}
