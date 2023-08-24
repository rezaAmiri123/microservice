package raft

import (
	"bytes"
	"crypto/tls"
	"github.com/hashicorp/raft"
	"github.com/rezaAmiri123/microservice/proglog/internal/constats"
	"github.com/stackus/errors"
	"net"
	"time"
)

var ErrNotARaftRPC = errors.Wrap(errors.ErrBadRequest, "is not a raft request")

var _ raft.StreamLayer = (*StreamLayer)(nil)

type StreamLayer struct {
	ln              net.Listener
	serverTLSConfig *tls.Config
	peerTLSConfig   *tls.Config
}

func NewStreamLayer(ln net.Listener, serverTLSConfig, peerTLSConfig *tls.Config) *StreamLayer {
	return &StreamLayer{
		ln:              ln,
		serverTLSConfig: serverTLSConfig,
		peerTLSConfig:   peerTLSConfig,
	}
}

func (s *StreamLayer) Accept() (net.Conn, error) {
	conn, err := s.ln.Accept()
	if err != nil {
		return nil, err
	}
	data := make([]byte, 1)
	_, err = conn.Read(data)
	if err != nil {
		return nil, err
	}
	if bytes.Compare([]byte{byte(constats.RaftRPC)}, data) != 0 {
		return nil, ErrNotARaftRPC
	}
	if s.serverTLSConfig != nil {
		return tls.Server(conn, s.serverTLSConfig), nil
	}
	return conn, nil
}
func (s *StreamLayer) Close() error {
	return s.ln.Close()
}
func (s *StreamLayer) Addr() net.Addr {
	return s.ln.Addr()
}
func (s *StreamLayer) Dial(address raft.ServerAddress, timeout time.Duration) (net.Conn, error) {
	dialer := &net.Dialer{Timeout: timeout}
	var conn, err = dialer.Dial("tcp", string(address))
	if err != nil {
		return nil, err
	}

	// identify to mux this is a raft rpc
	_, err = conn.Write([]byte{byte(constats.RaftRPC)})
	if err != nil {
		return nil, err
	}
	if s.peerTLSConfig != nil {
		conn = tls.Client(conn, s.peerTLSConfig)
	}
	return conn, nil
}
