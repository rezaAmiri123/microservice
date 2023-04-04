package app

import (
	"github.com/rezaAmiri123/microservice/users/internal/app/commands"
	"github.com/rezaAmiri123/microservice/users/internal/app/queries"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Queries struct {
	AuthorizeUser *queries.AuthorizeUserHandler
}

type Commands struct {
	RegisterUser *commands.RegisterUserHandler
	EnableUser   *commands.EnableUserHandler
}
