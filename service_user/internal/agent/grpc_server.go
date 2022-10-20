package agent

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	pkgTLS "github.com/rezaAmiri123/microservice/pkg/auth/tls"
	userGrpc "github.com/rezaAmiri123/microservice/service_user/internal/port/grpc"
	userService "github.com/rezaAmiri123/microservice/service_user/proto/grpc"

	// grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

const (
	maxConnectionIdle = 5
	gRPCTimeout       = 15
	maxConnectionAge  = 5
	gRPCTime          = 10
)

// Server serves gRPC requests for our banking service.
// type Server struct {
// 	userService.UnimplementedUserServiceServer
// }

func (a *Agent) setupGrpcServer() error {

	// server := &Server{}
	// grpcServer := grpc.NewServer()
	// userService.RegisterUserServiceServer(grpcServer, server)
	// reflection.Register(grpcServer)
	// listener, err := net.Listen("tcp", ":5080")
	// if err != nil {
	// 	log.Fatal("cannot create listener:", err)
	// }
	// err = grpcServer.Serve(listener)
	// if err != nil {
	// 	log.Fatal("cannot start gRPC server:", err)
	// }
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
			// im.Logger,
		)),
	)

	if a.GRPCServerTLSCertFile != "" && a.GRPCServerTLSKeyFile != "" {
		tlsConfig := pkgTLS.TLSConfig{
			CAFile:         a.GRPCServerTLSCAFile,
			CertFile:       a.GRPCServerTLSCertFile,
			KeyFile:        a.GRPCServerTLSKeyFile,
			ServerAddress:  a.GRPCServerTLSServerAddress,
			Server:         true,
			ClientAuthType: tls.RequireAndVerifyClientCert,
		}
		t, err := pkgTLS.SetupTLSConfig(tlsConfig)
		if err != nil {
			return err
		}
		creds := credentials.NewTLS(t)
		opts = append(opts, grpc.Creds(creds))
	}

	serverConfig := &userGrpc.Config{App: a.Application, Metric: a.metric, Logger: a.logger, Maker: a.Maker}
	server, _ := userGrpc.NewUserGRPCServer(serverConfig)
	grpcServer := grpc.NewServer(opts...)
	userService.RegisterUserServiceServer(grpcServer, server)
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
