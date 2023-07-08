package agent

import (
	"database/sql"

	"github.com/rezaAmiri123/microservice/pkg/db/postgresotel"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/users/internal/adapters/pg"
	"github.com/rezaAmiri123/microservice/users/internal/app"
	"github.com/rezaAmiri123/microservice/users/internal/constants"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
)

func (a *Agent) setupApplication() error {
	//dbConn, err := postgres.NewDB(postgres.Config{
	//	PGDriver:     a.PGDriver,
	//	PGHost:       a.PGHost,
	//	PGPort:       a.PGPort,
	//	PGUser:       a.PGUser,
	//	PGDBName:     a.PGDBName,
	//	PGPassword:   a.PGPassword,
	//	PGSearchPath: a.PGSearchPath,
	//})
	//if err != nil {
	//	return fmt.Errorf("cannot load db: %w", err)
	//}
	//
	//if err = postgres.MigrateUp(dbConn, migrations.FS); err != nil {
	//	return err
	//}
	//
	//a.container.AddSingleton(constants.DatabaseKey, func(c di.Container) (any, error) {
	//	return dbConn, nil
	//})
	//a.container.AddScoped(constants.DatabaseTransactionKey, func(c di.Container) (any, error) {
	//	return dbConn.Begin()
	//})

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
		return application, nil
	})

	return nil
}
