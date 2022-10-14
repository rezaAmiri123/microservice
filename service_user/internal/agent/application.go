package agent

import (
	"fmt"

	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	kafkaClient "github.com/rezaAmiri123/microservice/pkg/kafka"
	"github.com/rezaAmiri123/microservice/pkg/token"
	"github.com/rezaAmiri123/microservice/pkg/token/jwt"
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

	maker, err := jwt.NewJWTMaker(a.SecretKey)
	if err != nil {
		return fmt.Errorf("cannot make jwt maker: %w", err)
	}
	makerConfig := token.Config{
		AccessTokenDuration:  a.AccessTokenDuration,
		RefreshTokenDuration: a.RefreshTokenDuration,
	}

	producer := kafkaClient.NewProducer(a.logger, a.KafkaBrokers)

	application := &app.Application{
		Commands: app.Commands{
			CreateUser: command.NewCreateUserHandler(repo, a.logger, producer),
			UpdateUser: command.NewUpdateUserHandler(repo, a.logger),
			Login:      command.NewLoginHandler(repo, a.logger, maker, makerConfig),
		},
		Queries: app.Queries{},
	}
	a.Application = application
	a.Maker = maker
	return nil
}
