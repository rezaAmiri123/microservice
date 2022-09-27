package app

import "github.com/rezaAmiri123/microservice/service_finance/internal/app/commands"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Queries struct {
}

type Commands struct {
	CreateAccount *commands.CreateAccountHandler
}
