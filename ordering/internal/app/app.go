package app

import (
	"context"
	"github.com/rezaAmiri123/microservice/ordering/internal/app/commands"
	"github.com/rezaAmiri123/microservice/ordering/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
)

//type Application struct {
//	Commands Commands
//	Queries  Queries
//}
//
//type Queries struct {
//}
//
//type Commands struct {
//	CreateOrder   *commands.CreateOrderHandler
//	ReadyOrder    *commands.ReadyOrderHandler
//	ApproveOrder  *commands.ApproveOrderHandler
//	CompleteOrder *commands.CompleteOrderHandler
//	CancelOrder   *commands.CancelOrderHandler
//}

type (
	App interface {
		Commands
		Queries
	}

	Commands interface {
		CreateOrder(ctx context.Context, cmd commands.CreateOrder) error
		ReadyOrder(ctx context.Context, cmd commands.ReadyOrder) error
		ApproveOrder(ctx context.Context, cmd commands.ApproveOrder) error
		CompleteOrder(ctx context.Context, cmd commands.CompleteOrder) error
		CancelOrder(ctx context.Context, cmd commands.CancelOrder) error
	}
	Queries interface {
	}
	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateOrderHandler
		commands.ReadyOrderHandler
		commands.ApproveOrderHandler
		commands.CompleteOrderHandler
		commands.CancelOrderHandler
	}
	appQueries struct {
	}
)

var _ App = (*Application)(nil)

func New(
	order domain.OrderRepository,
	publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) Application {
	return Application{
		appCommands: appCommands{
			CreateOrderHandler:   commands.NewCreateOrderHandler(order, publisher, logger),
			ReadyOrderHandler:    commands.NewReadyOrderHandler(order, publisher, logger),
			ApproveOrderHandler:  commands.NewApproveOrderHandler(order, publisher, logger),
			CompleteOrderHandler: commands.NewCompleteOrderHandler(order, publisher, logger),
			CancelOrderHandler:   commands.NewCancelOrderHandler(order, publisher, logger),
		},
	}
}
