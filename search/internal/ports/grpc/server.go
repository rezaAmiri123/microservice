package grpc

import (
	"database/sql"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/search/internal/app"
	"github.com/rezaAmiri123/microservice/search/internal/constants"
	"github.com/rezaAmiri123/microservice/search/internal/domain"
	"github.com/rezaAmiri123/microservice/search/searchpb"
	"google.golang.org/grpc"
)

type (
	Config struct {
		App    app.App
		Logger logger.Logger
	}
	server struct {
		cfg *Config
		searchpb.UnimplementedSearchServiceServer
	}
	serverTx struct {
		c di.Container
		searchpb.UnimplementedSearchServiceServer
	}
)

var _ searchpb.SearchServiceServer = (*server)(nil)
var _ searchpb.SearchServiceServer = (*serverTx)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	searchpb.RegisterSearchServiceServer(registrar, serverTx{
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

func RegisterServer(application app.App, registrar grpc.ServiceRegistrar, logger logger.Logger) error {
	cfg := &Config{
		App:    application,
		Logger: logger,
	}
	searchpb.RegisterSearchServiceServer(registrar, server{cfg: cfg})
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
		panic(p)
	} else if err != nil {
		_ = tx.Rollback()
		return err
	} else {
		return tx.Commit()
	}
}

func (s server) orderFromDomain(order *domain.Order) *searchpb.Order {
	items := make([]*searchpb.Item, len(order.Items))
	for i, item := range order.Items {
		items[i] = s.itemFromDomain(item)
	}
	return &searchpb.Order{
		OrderId:  order.OrderID,
		UserId:   order.UserID,
		Username: order.Username,
		Items:    items,
		Status:   order.Status,
		Total:    order.Total,
	}
}
func (s server) itemFromDomain(item domain.Item) *searchpb.Item {
	return &searchpb.Item{
		StoreId:     item.StoreID,
		ProductId:   item.ProductID,
		StoreName:   item.StoreName,
		ProductName: item.ProductName,
		Price:       item.Price,
		Quantity:    int64(item.Quantity),
	}
}
