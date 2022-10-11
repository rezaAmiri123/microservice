package app

import "github.com/rezaAmiri123/microservice/service_message/internal/app/command"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Queries struct {
}

type Commands struct {
	CreateEmail *command.CreateEmailHandler
}
