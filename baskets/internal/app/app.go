package app

import (
	"context"
	"github.com/rezaAmiri123/microservice/baskets/internal/app/commands"
	"github.com/rezaAmiri123/microservice/baskets/internal/app/queries"
	"github.com/rezaAmiri123/microservice/baskets/internal/domain"
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
//	StartBasket    *commands.StartBasketHandler
//	AddItem        *commands.AddItemHandler
//	CheckoutBasket *commands.CheckoutBasketHandler
//	CancelBasket   *commands.CancelBasketHandler
//}

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		StartBasket(ctx context.Context, start commands.StartBasket) error
		CheckoutBasket(ctx context.Context, checkout commands.CheckoutBasket) error
		AddItem(ctx context.Context, add commands.AddItem) error
		CancelBasket(ctx context.Context, cancel commands.CancelBasket) error
	}
	Queries interface {
		GetBasket(ctx context.Context, cmd queries.GetBasket) (*domain.Basket, error)
	}
	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.StartBasketHandler
		commands.CheckoutBasketHandler
		commands.AddItemHandler
		commands.CancelBasketHandler
	}
	appQueries struct {
		queries.GetBasketHandler
	}
)

var _ App = (*Application)(nil)

func New(
	baskets domain.BasketRepository,
	stores domain.StoreRepository,
	products domain.ProductRepository,
	publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) Application {
	return Application{
		appCommands: appCommands{
			StartBasketHandler:    commands.NewStartBasketHandler(baskets, publisher, logger),
			CheckoutBasketHandler: commands.NewCheckoutBasketHandler(baskets, publisher, logger),
			AddItemHandler:        commands.NewAddItemHandler(baskets, stores, products, publisher, logger),
			CancelBasketHandler:   commands.NewCancelBasketHandler(baskets, publisher, logger),
		},
		appQueries: appQueries{
			GetBasketHandler: queries.NewGetBasketHandler(baskets, logger),
		},
	}
}
