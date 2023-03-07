package app

import "github.com/rezaAmiri123/microservice/stores/internal/app/commands"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Queries struct {
}

type Commands struct {
	AddProduct  *commands.AddProductHandler
	CreateStore *commands.CreateStoreHandler
}
