package app

import "github.com/rezaAmiri123/microservice/payments/internal/app/commands"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Queries struct{}

type Commands struct {
	CreateInvoice    *commands.CreateInvoiceHandler
	AuthorizePayment *commands.AuthorizePaymentHandler
	ConfirmPayment   *commands.ConfirmPaymentHandler
	PayInvoice       *commands.PayInvoiceHandler
	CancelInvoice    *commands.CancelInvoiceHandler
}
