package app

import (
	command "github.com/rezaAmiri123/microservice/service_api/internal/app/commands"
	query "github.com/rezaAmiri123/microservice/service_api/internal/app/queries"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Queries struct {
	GetTransfers *query.GetTransfersHandler
}

type Commands struct {
	CreateUser    *command.CreateUserHandler
	CreateAccount *command.CreateAccountHandler
	Login         *command.LoginHandler
	LoginVerify   *command.LoginVerifyHandler

	// Transfer
	CreateTransfer *command.CreateTransferHandler
}
