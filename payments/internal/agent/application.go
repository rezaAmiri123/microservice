package agent

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rezaAmiri123/microservice/payments/internal/adapters/migrations"
	"github.com/rezaAmiri123/microservice/payments/internal/adapters/pg"
	"github.com/rezaAmiri123/microservice/payments/internal/app"
	"github.com/rezaAmiri123/microservice/payments/internal/constants"
	"github.com/rezaAmiri123/microservice/payments/internal/domain"
	"github.com/rezaAmiri123/microservice/payments/internal/ports/handlers"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/amotel"
	"github.com/rezaAmiri123/microservice/pkg/amprom"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/pkg/db/postgresotel"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/di"
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

	if err = postgres.MigrateUp(dbConn, migrations.FS); err != nil {
		return err
	}

	a.container.AddSingleton(constants.DatabaseKey, func(c di.Container) (any, error) {
		return dbConn, nil
	})

	a.container.AddSingleton(constants.DomainDispatcherKey, func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.Event](), nil
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
			c.Get(constants.StreamKey).(am.MessageStream),
			amotel.OtelMessageContextInjector(),
			sentCounter,
			tm.OutboxPublisher(outboxStore),
		), nil
	})
	a.container.AddScoped(constants.MessageSubscriberKey, func(c di.Container) (any, error) {
		return am.NewMessageSubscriber(
			c.Get(constants.StreamKey).(am.MessageStream),
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

	a.container.AddScoped(constants.InvoicesRepoKey, func(c di.Container) (any, error) {
		//tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
		db := postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB))
		return pg.NewInvoiceRepository(constants.InvoicesTableName, db), nil
	})

	a.container.AddScoped(constants.PaymentsRepoKey, func(c di.Container) (any, error) {
		//tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
		db := postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB))
		return pg.NewPaymentRepository(constants.PaymentsTableName, db), nil
	})

	// setup application
	a.container.AddScoped(constants.ApplicationKey, func(c di.Container) (any, error) {
		publisher := c.Get(constants.DomainDispatcherKey).(ddd.EventPublisher[ddd.Event])
		invoices := c.Get(constants.InvoicesRepoKey).(domain.InvoiceRepository)
		payments := c.Get(constants.PaymentsRepoKey).(domain.PaymentRepository)
		log := c.Get(constants.LoggerKey).(logger.Logger)

		//fmt.Println("pubsher", publisher)
		application := app.NewInstrumentedApp(
			app.New(invoices, payments, publisher, log),
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
			c.Get(constants.ApplicationKey).(app.App),
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
		a.container.Get(constants.StreamKey).(am.MessageStream),
		postgres.NewOutboxStore(constants.OutboxTableName, dbConn),
	)
	// setup Driver adapters
	if err = handlers.RegisterIntegrationEventHandlersTx(a.container); err != nil {
		return err
	}
	handlers.RegisterDomainEventHandlersTx(a.container)
	if err = handlers.RegisterCommandHandlersTx(a.container); err != nil {
		return err
	}

	startOutboxProcessor(
		context.Background(),
		outboxProcessor,
		a.container.Get(constants.LoggerKey).(logger.Logger),
	)

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
