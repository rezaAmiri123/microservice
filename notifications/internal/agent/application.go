package agent

import (
	"database/sql"
	"fmt"
	"github.com/rezaAmiri123/microservice/notifications/internal/adapters/grpc"
	"github.com/rezaAmiri123/microservice/notifications/internal/adapters/pg"
	"github.com/rezaAmiri123/microservice/notifications/internal/app"
	"github.com/rezaAmiri123/microservice/notifications/internal/constants"
	"github.com/rezaAmiri123/microservice/pkg/db/postgresotel"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/logger"
)

func (a *Agent) setupApplication() error {
	a.container.AddScoped(constants.UsersRepoKey, func(c di.Container) (any, error) {
		grpcAddress := fmt.Sprintf("%s:%d", a.Config.GRPCUserClientAddr, a.Config.GRPCUserClientPort)
		return pg.NewUserCacheRepository(
			constants.UsersCacheTableName,
			postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB)),
			grpc.NewUserRepository(grpcAddress, c.Get(constants.LoggerKey).(logger.Logger)),
		), nil
	})

	a.container.AddScoped(constants.ApplicationKey, func(c di.Container) (any, error) {
		return app.New(c.Get(constants.UsersRepoKey).(app.UserCacheRepository)), nil
	})
	return nil
}
