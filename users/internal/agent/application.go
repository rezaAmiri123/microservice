package agent

import (
	"database/sql"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/jetstream"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/tm"
	"github.com/rezaAmiri123/microservice/users/internal/adapters/pg"
	"github.com/rezaAmiri123/microservice/users/internal/app"
	"github.com/rezaAmiri123/microservice/users/internal/app/commands"
	"github.com/rezaAmiri123/microservice/users/internal/handlers"
	"github.com/rezaAmiri123/microservice/users/userspb"
)

func (a *Agent) setupApplication() error {
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
	a.container.AddSingleton("db", func(c di.Container) (any, error) {
		return dbConn, nil
	})

	//repo, err := adapters.NewGORMArticleRepository(a.DBConfig)
	repo := pg.NewPGUserRepository(dbConn, a.logger)

	// setup Driven adapters
	reg := registry.New()
	a.container.AddSingleton("registry", func(c di.Container) (any, error) {
		return reg, nil
	})
	if err = userspb.Registrations(reg); err != nil {
		return err
	}
	js, err := a.nats()
	if err != nil {
		return err
	}
	stream := jetstream.NewStream("stream", js, a.logger)
	a.container.AddSingleton("stream", func(c di.Container) (any, error) {
		return stream, nil
	})
	eventStream := am.NewEventStream(reg, stream)
	//commandStream := am.NewCommandStream(reg, stream)
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	a.container.AddSingleton("domainDispatcher", func(c di.Container) (any, error) {
		return domainDispatcher, nil
	})
	a.container.AddSingleton("outboxProcessor", func(c di.Container) (any, error) {
		return tm.NewOutboxProcessor(
			c.Get("stream").(am.RawMessageStream),
			postgres.NewOutboxStore("users.outbox", c.Get("db").(*sql.DB)),
		), nil
	})
	a.container.AddScoped("tx", func(c di.Container) (any, error) {
		db := c.Get("db").(*sql.DB)
		return db.Begin()
	})
	a.container.AddScoped("users", func(c di.Container) (any, error) {
		return pg.NewPGUserRepository(c.Get("tx").(*sql.Tx), a.logger), nil
	})

	a.container.AddScoped("txStream", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		outboxStore := postgres.NewOutboxStore("userss.outbox", tx)
		return am.RawMessageStreamWithMiddleware(
			c.Get("stream").(am.RawMessageStream),
			tm.NewOutboxStreamMiddleware(outboxStore),
		), nil
	})
	a.container.AddScoped("eventStream", func(c di.Container) (any, error) {
		return am.NewEventStream(c.Get("registry").(registry.Registry), c.Get("txStream").(am.RawMessageStream)), nil
	})
	a.container.AddScoped("replyStream", func(c di.Container) (any, error) {
		return am.NewReplyStream(c.Get("registry").(registry.Registry), c.Get("txStream").(am.RawMessageStream)), nil
	})
	a.container.AddScoped("inboxMiddleware", func(c di.Container) (any, error) {
		tx := c.Get("tx").(*sql.Tx)
		inboxStore := postgres.NewInboxStore("users.inbox", tx)
		return tm.NewInboxHandlerMiddleware(inboxStore), nil
	})
	domainEventHandlers := handlers.NewDomainEventHandlers(eventStream)
	//integrationEventHandlers := commands.NewIntegrationEventHandlers(eventStream)
	//handlers.RegisterDomainEventHandlers(domainEventHandlers, domainDispatcher)

	application := &app.Application{
		Commands: app.Commands{
			RegisterUser: commands.NewRegisterUserHandler(repo, a.logger, domainDispatcher),
			EnableUser:   commands.NewEnableUserHandler(repo, a.logger, domainDispatcher),
		},
		Queries: app.Queries{},
	}
	a.Application = application

	a.container.AddScoped("app", func(c di.Container) (any, error) {
		return application, nil
	})
	a.container.AddScoped("domainEventHandlers", func(c di.Container) (any, error) {
		return domainEventHandlers, nil
	})
	commandHandler := handlers.NewCommandHandlers(application)
	a.container.AddScoped("commandHandlers", func(c di.Container) (any, error) {
		return commandHandler, nil
	})
	handlers.RegisterDomainEventHandlersTx(a.container)
	if err = handlers.RegisterCommandHandlersTx(a.container); err != nil {
		// TODO command handlers not working
		//return err
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
