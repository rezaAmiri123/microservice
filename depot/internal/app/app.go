package app

import (
	"github.com/rezaAmiri123/microservice/depot/internal/app/commands"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Queries struct {
}

type Commands struct {
	CreateShoppingList   *commands.CreateShoppingListHandler
	InitiateShopping     *commands.InitiateShoppingHandler
	AssignShoppingList   *commands.AssignShoppingListHandler
	CompleteShoppingList *commands.CompleteShoppingListHandler
}
