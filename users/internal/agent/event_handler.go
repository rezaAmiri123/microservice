package agent

import (
	"context"
	"database/sql"
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
	"github.com/rezaAmiri123/microservice/users/internal/app"
	"github.com/rezaAmiri123/microservice/users/internal/constants"
	"github.com/rezaAmiri123/microservice/users/internal/ports/handlers"
	"github.com/stackus/errors"
)

func (a *Agent) setupEventHandler() (err error) {
	a.container.AddSingleton(constants.DomainDispatcherKey, func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.AggregateEvent](), nil
	})

	sentCounter := amprom.SentMessagesCounter(constants.ServiceName)
	a.container.AddScoped(constants.MessagePublisherKey, func(c di.Container) (any, error) {
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
		db := postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB))
		return postgres.NewInboxStore(constants.InboxTableName, db), nil
	})

	a.container.AddScoped(constants.DomainEventHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewDomainEventHandlers(
			c.Get(constants.EventPublisherKey).(am.EventPublisher),
			c.Get(constants.LoggerKey).(logger.Logger),
		), nil
	})

	a.container.AddScoped(constants.CommandHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewCommandHandlers(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.LoggerKey).(logger.Logger),
			c.Get(constants.ApplicationKey).(app.App),
			c.Get(constants.ReplyPublisherKey).(am.ReplyPublisher),
			tm.InboxHandler(c.Get(constants.InboxStoreKey).(tm.InboxStore)),
		), nil
	})

	outboxProcessor := tm.NewOutboxProcessor(
		a.container.Get(constants.StreamKey).(am.MessageStream),
		postgres.NewOutboxStore(
			constants.OutboxTableName,
			a.container.Get(constants.DatabaseKey).(*sql.DB),
		),
	)

	handlers.RegisterDomainEventHandlersTx(a.container)
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
			logger.Error(errors.Wrap(err, "outbox processor encountered an error"))
			//logger.Error().Err(err).Msg("customers outbox processor encountered an error")
		}
	}()
}
