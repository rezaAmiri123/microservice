package app

import "github.com/rezaAmiri123/microservice/ordering/internal/app/commands"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Queries struct {
}

type Commands struct {
	CreateOrder   *commands.CreateOrderHandler
	ReadyOrder    *commands.ReadyOrderHandler
	ApproveOrder  *commands.ApproveOrderHandler
	CompleteOrder *commands.CompleteOrderHandler
}
