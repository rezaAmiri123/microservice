package agent

import (
	"fmt"

	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/service_finance/internal/adapters/pg"
	"github.com/rezaAmiri123/microservice/service_finance/internal/app"
	"github.com/rezaAmiri123/microservice/service_finance/internal/app/commands"
	"github.com/rezaAmiri123/microservice/service_finance/internal/app/queries"
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
	repo := pg.NewPGFinanceRepository(dbConn, a.logger, a.metric)

	application := &app.Application{
		Commands: app.Commands{
			CreateAccount:  commands.NewCreateAccountHandler(repo, a.logger),
			CreateTransfer: commands.NewCreateTransferHandler(repo, a.logger),
		},
		Queries: app.Queries{
			GetAccountByID: queries.NewGetAccountByIDHandler(repo, a.logger),
			GetTransfers:   queries.NewGetTransfersHandler(repo, a.logger),
		},
	}
	a.Application = application
	return nil
}
