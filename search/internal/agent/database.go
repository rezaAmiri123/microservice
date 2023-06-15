package agent

import (
	"context"
	"fmt"
	"github.com/rezaAmiri123/microservice/pkg/db/mongodb"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/search/internal/constants"
	"github.com/stackus/errors"
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

	return a.setMongoDB()
}

func (a *Agent) setMongoDB() error {
	mongoDBConn, err := mongodb.NewMongoDBConn(context.Background(), &mongodb.Config{
		URI:      a.MongoURI,
		User:     a.MongoUser,
		Password: a.MongoPassword,
	})
	if err != nil {
		return errors.Wrap(err, "NewMongoDBConn")
	}
	//s.mongoClient = mongoDBConn
	//defer mongoDBConn.Disconnect(ctx) // nolint: errcheck
	//s.log.Infof("Mongo connected: %v", mongoDBConn.NumberSessionsInProgress())

	a.container.AddSingleton(constants.MongoDBKey, func(c di.Container) (any, error) {
		return mongoDBConn, nil
	})

	return nil

}
