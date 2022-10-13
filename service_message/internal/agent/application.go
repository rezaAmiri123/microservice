package agent

import (
	"fmt"

	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/service_message/internal/adapters/pg"
	"github.com/rezaAmiri123/microservice/service_message/internal/app"
	"github.com/rezaAmiri123/microservice/service_message/internal/app/command"
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
	repo := pg.NewPGMessageRepository(dbConn, a.logger, a.metric)

	application := &app.Application{
		Commands: app.Commands{
			CreateEmail: command.NewCreateEmailHandler(repo, a.logger),
		},
		Queries: app.Queries{},
	}
	a.Application = application
	return nil
}
