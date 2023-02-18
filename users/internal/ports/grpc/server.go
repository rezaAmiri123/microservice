package grpc

import (
	"github.com/rezaAmiri123/microservice/users/userspb"
	"google.golang.org/grpc"
)

type (
	Config struct {
	}
	server struct {
		cfg *Config
		userspb.UnimplementedUserServiceServer
	}
)

func NewGrpcServer(cfg *Config, opts ...grpc.ServerOption) (*grpc.Server, error) {
	srv := server{cfg: cfg}
	gsrv := grpc.NewServer(opts...)
	userspb.RegisterUserServiceServer(gsrv, srv)
	return gsrv, nil
}
