package agent

import (
	"database/sql"
	"fmt"
	"github.com/rezaAmiri123/microservice/pkg/db/postgresotel"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/search/internal/adapters/grpc"
	mng "github.com/rezaAmiri123/microservice/search/internal/adapters/mongo"
	"github.com/rezaAmiri123/microservice/search/internal/adapters/pg"
	"github.com/rezaAmiri123/microservice/search/internal/app"
	"github.com/rezaAmiri123/microservice/search/internal/constants"
	"github.com/rezaAmiri123/microservice/search/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

func (a *Agent) setupApplication() error {

	//a.container.AddScoped(constants.InboxStoreKey, func(c di.Container) (any, error) {
	//	db := postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB))
	//	return postgres.NewInboxStore(constants.InboxTableName, db), nil
	//})

	a.container.AddScoped(constants.UsersRepoKey, func(c di.Container) (any, error) {
		grpcAddress := fmt.Sprintf("%s:%d", a.Config.GRPCUserClientAddr, a.Config.GRPCUserClientPort)
		return pg.NewUserCacheRepository(
			constants.UsersCacheTableName,
			postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
			grpc.NewUserRepository(grpcAddress, c.Get(constants.LoggerKey).(logger.Logger)),
		), nil
	})

	a.container.AddScoped(constants.OrdersRepoKey, func(c di.Container) (any, error) {
		return mng.NewOrderRepository(
			constants.ServiceName,
			constants.OrdersTableName,
			//postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
			c.Get(constants.MongoDBKey).(*mongo.Client),
		), nil
	})
	a.container.AddScoped(constants.ProductsRepoKey, func(c di.Container) (any, error) {
		grpcAddress := fmt.Sprintf("%s:%d", a.Config.GRPCStoreClientAddr, a.Config.GRPCStoreClientPort)
		return pg.NewProductCacheRepository(
			constants.ProductsCacheTableName,
			postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
			grpc.NewProductRepository(grpcAddress, c.Get(constants.LoggerKey).(logger.Logger)),
		), nil
	})

	a.container.AddScoped(constants.StoresRepoKey, func(c di.Container) (any, error) {
		grpcAddress := fmt.Sprintf("%s:%d", a.Config.GRPCStoreClientAddr, a.Config.GRPCStoreClientPort)
		return pg.NewStoreCacheRepository(
			constants.StoresCacheTableName,
			postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
			grpc.NewStoreRepository(grpcAddress, c.Get(constants.LoggerKey).(logger.Logger)),
		), nil
	})

	a.container.AddScoped(constants.ApplicationKey, func(c di.Container) (any, error) {
		return app.New(
			c.Get(constants.OrdersRepoKey).(domain.OrderRepository),
		), nil
	})
	return nil
}
