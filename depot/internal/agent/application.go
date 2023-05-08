package agent

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/rezaAmiri123/microservice/depot/internal/adapters/grpc"
	"github.com/rezaAmiri123/microservice/depot/internal/adapters/pg"
	"github.com/rezaAmiri123/microservice/depot/internal/app"
	"github.com/rezaAmiri123/microservice/depot/internal/constants"
	"github.com/rezaAmiri123/microservice/depot/internal/domain"
	"github.com/rezaAmiri123/microservice/depot/internal/ports/handlers"
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
		PGDriver:     a.PGDriver,
		PGHost:       a.PGHost,
		PGPort:       a.PGPort,
		PGUser:       a.PGUser,
		PGDBName:     a.PGDBName,
		PGPassword:   a.PGPassword,
		PGSearchPath: a.PGSearchPath,
	})
	if err != nil {
		return fmt.Errorf("cannot load db: %w", err)
	}

	a.container.AddSingleton(constants.DatabaseKey, func(c di.Container) (any, error) {
		return dbConn, nil
	})

	js, err := a.nats()
	if err != nil {
		return err
	}

	stream := jetstream.NewStream(a.NatsStream, js, a.container.Get(constants.LoggerKey).(logger.Logger))
	//a.container.AddSingleton(constants.DomainDispatcherKey, func(c di.Container) (any, error) {
	//	return ddd.NewEventDispatcher[ddd.AggregateEvent](), nil
	//})
	a.container.AddSingleton(constants.DomainDispatcherKey, func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.AggregateEvent](), nil
	})
	a.container.AddScoped(constants.DatabaseTransactionKey, func(c di.Container) (any, error) {
		return dbConn.Begin()
	})

	sentCounter := amprom.SentMessagesCounter(constants.ServiceName)
	a.container.AddScoped(constants.MessagePublisherKey, func(c di.Container) (any, error) {
		//tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
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

	a.container.AddScoped(constants.CommandPublisherKey, func(c di.Container) (any, error) {
		return am.NewCommandPublisher(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.MessagePublisherKey).(am.MessagePublisher),
		), nil
	})

	a.container.AddScoped(constants.ReplyPublisherKey, func(c di.Container) (any, error) {
		return am.NewReplyPublisher(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.MessagePublisherKey).(am.MessagePublisher),
		), nil
	})

	a.container.AddScoped(constants.InboxStoreKey, func(c di.Container) (any, error) {
		//tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
		db := postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB))
		return postgres.NewInboxStore(constants.InboxTableName, db), nil
	})

	a.container.AddScoped(constants.ShoppingListsRepoKey, func(c di.Container) (any, error) {
		//tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
		db := postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB))
		return pg.NewShoppingListRepository(constants.ShoppingListsTableName, db), nil
	})

	a.container.AddScoped(constants.StoresCacheRepoKey, func(c di.Container) (any, error) {
		//tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
		db := postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB))
		addr := fmt.Sprintf("%s:%d", a.Config.GRPCStoreClientAddr, a.Config.GRPCStoreClientPort)
		return pg.NewStoreCacheRepository(
			constants.StoresCacheTableName,
			db,
			grpc.NewStoreRepository(addr, c.Get(constants.LoggerKey).(logger.Logger)),
		), nil
	})

	a.container.AddScoped(constants.ProductsCacheRepoKey, func(c di.Container) (any, error) {
		//tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
		db := postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB))
		addr := fmt.Sprintf("%s:%d", a.Config.GRPCStoreClientAddr, a.Config.GRPCStoreClientPort)
		return pg.NewProductCacheRepository(
			constants.ProductsCacheTableName,
			db,
			grpc.NewProductRepository(addr, c.Get(constants.LoggerKey).(logger.Logger)),
		), nil
	})
	// setup application
	a.container.AddScoped(constants.ApplicationKey, func(c di.Container) (any, error) {
		//publisher := c.Get(constants.DomainDispatcherKey).(ddd.EventPublisher[ddd.Event])
		shoppingList := c.Get(constants.ShoppingListsRepoKey).(domain.ShoppingListRepository)
		stores := c.Get(constants.StoresCacheRepoKey).(domain.StoreCacheRepository)
		products := c.Get(constants.ProductsCacheRepoKey).(domain.ProductCacheRepository)
		dispatcher := c.Get(constants.DomainDispatcherKey).(*ddd.EventDispatcher[ddd.AggregateEvent])
		log := c.Get(constants.LoggerKey).(logger.Logger)

		//fmt.Println("pubsher", publisher)
		application := app.NewInstrumentedApp(
			app.New(shoppingList, stores, products, dispatcher, log),
		)
		//a.Application = application
		return application, nil
	})

	a.container.AddScoped(constants.DomainEventHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewDomainEventHandlers(c.Get(constants.EventPublisherKey).(am.EventPublisher)), nil
	})

	a.container.AddScoped(constants.IntegrationEventHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewIntegrationEventHandlers(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.StoresCacheRepoKey).(domain.StoreCacheRepository),
			c.Get(constants.ProductsCacheRepoKey).(domain.ProductCacheRepository),
			tm.InboxHandler(c.Get(constants.InboxStoreKey).(tm.InboxStore)),
		), nil
	})

	a.container.AddScoped(constants.CommandHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewCommandHandlers(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.ApplicationKey).(app.App),
			c.Get(constants.ReplyPublisherKey).(am.ReplyPublisher),
			tm.InboxHandler(c.Get(constants.InboxStoreKey).(tm.InboxStore)),
		), nil
	})

	outboxProcessor := tm.NewOutboxProcessor(
		stream,
		postgres.NewOutboxStore(constants.OutboxTableName, dbConn),
	)
	// setup Driver adapters
	//if err = handlers.RegisterIntegrationEventHandlersTx(a.container); err != nil {
	//	return err
	//}
	handlers.RegisterDomainEventHandlersTx(a.container)
	if err = handlers.RegisterIntegrationEventHandlersTx(a.container); err != nil {
		return err
	}
	if err = handlers.RegisterCommandHandlersTx(a.container); err != nil {
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
