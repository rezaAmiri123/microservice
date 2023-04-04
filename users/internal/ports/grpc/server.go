package grpc

import (
	"database/sql"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/users/internal/app"
	"github.com/rezaAmiri123/microservice/users/internal/constants"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"google.golang.org/grpc"
)

type (
	Config struct {
		App    *app.Application
		Logger logger.Logger
	}
	server struct {
		cfg *Config
		userspb.UnimplementedUserServiceServer
	}
	serverTx struct {
		c di.Container
		userspb.UnimplementedUserServiceServer
	}
)

var _ userspb.UserServiceServer = (*server)(nil)
var _ userspb.UserServiceServer = (*serverTx)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	userspb.RegisterUserServiceServer(registrar, serverTx{
		c: container,
	})
	return nil
}
func (s serverTx) getNextServer() server {
	cfg := &Config{
		App:    s.c.Get(constants.ApplicationKey).(*app.Application),
		Logger: s.c.Get(constants.LoggerKey).(logger.Logger),
	}
	return server{cfg: cfg}
}

//func NewServer(app *app.Application, logger logger.Logger) *server {
//	cfg := &Config{
//		App:    app,
//		Logger: logger,
//	}
//	return &server{cfg: cfg}
//}

func (s serverTx) closeTx(tx *sql.Tx, err error) error {
	if p := recover(); p != nil {
		_ = tx.Rollback()
		panic(p)
	} else if err != nil {
		_ = tx.Rollback()
		return err
	} else {
		return tx.Commit()
	}
}
