package serf

import (
	"github.com/hashicorp/raft"
	"github.com/hashicorp/serf/serf"
	"github.com/rezaAmiri123/microservice/proglog/internal/constats"
	"github.com/rezaAmiri123/microservice/proglog/internal/discovery"
	"go.uber.org/zap"
	"net"
)

type Config struct {
	NodeName       string
	BindAddr       string
	Tags           map[string]string
	StartJoinAddrs []string
}

type MemberShip struct {
	Config
	handler discovery.Handler
	serf    *serf.Serf
	events  chan serf.Event
	logger  *zap.Logger
}

func New(handler discovery.Handler, config Config) (*MemberShip, error) {
	c := &MemberShip{
		Config:  config,
		handler: handler,
		logger:  zap.L().Named("membership"),
	}
	if err := c.setupSerf(); err != nil {
		return nil, err
	}
	return c, nil
}

func (m *MemberShip) setupSerf() (err error) {
	addr, err := net.ResolveTCPAddr("tcp", m.BindAddr)
	if err != nil {
		return err
	}
	config := serf.DefaultConfig()
	config.Init()
	config.MemberlistConfig.BindAddr = addr.IP.String()
	config.MemberlistConfig.BindPort = addr.Port
	m.events = make(chan serf.Event)
	config.EventCh = m.events
	config.Tags = m.Tags
	config.NodeName = m.Config.NodeName
	m.serf, err = serf.Create(config)
	if err != nil {
		return err
	}

	go m.eventHandler() //<label id="handlergoroutine" />

	if m.StartJoinAddrs != nil {
		_, err = m.serf.Join(m.StartJoinAddrs, true)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MemberShip) eventHandler() {
	for e := range m.events {
		switch e.EventType() {
		case serf.EventMemberJoin:
			for _, member := range e.(serf.MemberEvent).Members {
				if m.isLocal(member) {
					continue
				}
				m.handleJoin(member)
			}
		case serf.EventMemberLeave, serf.EventMemberFailed:
			for _, member := range e.(serf.MemberEvent).Members {
				if m.isLocal(member) {
					return
				}
				m.handleLeave(member)
			}
		}
	}
}

func (m *MemberShip) handleJoin(member serf.Member) {
	err := m.handler.Join(member.Name, member.Tags[constats.EventRPCAddress])
	if err != nil {
		m.logError(err, "failed to join", member)
	}
}

func (m *MemberShip) handleLeave(member serf.Member) {
	err := m.handler.Leave(member.Name)
	if err != nil {
		m.logError(err, "failed to leave", member)
	}
}

func (m *MemberShip) isLocal(member serf.Member) bool {
	return m.serf.LocalMember().Name == member.Name
}

func (m *MemberShip) Members() []serf.Member {
	return m.serf.Members()
}

func (m *MemberShip) Leave() error {
	return m.serf.Leave()
}

func (m *MemberShip) logError(err error, msg string, member serf.Member) {
	log := m.logger.Error
	if err == raft.ErrNotLeader {
		log = m.logger.Debug
	}
	log(
		msg,
		zap.Error(err),
		zap.String("name", member.Name),
		zap.String(constats.EventRPCAddress, member.Tags[constats.EventRPCAddress]),
	)
}
