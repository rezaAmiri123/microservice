package agent

import (
	"database/sql"
	"fmt"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/pkg/es"
	"github.com/rezaAmiri123/microservice/pkg/registry"

	"github.com/rezaAmiri123/microservice/baskets/internal/adapters/grpc"
	"github.com/rezaAmiri123/microservice/baskets/internal/adapters/pg"
	"github.com/rezaAmiri123/microservice/baskets/internal/app"
	"github.com/rezaAmiri123/microservice/baskets/internal/constants"
	"github.com/rezaAmiri123/microservice/baskets/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/db/postgresotel"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/logger"
)

func (a *Agent) setupApplication() error {
	//a.container.AddScoped(constants.BasketsRepoKey, func(c di.Container) (any, error) {
	//	return pg.NewBasketRepository(
	//		constants.BasketTableName,
	//		postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB)),
	//	), nil
	//})
	a.container.AddScoped(constants.BasketsRepoKey, func(c di.Container) (any, error) {
		db := postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB))
		reg := c.Get(constants.RegistryKey).(registry.Registry)
		return es.NewAggregateRepository[*domain.Basket](
			domain.BasketAggregate,
			reg,
			es.AggregateStoreWithMiddleware(
				postgres.NewEventStore(constants.EventsTableName, db, reg),
				postgres.NewSnapshotStore(constants.SnapshotsTableName, db, reg),
			),
		), nil
	})
	a.container.AddScoped(constants.StoresRepoKey, func(c di.Container) (any, error) {
		addr := fmt.Sprintf("%s:%d", a.Config.GRPCStoreClientAddr, a.Config.GRPCStoreClientPort)
		return pg.NewStoreCacheRepository(
			constants.StoresCacheTableName,
			postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB)),
			grpc.NewStoreRepository(addr, c.Get(constants.LoggerKey).(logger.Logger)),
		), nil
	})

	a.container.AddScoped(constants.ProductsRepoKey, func(c di.Container) (any, error) {
		addr := fmt.Sprintf("%s:%d", a.Config.GRPCStoreClientAddr, a.Config.GRPCStoreClientPort)
		return pg.NewProductCacheRepository(
			constants.ProductsCacheTableName,
			postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB)),
			grpc.NewProductRepository(addr, c.Get(constants.LoggerKey).(logger.Logger)),
		), nil
	})

	a.container.AddScoped(constants.ApplicationKey, func(c di.Container) (any, error) {
		publisher := c.Get(constants.DomainDispatcherKey).(ddd.EventPublisher[ddd.Event])
		baskets := c.Get(constants.BasketsRepoKey).(domain.BasketRepository)
		stores := c.Get(constants.StoresRepoKey).(domain.StoreCacheRepository)
		products := c.Get(constants.ProductsRepoKey).(domain.ProductCacheRepository)

		log := c.Get(constants.LoggerKey).(logger.Logger)
		application := app.NewInstrumentedApp(
			app.New(baskets, stores, products, publisher, log),
		)
		return application, nil
	})
	return nil
}
