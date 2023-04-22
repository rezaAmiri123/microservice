package grpc

import (
	"database/sql"
	"github.com/rezaAmiri123/microservice/depot/depotpb"
	"github.com/rezaAmiri123/microservice/depot/internal/app"
	"github.com/rezaAmiri123/microservice/depot/internal/app/commands"
	"github.com/rezaAmiri123/microservice/depot/internal/constants"
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
		depotpb.UnimplementedDepotServiceServer
	}
	serverTx struct {
		c di.Container
		depotpb.UnimplementedDepotServiceServer
	}
)

var _ depotpb.DepotServiceServer = (*server)(nil)
var _ depotpb.DepotServiceServer = (*serverTx)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	depotpb.RegisterDepotServiceServer(registrar, serverTx{
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

func (s server) itemToDomain(item *depotpb.OrderItem) commands.OrderItem {
	return commands.OrderItem{
		StoreID:   item.GetStoreId(),
		ProductID: item.GetProductId(),
		Quantity:  int(item.GetQuantity()),
	}
}
