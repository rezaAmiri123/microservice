package app

import (
	"context"
	"github.com/rezaAmiri123/microservice/depot/internal/app/commands"
	"github.com/rezaAmiri123/microservice/depot/internal/domain"
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
//	CreateShoppingList   *commands.CreateShoppingListHandler
//	InitiateShopping     *commands.InitiateShoppingHandler
//	AssignShoppingList   *commands.AssignShoppingListHandler
//	CompleteShoppingList *commands.CompleteShoppingListHandler
//}

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		CreateShoppingList(ctx context.Context, cmd commands.CreateShoppingList) error
		InitiateShopping(ctx context.Context, cmd commands.InitiateShopping) error
		AssignShoppingList(ctx context.Context, cmd commands.AssignShoppingList) error
		CompleteShoppingList(ctx context.Context, cmd commands.CompleteShoppingList) error
	}
	Queries interface {
	}
	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateShoppingListHandler
		commands.InitiateShoppingHandler
		commands.AssignShoppingListHandler
		commands.CompleteShoppingListHandler
	}
	appQueries struct {
	}
)

var _ App = (*Application)(nil)

func New(
	shoppingLists domain.ShoppingListRepository,
	stores domain.StoreRepository,
	products domain.ProductRepository,
	domainPublisher ddd.EventPublisher[ddd.AggregateEvent],
	log logger.Logger,
) Application {
	return Application{
		appCommands: appCommands{
			CreateShoppingListHandler:   commands.NewCreateShoppingListHandler(shoppingLists, stores, products, domainPublisher, log),
			InitiateShoppingHandler:     commands.NewInitiateShoppingHandler(shoppingLists, domainPublisher, log),
			AssignShoppingListHandler:   commands.NewAssignShoppingListHandler(shoppingLists, domainPublisher, log),
			CompleteShoppingListHandler: commands.NewCompleteShoppingListHandler(shoppingLists, domainPublisher, log),
		},
		appQueries: appQueries{},
	}
}
