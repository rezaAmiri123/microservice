package agent

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/jetstream"
	"github.com/rezaAmiri123/microservice/pkg/registry"
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

	//repo, err := adapters.NewGORMArticleRepository(a.DBConfig)
	repo := pg.NewPGUserRepository(dbConn, a.logger)

	// setup Driven adapters
	reg := registry.New()
	if err = userspb.Registrations(reg); err != nil {
		return err
	}
	js, err := a.nats()
	if err != nil {
		return err
	}
	stream := jetstream.NewStream("stream", js, a.logger)
	eventStream := am.NewEventStream(reg, stream)
	commandStream := am.NewCommandStream(reg, stream)
	domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()
	domainEventHandlers := handlers.NewDomainEventHandlers(eventStream)
	//integrationEventHandlers := commands.NewIntegrationEventHandlers(eventStream)
	handlers.RegisterDomainEventHandlers(domainEventHandlers, domainDispatcher)

	application := &app.Application{
		Commands: app.Commands{
			RegisterUser: commands.NewRegisterUserHandler(repo, a.logger, domainDispatcher),
			EnableUser:   commands.NewEnableUserHandler(repo, a.logger, domainDispatcher),
		},
		Queries: app.Queries{},
	}
	a.Application = application

	commandHandler := handlers.NewCommandHandlers(application)
	if err = handlers.RegisterCommandHandlers(commandStream, commandHandler); err != nil {
		// TODO there is a problem
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
