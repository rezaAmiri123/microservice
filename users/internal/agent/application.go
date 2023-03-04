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
	"github.com/rezaAmiri123/microservice/pkg/jetstream"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/tm"
	"github.com/rezaAmiri123/microservice/users/internal/adapters/pg"
	"github.com/rezaAmiri123/microservice/users/internal/app"
	"github.com/rezaAmiri123/microservice/users/internal/app/commands"
	"github.com/rezaAmiri123/microservice/users/internal/constants"
	"github.com/rezaAmiri123/microservice/users/internal/handlers"
	"github.com/rezaAmiri123/microservice/users/userspb"
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
	//if err = postgres.MigrateUp(dbConn, migrations.FS); err != nil {
	//	return err
	//}
	a.container.AddSingleton(constants.DatabaseKey, func(c di.Container) (any, error) {
		return dbConn, nil
	})

	//repo, err := adapters.NewGORMArticleRepository(a.DBConfig)
	repo := pg.NewPGUserRepository(dbConn, a.logger)

	// setup Driven adapters
	reg := registry.New()
	a.container.AddSingleton(constants.RegistryKey, func(c di.Container) (any, error) {
		return reg, nil
	})
	if err = userspb.Registrations(reg); err != nil {
		return err
	}
	js, err := a.nats()
	if err != nil {
		return err
	}
	stream := jetstream.NewStream(a.NatsStream, js, a.logger)
	//a.container.AddSingleton("stream", func(c di.Container) (any, error) {
	//	return stream, nil
	//})
	a.container.AddScoped(constants.DatabaseTransactionKey, func(c di.Container) (any, error) {
		return dbConn.Begin()
	})
	a.container.AddSingleton(constants.DomainDispatcherKey, func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.AggregateEvent](), nil
	})

	a.container.AddScoped(constants.UsersRepoKey, func(c di.Container) (any, error) {
		return pg.NewPGUserRepository(c.Get(constants.DatabaseTransactionKey).(*sql.Tx), a.logger), nil
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

	a.container.AddScoped(constants.ReplyPublisherKey, func(c di.Container) (any, error) {
		return am.NewReplyPublisher(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.MessagePublisherKey).(am.MessagePublisher),
		), nil
	})

	a.container.AddScoped(constants.InboxStoreKey, func(c di.Container) (any, error) {
		tx := c.Get(constants.DatabaseTransactionKey).(*sql.Tx)
		return postgres.NewInboxStore(constants.InboxTableName, tx), nil
	})

	a.container.AddScoped(constants.ApplicationKey, func(c di.Container) (any, error) {
		domainDispatcher := a.container.Get(constants.DomainDispatcherKey).(*ddd.EventDispatcher[ddd.AggregateEvent])
		application := &app.Application{
			Commands: app.Commands{
				RegisterUser: commands.NewRegisterUserHandler(repo, a.logger, domainDispatcher),
				EnableUser:   commands.NewEnableUserHandler(repo, a.logger, domainDispatcher),
			},
			Queries: app.Queries{},
		}
		a.Application = application
		return application, nil
	})
	a.container.AddScoped(constants.DomainEventHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewDomainEventHandlers(c.Get(constants.EventPublisherKey).(am.EventPublisher)), nil
	})

	a.container.AddScoped(constants.CommandHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewCommandHandlers(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.ApplicationKey).(*app.Application),
			c.Get(constants.ReplyPublisherKey).(am.ReplyPublisher),
			tm.InboxHandler(c.Get(constants.InboxStoreKey).(tm.InboxStore)),
		), nil
	})

	outboxProcessor := tm.NewOutboxProcessor(
		stream,
		postgres.NewOutboxStore(constants.OutboxTableName, dbConn),
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
