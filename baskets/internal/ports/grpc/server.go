package grpc

import (
	"github.com/rezaAmiri123/microservice/baskets/basketspb"
	"github.com/rezaAmiri123/microservice/baskets/internal/app"
	"github.com/rezaAmiri123/microservice/baskets/internal/domain"
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
		basketspb.UnimplementedBasketServiceServer
	}
)

var _ basketspb.BasketServiceServer = (*server)(nil)

func RegisterServer(application app.App, registrar grpc.ServiceRegistrar, logger logger.Logger) error {
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

func (s server) basketFromDomain(basket *domain.Basket) *basketspb.Basket {
	protoBasket := &basketspb.Basket{
		Id: basket.ID(),
	}

	protoBasket.Items = make([]*basketspb.Item, 0, len(basket.Items))

	for _, item := range basket.Items {
		protoBasket.Items = append(protoBasket.Items, &basketspb.Item{
			StoreId:      item.StoreID,
			StoreName:    item.StoreName,
			ProductId:    item.ProductID,
			ProductName:  item.ProductName,
			ProductPrice: item.ProductPrice,
			Quantity:     int32(item.Quantity),
		})
	}

	return protoBasket
}
