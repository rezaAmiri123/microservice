package grpc

import (
	"database/sql"
	"fmt"
	"github.com/rezaAmiri123/microservice/baskets/basketspb"
	"github.com/rezaAmiri123/microservice/baskets/internal/app"
	"github.com/rezaAmiri123/microservice/baskets/internal/constants"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"google.golang.org/grpc"
)

type (
	Config struct {
		App    *app.Application
		Logger logger.Logger
	}
	server struct {
		cfg *Config
		basketspb.UnimplementedBasketServiceServer
	}
	serverTx struct {
		c di.Container
		basketspb.UnimplementedBasketServiceServer
	}
)

var _ basketspb.BasketServiceServer = (*server)(nil)
var _ basketspb.BasketServiceServer = (*serverTx)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	basketspb.RegisterBasketServiceServer(registrar, serverTx{
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

func RegisterServer(application *app.Application, registrar grpc.ServiceRegistrar, logger logger.Logger) error {
	cfg := &Config{
		App:    application,
		Logger: logger,
	}
	basketspb.RegisterBasketServiceServer(registrar, server{cfg: cfg})
	return nil
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
		fmt.Println("rollback")
		panic(p)
	} else if err != nil {
		fmt.Println("rollback")
		_ = tx.Rollback()
		return err
	} else {
		fmt.Println("commit")
		return tx.Commit()
	}
}
