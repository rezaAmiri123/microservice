package agent

import (
	"database/sql"
	"github.com/rezaAmiri123/microservice/ordering/internal/app"
	"github.com/rezaAmiri123/microservice/ordering/internal/constants"
	"github.com/rezaAmiri123/microservice/ordering/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/es"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/pkg/registry"
)

func (a *Agent) setupApplication() error {
	//a.container.AddSingleton(constants.DomainDispatcherKey, func(c di.Container) (any, error) {
	//	return ddd.NewEventDispatcher[ddd.Event](), nil
	//})
	//
	//sentCounter := amprom.SentMessagesCounter(constants.ServiceName)
	//a.container.AddScoped(constants.MessagePublisherKey, func(c di.Container) (any, error) {
	//	//tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
	//	db := c.Get(constants.DatabaseKey).(*sql.DB)
	//	outboxStore := postgres.NewOutboxStore(constants.OutboxTableName, db)
	//	return am.NewMessagePublisher(
	//		c.Get(constants.StreamKey).(am.MessageStream),
	//		amotel.OtelMessageContextInjector(),
	//		sentCounter,
	//		tm.OutboxPublisher(outboxStore),
	//	), nil
	//})
	//a.container.AddScoped(constants.MessageSubscriberKey, func(c di.Container) (any, error) {
	//	return am.NewMessageSubscriber(
	//		c.Get(constants.StreamKey).(am.MessageStream),
	//		amotel.OtelMessageContextExtractor(),
	//		amprom.ReceivedMessagesCounter(constants.ServiceName),
	//	), nil
	//})
	//a.container.AddScoped(constants.EventPublisherKey, func(c di.Container) (any, error) {
	//	return am.NewEventPublisher(
	//		c.Get(constants.RegistryKey).(registry.Registry),
	//		c.Get(constants.MessagePublisherKey).(am.MessagePublisher),
	//	), nil
	//})

	//a.container.AddScoped(constants.CommandPublisherKey, func(c di.Container) (any, error) {
	//	return am.NewCommandPublisher(
	//		c.Get(constants.RegistryKey).(registry.Registry),
	//		c.Get(constants.MessagePublisherKey).(am.MessagePublisher),
	//	), nil
	//})

	//a.container.AddScoped(constants.InboxStoreKey, func(c di.Container) (any, error) {
	//	//tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
	//	db := postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB))
	//	return postgres.NewInboxStore(constants.InboxTableName, db), nil
	//})

	//a.container.AddScoped(constants.OrdersRepoKey, func(c di.Container) (any, error) {
	//	//tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
	//	db := postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB))
	//	return pg.NewOrderRepository(constants.OrdersTableName, db), nil
	//})

	a.container.AddScoped(constants.OrdersRepoKey, func(c di.Container) (any, error) {
		db := c.Get(constants.DatabaseKey).(*sql.DB)
		reg := c.Get(constants.RegistryKey).(registry.Registry)
		return es.NewAggregateRepository[*domain.Order](
			domain.OrderAggregate,
			c.Get(constants.RegistryKey).(registry.Registry),
			es.AggregateStoreWithMiddleware(
				postgres.NewEventStore(constants.EventsTableName, db, reg),
				postgres.NewSnapshotStore(constants.SnapshotsTableName, db, reg),
			),
		), nil
	})

	// setup application
	a.container.AddScoped(constants.ApplicationKey, func(c di.Container) (any, error) {
		orders := c.Get(constants.OrdersRepoKey).(domain.OrderRepository)
		dispatcher := c.Get(constants.DomainDispatcherKey).(*ddd.EventDispatcher[ddd.Event])
		log := c.Get(constants.LoggerKey).(logger.Logger)

		//fmt.Println("pubsher", publisher)
		application := app.NewInstrumentedApp(
			app.New(orders, dispatcher, log),
		)
		return application, nil
	})

	//a.container.AddScoped(constants.DomainEventHandlersKey, func(c di.Container) (any, error) {
	//	return handlers.NewDomainEventHandlers(c.Get(constants.EventPublisherKey).(am.EventPublisher)), nil
	//})

	//a.container.AddScoped(constants.IntegrationEventHandlersKey, func(c di.Container) (any, error) {
	//	return handlers.NewIntegrationEventHandlers(
	//		c.Get(constants.RegistryKey).(registry.Registry),
	//		c.Get(constants.ApplicationKey).(app.App),
	//		tm.InboxHandler(c.Get(constants.InboxStoreKey).(tm.InboxStore)),
	//	), nil
	//})

	//a.container.AddScoped(constants.CommandHandlersKey, func(c di.Container) (any, error) {
	//	return handlers.NewCommandHandlers(
	//		c.Get(constants.RegistryKey).(registry.Registry),
	//		c.Get(constants.ApplicationKey).(app.App),
	//		c.Get(constants.ReplyPublisherKey).(am.ReplyPublisher),
	//		tm.InboxHandler(c.Get(constants.InboxStoreKey).(tm.InboxStore)),
	//	), nil
	//})

	//outboxProcessor := tm.NewOutboxProcessor(
	//	a.container.Get(constants.StreamKey).(am.MessageStream),
	//	postgres.NewOutboxStore(constants.OutboxTableName, dbConn),
	//)
	// setup Driver adapters
	//if err = handlers.RegisterIntegrationEventHandlersTx(a.container); err != nil {
	//	return err
	//}
	//handlers.RegisterDomainEventHandlersTx(a.container)
	//if err = handlers.RegisterIntegrationEventHandlersTx(a.container); err != nil {
	//	return err
	//}
	//if err = handlers.RegisterCommandHandlersTx(a.container); err != nil {
	//	return err
	//}
	//
	//startOutboxProcessor(context.Background(), outboxProcessor, a.container.Get(constants.LoggerKey).(logger.Logger))

	return nil
}
