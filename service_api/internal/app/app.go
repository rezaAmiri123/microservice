package app

import command "github.com/rezaAmiri123/microservice/service_api/internal/app/commands"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Queries struct {
}

type Commands struct {
	CreateUser    *command.CreateUserHandler
	CreateAccount *command.CreateAccountHandler
	Login         *command.LoginHandler
	LoginVerify   *command.LoginVerifyHandler

	// Transfer
	CreateTransfer *command.CreateTransferHandler
}
