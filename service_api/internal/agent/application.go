package agent

import (
	"fmt"

	pkgGrpc "github.com/rezaAmiri123/microservice/pkg/grpc"
	"github.com/rezaAmiri123/microservice/service_api/internal/app"
	"github.com/rezaAmiri123/microservice/service_api/internal/app/commands"
	userrpc "github.com/rezaAmiri123/microservice/service_user/proto/grpc"
)

func (a *Agent) setupApplication() error {
	addr := fmt.Sprintf("%s:%d", a.GRPCUserClientAddr, a.GRPCUserClientPort)
	userConn, err := pkgGrpc.NewGrpcClient(addr, a.GRPCUserClientTLSConfig, a.logger)
	if err != nil {
		return err
	}
	userClient := userrpc.NewUserServiceClient(userConn)

	application := &app.Application{
		Commands: app.Commands{
			CreateUser: commands.NewCreateUserHandler(userClient, a.logger),
		},
		Queries: app.Queries{},
	}

	a.Application = application
	return nil
}
