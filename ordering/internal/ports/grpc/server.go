package grpc

import (
	"database/sql"
	"github.com/rezaAmiri123/microservice/ordering/internal/app"
	"github.com/rezaAmiri123/microservice/ordering/internal/constants"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
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
		orderingpb.UnimplementedOrderingServiceServer
	}
	serverTx struct {
		c di.Container
		orderingpb.UnimplementedOrderingServiceServer
	}
)

var _ orderingpb.OrderingServiceServer = (*server)(nil)
var _ orderingpb.OrderingServiceServer = (*serverTx)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	orderingpb.RegisterOrderingServiceServer(registrar, serverTx{
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
func NewServer(app *app.Application, logger logger.Logger) *server {
	cfg := &Config{
		App:    app,
		Logger: logger,
	}
	return &server{cfg: cfg}
}

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
