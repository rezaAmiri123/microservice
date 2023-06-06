package grpc

import (
	"github.com/rezaAmiri123/microservice/ordering/internal/app"
	"github.com/rezaAmiri123/microservice/ordering/internal/domain"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
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
		orderingpb.UnimplementedOrderingServiceServer
	}
)

var _ orderingpb.OrderingServiceServer = (*server)(nil)

//func NewServer(app app.App, logger logger.Logger) *server {
//	cfg := &Config{
//		App:    app,
//		Logger: logger,
//	}
//	return &server{cfg: cfg}
//}

func RegisterServer(application app.App, registrar grpc.ServiceRegistrar, logger logger.Logger) error {
	cfg := &Config{
		App:    application,
		Logger: logger,
	}
	orderingpb.RegisterOrderingServiceServer(registrar, server{cfg: cfg})
	return nil
}

func (s server) orderFromDomain(order *domain.Order) *orderingpb.Order {
	items := make([]*orderingpb.Item, len(order.Items))
	for i, item := range order.Items {
		items[i] = s.itemFromDomain(item)
	}
	return &orderingpb.Order{
		Id:        order.ID(),
		UserId:    order.UserID,
		PaymentId: order.PaymentID,
		Items:     items,
		Status:    order.Status.String(),
	}
}
