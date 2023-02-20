package agent

import (
	"fmt"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/users/internal/adapters/pg"
	"github.com/rezaAmiri123/microservice/users/internal/app"
	"github.com/rezaAmiri123/microservice/users/internal/app/commands"
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
	domainDispatcher := ddd.NewEventDispatcher()

	application := &app.Application{
		Commands: app.Commands{
			RegisterUser: commands.NewRegisterUserHandler(repo, a.logger, domainDispatcher),
			EnableUser:   commands.NewEnableUserHandler(repo, a.logger, domainDispatcher),
		},
		Queries: app.Queries{},
	}
	a.Application = application

	return nil
}
