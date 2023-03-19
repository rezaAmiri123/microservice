package agent

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/es"
	"github.com/rezaAmiri123/microservice/pkg/jetstream"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/tm"
	"github.com/rezaAmiri123/microservice/stores/internal/adapters/pg"
	"github.com/rezaAmiri123/microservice/stores/internal/app"
	"github.com/rezaAmiri123/microservice/stores/internal/app/commands"
	"github.com/rezaAmiri123/microservice/stores/internal/app/queries"
	"github.com/rezaAmiri123/microservice/stores/internal/constants"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
	"github.com/rezaAmiri123/microservice/stores/internal/ports/handlers"
)

func (a *Agent) setupApplication() error {
	dbConn, err := postgres.NewDB(postgres.Config{
		PGDriver:   a.PGDriver,
		PGHost:     a.PGHost,
		PGPort:     a.PGPort,
		PGUser:     a.PGUser,
		PGDBName:   a.PGDBName,
		PGPassword: a.PGPassword,
	})
	if err != nil {
		return fmt.Errorf("cannot load db: %w", err)
	}

	a.container.AddSingleton(constants.DatabaseKey, func(c di.Container) (any, error) {
		return dbConn, nil
	})

	//repo, err := adapters.NewGORMArticleRepository(a.DBConfig)
	//repo := pg.NewPGUserRepository(dbConn, a.logger)

	// setup Driven adapters

	//if err = userspb.Registrations(reg); err != nil {
	//	return err
	//}
	js, err := a.nats()
	if err != nil {
		return err
	}
	//log := a.container.Get(constants.LoggerKey).(logger.Logger)
	stream := jetstream.NewStream(a.NatsStream, js, a.container.Get(constants.LoggerKey).(logger.Logger))
	//a.container.AddSingleton(constants.DomainDispatcherKey, func(c di.Container) (any, error) {
	//	return ddd.NewEventDispatcher[ddd.AggregateEvent](), nil
	//})
	a.container.AddSingleton(constants.DomainDispatcherKey, func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.Event](), nil
	})
	a.container.AddScoped(constants.DatabaseTransactionKey, func(c di.Container) (any, error) {
		return dbConn.Begin()
	})
	a.container.AddScoped(constants.MessagePublisherKey, func(c di.Container) (any, error) {
		tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
		outboxStore := postgres.NewOutboxStore(constants.OutboxTableName, tx)
		return am.NewMessagePublisher(
			stream,
			tm.OutboxPublisher(outboxStore),
		), nil
	})
	a.container.AddScoped(constants.MessageSubscriberKey, func(c di.Container) (any, error) {
		return am.NewMessageSubscriber(stream), nil
	})
	a.container.AddScoped(constants.EventPublisherKey, func(c di.Container) (any, error) {
		return am.NewEventPublisher(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.MessagePublisherKey).(am.MessagePublisher),
		), nil
	})

	a.container.AddScoped(constants.InboxStoreKey, func(c di.Container) (any, error) {
		tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
		return postgres.NewInboxStore(constants.InboxTableName, tx), nil
	})

	a.container.AddScoped(constants.AggregateStoreKey, func(c di.Container) (any, error) {
		tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
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
			c.Get(constants.DatabaseTransactionKey).(*sql.Tx),
		), nil
	})

	a.container.AddScoped(constants.MallRepoKey, func(c di.Container) (any, error) {
		return pg.NewMallRepository(
			constants.MallTableName,
			c.Get(constants.DatabaseTransactionKey).(*sql.Tx),
		), nil
	})

	a.container.AddScoped(constants.ApplicationKey, func(c di.Container) (any, error) {
		publisher := c.Get(constants.DomainDispatcherKey).(ddd.EventPublisher[ddd.Event])
		stores := c.Get(constants.StoresRepoKey).(domain.StoreRepository)
		products := c.Get(constants.ProductsRepoKey).(domain.ProductRepository)
		catalog := c.Get(constants.CatalogRepoKey).(domain.CatalogRepository)
		malls := c.Get(constants.MallRepoKey).(domain.MallRepository)
		log := c.Get(constants.LoggerKey).(logger.Logger)

		//fmt.Println("pubsher", publisher)
		application := &app.Application{
			Commands: app.Commands{
				CreateStore: commands.NewCreateStoreHandler(stores, publisher, log),
				AddProduct:  commands.NewAddProductHandler(products, publisher, log),
			},
			Queries: app.Queries{
				GetProduct: queries.NewGetProductHandler(catalog, log),
				GetStore:   queries.NewGetStoreHandler(malls, log),
			},
		}
		//a.Application = application
		return application, nil
	})

	a.container.AddScoped(constants.CatalogHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewCatalogHandlers(c.Get(constants.CatalogRepoKey).(domain.CatalogRepository)), nil
	})

	a.container.AddScoped(constants.MallHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewMallHandlers(c.Get(constants.MallRepoKey).(domain.MallRepository)), nil
	})

	a.container.AddScoped(constants.DomainEventHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewDomainEventHandlers(c.Get(constants.EventPublisherKey).(am.EventPublisher)), nil
	})
	outboxProcessor := tm.NewOutboxProcessor(
		stream,
		postgres.NewOutboxStore(constants.OutboxTableName, dbConn),
	)
	handlers.RegisterCatalogHandlersTx(a.container)
	handlers.RegisterMallHandlersTx(a.container)
	handlers.RegisterDomainEventHandlersTx(a.container)

	startOutboxProcessor(context.Background(), outboxProcessor, a.container.Get(constants.LoggerKey).(logger.Logger))
	return nil
}
func startOutboxProcessor(ctx context.Context, outboxProcessor tm.OutboxProcessor, logger logger.Logger) {
	go func() {
		err := outboxProcessor.Start(ctx)
		if err != nil {
			//logger.Error().Err(err).Msg("customers outbox processor encountered an error")
		}
	}()
}
func (a *Agent) nats() (nats.JetStreamContext, error) {
	nc, err := nats.Connect(a.NatsURL)
	if err != nil {
		return nil, err
	}
	// defer nc.Close()
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     a.NatsStream,
		Subjects: []string{fmt.Sprintf("%s.>", a.NatsStream)},
	})

	return js, err
}
