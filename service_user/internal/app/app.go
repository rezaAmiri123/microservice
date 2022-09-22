package app

import "github.com/rezaAmiri123/microservice/service_user/internal/app/command"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Queries struct {
}

type Commands struct {
	CreateUser *command.CreateUserHandler
	Login      *command.LoginHandler
}
