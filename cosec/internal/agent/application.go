package agent

import (
	"database/sql"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/rezaAmiri123/microservice/cosec/internal"
	"github.com/rezaAmiri123/microservice/cosec/internal/constants"
	"github.com/rezaAmiri123/microservice/cosec/internal/domain"
	"github.com/rezaAmiri123/microservice/cosec/internal/handlers"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/jetstream"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/registry/serdes"
	"github.com/rezaAmiri123/microservice/pkg/sec"
	"github.com/rezaAmiri123/microservice/pkg/tm"
	"github.com/rezaAmiri123/microservice/users/userspb"
)

func (a *Agent) setupApplication() (err error) {
	container := di.New()
	// setup Driven adapters
	container.AddSingleton(constants.RegistryKey, func(c di.Container) (any, error) {
		reg := registry.New()
		if err := registrations(reg); err != nil {
			return nil, err
		}
		//if err := orderingpb.Registrations(reg); err != nil {
		//	return nil, err
		//}
		if err := userspb.Registrations(reg); err != nil {
			return nil, err
		}
		//if err := depotpb.Registrations(reg); err != nil {
		//	return nil, err
		//}
		//if err := paymentspb.Registrations(reg); err != nil {
		//	return nil, err
		//}
		return reg, nil
	})
	//stream := jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger())
	js, _ := a.nats()
	stream := jetstream.NewStream("stream", js, a.logger)
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

	container.AddScoped(constants.DatabaseTransactionKey, func(c di.Container) (any, error) {
		return dbConn.Begin()
	})
	//sentCounter := amprom.SentMessagesCounter(constants.ServiceName)
	container.AddScoped(constants.MessagePublisherKey, func(c di.Container) (any, error) {
		tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
		outboxStore := postgres.NewOutboxStore(constants.OutboxTableName, tx)
		publisher := am.NewMessagePublisher(
			stream,
			//amotel.OtelMessageContextInjector(),
			//sentCounter,
			tm.OutboxPublisher(outboxStore),
		)
		return publisher, nil
	})
	container.AddSingleton(constants.MessageSubscriberKey, func(c di.Container) (any, error) {
		return am.NewMessageSubscriber(
			stream,
			//amotel.OtelMessageContextExtractor(),
			//amprom.ReceivedMessagesCounter(constants.ServiceName),
		), nil
	})
	container.AddScoped(constants.CommandPublisherKey, func(c di.Container) (any, error) {
		return am.NewCommandPublisher(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.MessagePublisherKey).(am.MessagePublisher),
		), nil
	})
	container.AddScoped(constants.InboxStoreKey, func(c di.Container) (any, error) {
		tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
		return postgres.NewInboxStore(constants.InboxTableName, tx), nil
	})
	container.AddScoped(constants.SagaStoreKey, func(c di.Container) (any, error) {
		reg := c.Get(constants.RegistryKey).(registry.Registry)
		return sec.NewSagaRepository[*domain.CreateOrderData](
			reg,
			postgres.NewSagaStore(
				constants.SagasTableName,
				c.Get(constants.DatabaseTransactionKey).(*sql.Tx),
				reg,
			),
		), nil
	})
	container.AddSingleton(constants.SagaKey, func(c di.Container) (any, error) {
		return internal.NewCreateOrderSaga(), nil
	})

	// setup application
	container.AddScoped(constants.OrchestratorKey, func(c di.Container) (any, error) {
		return sec.NewOrchestrator[*domain.CreateOrderData](
			c.Get(constants.SagaKey).(sec.Saga[*domain.CreateOrderData]),
			c.Get(constants.SagaStoreKey).(sec.SagaRepository[*domain.CreateOrderData]),
			c.Get(constants.CommandPublisherKey).(am.CommandPublisher),
		), nil
	})
	container.AddScoped(constants.IntegrationEventHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewIntegrationEventHandlers(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.OrchestratorKey).(sec.Orchestrator[*domain.CreateOrderData]),
			tm.InboxHandler(c.Get(constants.InboxStoreKey).(tm.InboxStore)),
		), nil
	})
	container.AddScoped(constants.ReplyHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewReplyHandlers(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.OrchestratorKey).(sec.Orchestrator[*domain.CreateOrderData]),
			tm.InboxHandler(c.Get(constants.InboxStoreKey).(tm.InboxStore)),
		), nil
	})
	//outboxProcessor := tm.NewOutboxProcessor(
	//	stream,
	//	postgres.NewOutboxStore(constants.OutboxTableName, dbConn),
	//)

	// setup Driver adapters
	if err = handlers.RegisterIntegrationEventHandlersTx(container); err != nil {
		return err
	}
	if err = handlers.RegisterReplyHandlersTx(container); err != nil {
		return err
	}
	//startOutboxProcessor(ctx, outboxProcessor, svc.Logger())

	return
}
func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Saga data
	if err = serde.RegisterKey(internal.CreateOrderSagaName, domain.CreateOrderData{}); err != nil {
		return err
	}

	return nil
}

func (a *Agent) nats() (nats.JetStreamContext, error) {
	nc, err := nats.Connect("localhost")
	if err != nil {
		return nil, err
	}
	// defer nc.Close()
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "stream",
		Subjects: []string{fmt.Sprintf("%s.>", "stream")},
	})

	return js, err
}
