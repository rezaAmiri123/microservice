package agent

import (
	"crypto/tls"
	"fmt"

	pkgTLS "github.com/rezaAmiri123/microservice/pkg/auth/tls"
	pkgGrpc "github.com/rezaAmiri123/microservice/pkg/grpc"
	"github.com/rezaAmiri123/microservice/service_api/internal/app"
	"github.com/rezaAmiri123/microservice/service_api/internal/app/commands"
	financerpc "github.com/rezaAmiri123/microservice/service_finance/proto/grpc"
	userrpc "github.com/rezaAmiri123/microservice/service_user/proto/grpc"
)

func (a *Agent) setupApplication() error {
	var clientTLSConfig *tls.Config
	if a.GRPCUserClientTLSEnabled {
		tlsConfig := pkgTLS.TLSConfig{
			CAFile:        a.GRPCUserClientTLSCAFile,
			CertFile:      a.GRPCUserClientTLSCertFile,
			KeyFile:       a.GRPCUserClientTLSKeyFile,
			ServerAddress: a.GRPCUserClientTLSServerAddress,
			Server:        false,
		}
		t, err := pkgTLS.SetupTLSConfig(tlsConfig)
		if err != nil {
			return err
		}
		clientTLSConfig = t
	}
	addr := fmt.Sprintf("%s:%d", a.GRPCUserClientAddr, a.GRPCUserClientPort)
	userConn, err := pkgGrpc.NewGrpcClient(addr, clientTLSConfig, a.logger)
	if err != nil {
		return err
	}
	userClient := userrpc.NewUserServiceClient(userConn)

	addr = fmt.Sprintf("%s:%d", a.GRPCFinanceClientAddr, a.GRPCFinanceClientPort)
	financeConn, err := pkgGrpc.NewGrpcClient(addr, a.GRPCFinanceClientTLSConfig, a.logger)
	if err != nil {
		return err
	}
	financeClient := financerpc.NewFinanceServiceClient(financeConn)

	application := &app.Application{
		Commands: app.Commands{
			// User RPC
			CreateUser:  commands.NewCreateUserHandler(userClient, a.logger),
			Login:       commands.NewLoginHandler(userClient, a.logger),
			LoginVerify: commands.NewLoginVerifyHandler(userClient, a.logger),
			// Finance RPC
			CreateAccount:  commands.NewCreateAccountHandler(financeClient, a.logger),
			CreateTransfer: commands.NewCreateTransferHandler(financeClient, a.logger),
		},
		Queries: app.Queries{},
	}

	a.Application = application
	return nil
}
