package agent

import (
	"context"
	"database/sql"
	"fmt"
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
	"github.com/rezaAmiri123/microservice/users/internal/adapters/pg"
	"github.com/rezaAmiri123/microservice/users/internal/app"
	"github.com/rezaAmiri123/microservice/users/internal/constants"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
	"github.com/rezaAmiri123/microservice/users/internal/handlers"
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
	//if err = postgres.MigrateUp(dbConn, migrations.FS); err != nil {
	//	return err
	//}
	a.container.AddSingleton(constants.DatabaseKey, func(c di.Container) (any, error) {
		return dbConn, nil
	})

	//repo, err := adapters.NewGORMArticleRepository(a.DBConfig)
	//repo := pg.NewPGUserRepository(dbConn, a.logger)

	// setup Driven adapters
	//reg := registry.New()
	//a.container.AddSingleton(constants.RegistryKey, func(c di.Container) (any, error) {
	//	return reg, nil
	//})
	//if err = userspb.Registrations(reg); err != nil {
	//	return err
	//}
	//log := a.container.Get(constants.LoggerKey).(logger.Logger)
	//js, err := a.nats()
	//if err != nil {
	//	return err
	//}
	//stream1 := jetstream.NewStream(a.NatsStream, js, log)
	//fmt.Println(stream1)
	//fmt.Println(log)
	//stream := kafkastream.NewStream(a.NatsStream, log, []string{"kafka:9092"})
	a.container.AddScoped(constants.DatabaseTransactionKey, func(c di.Container) (any, error) {
		return dbConn.Begin()
	})
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

	a.container.AddScoped(constants.UsersRepoKey, func(c di.Container) (any, error) {
		db := postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB))
		return pg.NewUserRepository(constants.UsersTableName, db), nil
	})

	a.container.AddScoped(constants.ApplicationKey, func(c di.Container) (any, error) {
		users := c.Get(constants.UsersRepoKey).(domain.UserRepository)
		domainDispatcher := c.Get(constants.DomainDispatcherKey).(*ddd.EventDispatcher[ddd.AggregateEvent])
		application := app.NewInstrumentedApp(
			app.New(users, domainDispatcher),
		)
		//application := &app.Application{
		//	Commands: app.Commands{
		//		RegisterUser: commands.NewRegisterUserHandler(repo, a.logger, domainDispatcher),
		//		EnableUser:   commands.NewEnableUserHandler(repo, a.logger, domainDispatcher),
		//	},
		//	Queries: app.Queries{
		//		AuthorizeUser: queries.NewAuthorizeUserHandler(repo, a.logger, domainDispatcher),
		//	},
		//}
		return application, nil
	})
	a.container.AddScoped(constants.DomainEventHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewDomainEventHandlers(c.Get(constants.EventPublisherKey).(am.EventPublisher)), nil
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
