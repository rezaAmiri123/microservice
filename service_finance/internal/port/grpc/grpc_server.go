package grpc

import (
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/service_finance/internal/app"
	"github.com/rezaAmiri123/microservice/service_finance/internal/metrics"
	financeService "github.com/rezaAmiri123/microservice/service_finance/proto/grpc"
)

type Config struct {
	Metric *metrics.FinanceServiceMetric
	App    *app.Application
	Logger logger.Logger
}

var _ financeService.FinanceServiceServer = (*FinanceGRPCServer)(nil)

type FinanceGRPCServer struct {
	cfg *Config
	financeService.UnimplementedFinanceServiceServer
}

func NewFinanceGRPCServer(config *Config) (*FinanceGRPCServer, error) {
	srv := &FinanceGRPCServer{
		cfg: config,
	}
	return srv, nil
}

// func NewGRPCServer(config *Config, opts ...grpc.ServerOption) (*grpc.Server, error) {
// 	gsrv := grpc.NewServer(opts...)
// 	srv, err := NewUserGRPCServer(config)
// 	if err != nil {
// 		return nil, err
// 	}
// 	userService.RegisterUserServiceServer(gsrv, srv)
// 	return gsrv, nil
// }
