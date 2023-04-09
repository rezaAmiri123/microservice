package app

import "github.com/rezaAmiri123/microservice/baskets/internal/app/commands"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Queries struct {
}

type Commands struct {
	StartBasket    *commands.StartBasketHandler
	AddItem        *commands.AddItemHandler
	CheckoutBasket *commands.CheckoutBasketHandler
	CancelBasket   *commands.CancelBasketHandler
}
