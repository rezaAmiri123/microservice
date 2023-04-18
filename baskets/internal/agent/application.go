package agent

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nats-io/nats.go"

	"github.com/rezaAmiri123/microservice/baskets/internal/adapters/grpc"
	"github.com/rezaAmiri123/microservice/baskets/internal/adapters/pg"
	"github.com/rezaAmiri123/microservice/baskets/internal/app"
	"github.com/rezaAmiri123/microservice/baskets/internal/constants"
	"github.com/rezaAmiri123/microservice/baskets/internal/domain"
	"github.com/rezaAmiri123/microservice/baskets/internal/ports/handlers"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/amotel"
	"github.com/rezaAmiri123/microservice/pkg/amprom"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/pkg/db/postgresotel"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/jetstream"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/tm"
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
	if err = dbConn.Ping(); err != nil {
		fmt.Println("cannot ping db")
		return fmt.Errorf("cannot ping db: %w", err)
	}

	a.container.AddSingleton(constants.DatabaseKey, func(c di.Container) (any, error) {
		return dbConn, nil
	})

	js, err := a.nats()
	if err != nil {
		return err
	}

	stream := jetstream.NewStream(a.NatsStream, js, a.container.Get(constants.LoggerKey).(logger.Logger))

	a.container.AddSingleton(constants.DomainDispatcherKey, func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.Event](), nil
	})
	a.container.AddScoped(constants.DatabaseTransactionKey, func(c di.Container) (any, error) {
		//return c.Get(constants.DatabaseKey).(*sql.DB).Begin()
		return dbConn.Begin()
	})

	sentCounter := amprom.SentMessagesCounter(constants.ServiceName)
	a.container.AddScoped(constants.MessagePublisherKey, func(c di.Container) (any, error) {
		db := postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB))
		outboxStore := postgres.NewOutboxStore(constants.OutboxTableName, db)
		return am.NewMessagePublisher(
			stream,
			amotel.OtelMessageContextInjector(),
			sentCounter,
			tm.OutboxPublisher(outboxStore),
		), nil
	})
	a.container.AddScoped(constants.MessageSubscriberKey, func(c di.Container) (any, error) {
		return am.NewMessageSubscriber(
			stream,
			amotel.OtelMessageContextExtractor(),
			amprom.ReceivedMessagesCounter(constants.ServiceName),
		), nil
	})
	a.container.AddScoped(constants.EventPublisherKey, func(c di.Container) (any, error) {
		return am.NewEventPublisher(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.MessagePublisherKey).(am.MessagePublisher),
		), nil
	})

	a.container.AddScoped(constants.InboxStoreKey, func(c di.Container) (any, error) {
		db := postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB))
		return postgres.NewInboxStore(constants.InboxTableName, db), nil
	})

	//a.container.AddScoped(constants.AggregateStoreKey, func(c di.Container) (any, error) {
	//	tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
	//	reg := c.Get(constants.RegistryKey).(registry.Registry)
	//	return es.AggregateStoreWithMiddleware(
	//		postgres.NewEventStore(constants.EventsTableName, tx, reg),
	//		postgres.NewSnapshotStore(constants.SnapshotsTableName, tx, reg),
	//	), nil
	//})

	a.container.AddScoped(constants.BasketsRepoKey, func(c di.Container) (any, error) {
		//tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
		//reg := c.Get(constants.RegistryKey).(registry.Registry)
		//return es.NewAggregateRepository[*domain.Basket](
		//	domain.BasketAggregate,
		//	reg,
		//	es.AggregateStoreWithMiddleware(
		//		postgres.NewEventStore(constants.EventsTableName, tx, reg),
		//		postgres.NewSnapshotStore(constants.SnapshotsTableName, tx, reg),
		//	),
		//), nil
		//return pg.NewBasketRepository(
		//	constants.BasketTableName,
		//	c.Get(constants.DatabaseTransactionKey).(*sql.Tx),
		//), nil
		return pg.NewBasketRepository(
			constants.BasketTableName,
			postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB)),
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

	//a.container.AddScoped(constants.ProductsRepoKey, func(c di.Container) (any, error) {
	//	return es.NewAggregateRepository[*domain.Product](
	//		domain.ProductAggregate,
	//		c.Get(constants.RegistryKey).(registry.Registry),
	//		c.Get(constants.AggregateStoreKey).(es.AggregateStore),
	//	), nil
	//})
	//a.container.AddScoped(constants.CatalogRepoKey, func(c di.Container) (any, error) {
	//	return pg.NewCatalogRepository(
	//		constants.CatalogTableName,
	//		c.Get(constants.DatabaseTransactionKey).(*sql.Tx),
	//	), nil
	//})
	//
	//a.container.AddScoped(constants.MallRepoKey, func(c di.Container) (any, error) {
	//	return pg.NewMallRepository(
	//		constants.MallTableName,
	//		c.Get(constants.DatabaseTransactionKey).(*sql.Tx),
	//	), nil
	//})

	a.container.AddScoped(constants.ApplicationKey, func(c di.Container) (any, error) {
		publisher := c.Get(constants.DomainDispatcherKey).(ddd.EventPublisher[ddd.Event])
		baskets := c.Get(constants.BasketsRepoKey).(domain.BasketRepository)
		stores := c.Get(constants.StoresRepoKey).(domain.StoreCacheRepository)
		products := c.Get(constants.ProductsRepoKey).(domain.ProductCacheRepository)

		log := c.Get(constants.LoggerKey).(logger.Logger)

		//fmt.Println("pubsher", publisher)
		application := app.NewInstrumentedApp(
			app.New(baskets, stores, products, publisher, log),
		)
		//application := &app.Application{
		//	Commands: app.Commands{
		//		StartBasket:    commands.NewStartBasketHandler(baskets, publisher, log),
		//		AddItem:        commands.NewAddItemHandler(baskets, stores, products, publisher, log),
		//		CheckoutBasket: commands.NewCheckoutBasketHandler(baskets, publisher, log),
		//		CancelBasket:   commands.NewCancelBasketHandler(baskets, publisher, log),
		//	},
		//	Queries: app.Queries{},
		//}
		//a.Application = application
		return application, nil
	})
	a.container.AddScoped(constants.DomainEventHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewDomainEventHandlers(c.Get(constants.EventPublisherKey).(am.EventPublisher)), nil
	})
	a.container.AddScoped(constants.IntegrationEventHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewIntegrationEventHandlers(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.StoresRepoKey).(domain.StoreCacheRepository),
			c.Get(constants.ProductsRepoKey).(domain.ProductCacheRepository),
			tm.InboxHandler(c.Get(constants.InboxStoreKey).(tm.InboxStore)),
		), nil
	})

	outboxProcessor := tm.NewOutboxProcessor(
		stream,
		postgres.NewOutboxStore(constants.OutboxTableName, dbConn),
	)
	handlers.RegisterDomainEventHandlersTx(a.container)
	if err = handlers.RegisterIntegrationEventHandlersTx(a.container); err != nil {
		return err
	}
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
