package app

import "github.com/rezaAmiri123/microservice/users/internal/app/commands"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Queries struct {
}

type Commands struct {
	RegisterUser *commands.RegisterUserHandler
	EnableUser   *commands.EnableUserHandler
}
