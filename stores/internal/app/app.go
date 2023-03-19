package app

import (
	"github.com/rezaAmiri123/microservice/stores/internal/app/commands"
	"github.com/rezaAmiri123/microservice/stores/internal/app/queries"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Queries struct {
	GetProduct *queries.GetProductHandler
	GetStore   *queries.GetStoreHandler
}

type Commands struct {
	AddProduct  *commands.AddProductHandler
	CreateStore *commands.CreateStoreHandler
}
