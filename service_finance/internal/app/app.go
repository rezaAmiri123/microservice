package app

import (
	"github.com/rezaAmiri123/microservice/service_finance/internal/app/commands"
	"github.com/rezaAmiri123/microservice/service_finance/internal/app/queries"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Queries struct {
	GetAccountByID *queries.GetAccountByIDHandler
	GetTransfers   *queries.GetTransfersHandler
}

type Commands struct {
	CreateAccount  *commands.CreateAccountHandler
	CreateTransfer *commands.CreateTransferHandler
}
