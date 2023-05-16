package agent

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rezaAmiri123/microservice/cosec/internal"
	"github.com/rezaAmiri123/microservice/cosec/internal/constants"
	"github.com/rezaAmiri123/microservice/cosec/internal/domain"
	"github.com/rezaAmiri123/microservice/cosec/internal/handlers"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/amotel"
	"github.com/rezaAmiri123/microservice/pkg/amprom"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/pkg/db/postgresotel"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/sec"
	"github.com/rezaAmiri123/microservice/pkg/tm"
)

func (a *Agent) setupApplication() (err error) {
	//stream := jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger())
	//js, _ := a.nats()
	//stream := jetstream.NewStream(a.NatsStream, js, a.container.Get(constants.LoggerKey).(logger.Logger))
	dbConn, err := postgres.NewPsqlDB(postgres.Config{
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

	a.container.AddScoped(constants.DatabaseTransactionKey, func(c di.Container) (any, error) {
		return dbConn.Begin()
	})
	sentCounter := amprom.SentMessagesCounter(constants.ServiceName)
	a.container.AddScoped(constants.MessagePublisherKey, func(c di.Container) (any, error) {
		tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
		outboxStore := postgres.NewOutboxStore(constants.OutboxTableName, tx)
		publisher := am.NewMessagePublisher(
			c.Get(constants.StreamKey).(am.MessageStream),
			amotel.OtelMessageContextInjector(),
			sentCounter,
			tm.OutboxPublisher(outboxStore),
		)
		return publisher, nil
	})
	a.container.AddSingleton(constants.MessageSubscriberKey, func(c di.Container) (any, error) {
		return am.NewMessageSubscriber(
			c.Get(constants.StreamKey).(am.MessageStream),
			amotel.OtelMessageContextExtractor(),
			amprom.ReceivedMessagesCounter(constants.ServiceName),
		), nil
	})
	a.container.AddScoped(constants.CommandPublisherKey, func(c di.Container) (any, error) {
		return am.NewCommandPublisher(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.MessagePublisherKey).(am.MessagePublisher),
		), nil
	})
	a.container.AddScoped(constants.InboxStoreKey, func(c di.Container) (any, error) {
		tx := postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx))
		return postgres.NewInboxStore(constants.InboxTableName, tx), nil
	})
	a.container.AddScoped(constants.SagaStoreKey, func(c di.Container) (any, error) {
		reg := c.Get(constants.RegistryKey).(registry.Registry)
		return sec.NewSagaRepository[*domain.CreateOrderData](
			reg,
			postgres.NewSagaStore(
				constants.SagasTableName,
				postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
				reg,
			),
		), nil
	})
	a.container.AddSingleton(constants.SagaKey, func(c di.Container) (any, error) {
		return internal.NewCreateOrderSaga(), nil
	})

	// setup application
	a.container.AddScoped(constants.OrchestratorKey, func(c di.Container) (any, error) {
		return sec.NewOrchestrator[*domain.CreateOrderData](
			c.Get(constants.SagaKey).(sec.Saga[*domain.CreateOrderData]),
			c.Get(constants.SagaStoreKey).(sec.SagaRepository[*domain.CreateOrderData]),
			c.Get(constants.CommandPublisherKey).(am.CommandPublisher),
		), nil
	})
	a.container.AddScoped(constants.IntegrationEventHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewIntegrationEventHandlers(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.OrchestratorKey).(sec.Orchestrator[*domain.CreateOrderData]),
			tm.InboxHandler(c.Get(constants.InboxStoreKey).(tm.InboxStore)),
		), nil
	})
	a.container.AddScoped(constants.ReplyHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewReplyHandlers(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.OrchestratorKey).(sec.Orchestrator[*domain.CreateOrderData]),
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
	if err = handlers.RegisterReplyHandlersTx(a.container); err != nil {
		return err
	}
	startOutboxProcessor(context.Background(), outboxProcessor, a.container.Get(constants.LoggerKey).(logger.Logger))

	return
}
func startOutboxProcessor(ctx context.Context, outboxProcessor tm.OutboxProcessor, logger logger.Logger) {
	go func() {
		err := outboxProcessor.Start(ctx)
		if err != nil {
			//logger.Error().Err(err).Msg("customers outbox processor encountered an error")
		}
	}()
}
