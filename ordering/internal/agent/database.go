package agent

import (
	"fmt"
	"github.com/rezaAmiri123/microservice/ordering/internal/constants"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/pkg/di"
)

func (a *Agent) setupDatabase() error {
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

	a.container.AddSingleton(constants.DatabaseKey, func(c di.Container) (any, error) {
		return dbConn, nil
	})

	a.container.AddScoped(constants.DatabaseTransactionKey, func(c di.Container) (any, error) {
		return dbConn.Begin()
	})

	return nil
}
