package agent

import (
	"database/sql"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/pkg/db/postgresotel"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/es"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/stores/internal/adapters/pg"
	"github.com/rezaAmiri123/microservice/stores/internal/app"
	"github.com/rezaAmiri123/microservice/stores/internal/constants"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
)

func (a *Agent) setupApplication() error {
	a.container.AddScoped(constants.AggregateStoreKey, func(c di.Container) (any, error) {
		tx := postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx))
		reg := c.Get(constants.RegistryKey).(registry.Registry)
		return es.AggregateStoreWithMiddleware(
			postgres.NewEventStore(constants.EventsTableName, tx, reg),
			postgres.NewSnapshotStore(constants.SnapshotsTableName, tx, reg),
		), nil
	})

	a.container.AddScoped(constants.StoresRepoKey, func(c di.Container) (any, error) {
		return es.NewAggregateRepository[*domain.Store](
			domain.StoreAggregate,
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.AggregateStoreKey).(es.AggregateStore),
		), nil
	})

	a.container.AddScoped(constants.ProductsRepoKey, func(c di.Container) (any, error) {
		return es.NewAggregateRepository[*domain.Product](
			domain.ProductAggregate,
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.AggregateStoreKey).(es.AggregateStore),
		), nil
	})
	a.container.AddScoped(constants.CatalogRepoKey, func(c di.Container) (any, error) {
		return pg.NewCatalogRepository(
			constants.CatalogTableName,
			postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
		), nil
	})

	a.container.AddScoped(constants.MallRepoKey, func(c di.Container) (any, error) {
		return pg.NewMallRepository(
			constants.MallTableName,
			postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
		), nil
	})

	a.container.AddScoped(constants.ApplicationKey, func(c di.Container) (any, error) {
		publisher := c.Get(constants.DomainDispatcherKey).(ddd.EventPublisher[ddd.Event])
		stores := c.Get(constants.StoresRepoKey).(domain.StoreRepository)
		products := c.Get(constants.ProductsRepoKey).(domain.ProductRepository)
		catalog := c.Get(constants.CatalogRepoKey).(domain.CatalogRepository)
		malls := c.Get(constants.MallRepoKey).(domain.MallRepository)
		log := c.Get(constants.LoggerKey).(logger.Logger)

		application := app.New(stores, products, catalog, malls, publisher, log)

		return application, nil
	})

	return nil
}
