package grpc

import (
	"database/sql"
	"github.com/rezaAmiri123/microservice/payments/internal/app"
	"github.com/rezaAmiri123/microservice/payments/internal/constants"
	"github.com/rezaAmiri123/microservice/payments/paymentspb"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"google.golang.org/grpc"
)

type (
	Config struct {
		App    app.App
		Logger logger.Logger
	}
	server struct {
		cfg *Config
		paymentspb.UnimplementedPaymentsServiceServer
	}
	serverTx struct {
		c di.Container
		paymentspb.UnimplementedPaymentsServiceServer
	}
)

var _ paymentspb.PaymentsServiceServer = (*server)(nil)
var _ paymentspb.PaymentsServiceServer = (*serverTx)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	paymentspb.RegisterPaymentsServiceServer(registrar, serverTx{
		c: container,
	})
	return nil
}
func (s serverTx) getNextServer() server {
	cfg := &Config{
		App:    s.c.Get(constants.ApplicationKey).(app.App),
		Logger: s.c.Get(constants.LoggerKey).(logger.Logger),
	}
	return server{cfg: cfg}
}
func NewServer(app app.App, logger logger.Logger) *server {
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
