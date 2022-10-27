package agent

import (
	"fmt"
	"net"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	pkgGrpc "github.com/rezaAmiri123/microservice/pkg/grpc"
	financeGrpc "github.com/rezaAmiri123/microservice/service_finance/internal/port/grpc"
	financeService "github.com/rezaAmiri123/microservice/service_finance/proto/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

const (
	maxConnectionIdle = 5
	gRPCTimeout       = 15
	maxConnectionAge  = 5
	gRPCTime          = 10
)

func (a *Agent) setupGrpcServer() error {
	im := pkgGrpc.NewInterceptorManager(a.logger)
	var opts []grpc.ServerOption
	opts = append(opts,
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: maxConnectionIdle * time.Minute,
			Timeout:           gRPCTimeout * time.Second,
			MaxConnectionAge:  maxConnectionAge * time.Minute,
			Time:              gRPCTime * time.Minute,
		}),
		// grpc.StreamInterceptor(
		// grpc_middleware.ChainStreamServer(
		// 	grpc_auth.StreamServerInterceptor(auth.Authenticate),
		// )),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
			// grpc_auth.UnaryServerInterceptor(auth.Authenticate),
			im.Payload,
		)),
	)
	serverConfig := &financeGrpc.Config{App: a.Application, Metric: a.metric, Logger: a.logger}
	server, _ := financeGrpc.NewFinanceGRPCServer(serverConfig)
	grpcServer := grpc.NewServer(opts...)
	financeService.RegisterFinanceServiceServer(grpcServer, server)
	reflection.Register(grpcServer)
	grpc_prometheus.Register(grpcServer)
	grpcAddress := fmt.Sprintf("%s:%d", a.Config.GRPCServerAddr, a.Config.GRPCServerPort)
	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return err
	}
	a.grpcServer = grpcServer
	go func() {
		err = a.grpcServer.Serve(listener)
		if err != nil {
			_ = a.Shutdown()
		}
	}()
	return err
}
