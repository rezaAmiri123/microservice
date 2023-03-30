package agent

import (
	"fmt"
	"net"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/rezaAmiri123/microservice/depot/internal/constants"
	depotGrpc "github.com/rezaAmiri123/microservice/depot/internal/ports/grpc"
	pkgTLS "github.com/rezaAmiri123/microservice/pkg/auth/tls"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"google.golang.org/grpc/reflection"
	// grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
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

	if a.GRPCServerTLSEnabled {
		tlsConfig := pkgTLS.TLSConfig{
			CAFile:         a.GRPCServerTLSCAFile,
			CertFile:       a.GRPCServerTLSCertFile,
			KeyFile:        a.GRPCServerTLSKeyFile,
			ServerAddress:  a.GRPCServerTLSServerAddress,
			Server:         true,
			ClientAuthType: pkgTLS.GetClientAuthType(a.GRPCServerTLSClientAuthType),
			// ClientAuthType: tls.NoClientCert,
		}
		t, err := pkgTLS.SetupTLSConfig(tlsConfig)
		if err != nil {
			return err
		}
		creds := credentials.NewTLS(t)
		opts = append(opts, grpc.Creds(creds))
	}
	//application := a.container.Get(constants.ApplicationKey).(*app.Application)
	//serverConfig := &userGrpc.Config{
	//	App:    application,
	//	Logger: a.logger,
	//}
	grpcServer := grpc.NewServer(opts...)
	reflection.Register(grpcServer)
	grpc_prometheus.Register(grpcServer)
	_ = depotGrpc.RegisterServerTx(a.container, grpcServer)
	//grpcServer := grpc.NewServer(opts...)
	//userService.RegisterUserServiceServer(grpcServer, server)
	//reflection.Register(grpcServer)

	grpcAddress := fmt.Sprintf("%s:%d", a.Config.GRPCServerAddr, a.Config.GRPCServerPort)
	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return err
	}
	//a.grpcServer = grpcServer
	a.container.AddSingleton(constants.GrpcServerKey, func(c di.Container) (any, error) {
		return grpcServer, nil
	})
	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			_ = a.Shutdown()
		}
	}()
	return err
}
