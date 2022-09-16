package grpc

import (
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/service_user/internal/app"
	"github.com/rezaAmiri123/microservice/service_user/internal/metrics"
	userService "github.com/rezaAmiri123/microservice/service_user/proto/grpc"
	"google.golang.org/grpc"
)

type Config struct {
	Metric *metrics.UserServiceMetric
	App    *app.Application
	Logger logger.Logger
}

var _ userService.UserServiceServer = (*UserGRPCServer)(nil)

type UserGRPCServer struct {
	cfg *Config
	userService.UnimplementedUserServiceServer
}

func NewUserGRPCServer(config *Config) (*UserGRPCServer, error) {
	srv := &UserGRPCServer{
		cfg: config,
	}
	return srv, nil
}

func NewGRPCServer(config *Config, opts ...grpc.ServerOption) (*grpc.Server, error) {
	gsrv := grpc.NewServer(opts...)
	srv, err := NewUserGRPCServer(config)
	if err != nil {
		return nil, err
	}
	userService.RegisterUserServiceServer(gsrv, srv)
	return gsrv, nil
}
