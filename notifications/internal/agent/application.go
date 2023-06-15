package agent

import (
	"fmt"
	"github.com/rezaAmiri123/microservice/notifications/internal/adapters/grpc"
	mgo "github.com/rezaAmiri123/microservice/notifications/internal/adapters/mongo"
	"github.com/rezaAmiri123/microservice/notifications/internal/app"
	"github.com/rezaAmiri123/microservice/notifications/internal/constants"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func (a *Agent) setupApplication() error {
	a.container.AddScoped(constants.UsersRepoKey, func(c di.Container) (any, error) {
		grpcAddress := fmt.Sprintf("%s:%d", a.Config.GRPCUserClientAddr, a.Config.GRPCUserClientPort)
		return mgo.NewUserCacheRepository(
			constants.ServiceName,
			constants.UsersCacheTableName,
			c.Get(constants.DatabaseKey).(*mongo.Client),
			grpc.NewUserRepository(grpcAddress, c.Get(constants.LoggerKey).(logger.Logger)),
		), nil
	})

	a.container.AddScoped(constants.ApplicationKey, func(c di.Container) (any, error) {
		return app.New(c.Get(constants.UsersRepoKey).(app.UserCacheRepository)), nil
	})
	return nil
}
