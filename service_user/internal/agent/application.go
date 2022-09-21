package agent

import (
	"fmt"

	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/service_user/internal/adapters/pg"
	"github.com/rezaAmiri123/microservice/service_user/internal/app"
	"github.com/rezaAmiri123/microservice/service_user/internal/app/command"
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
	repo := pg.NewPGUserRepository(dbConn, a.logger, a.metric)

	application := &app.Application{
		Commands: app.Commands{
			CreateUser: command.NewCreateUserHandler(repo, a.logger),
		},
		Queries: app.Queries{},
	}
	a.Application = application
	return nil
}