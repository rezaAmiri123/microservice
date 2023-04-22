package app

import (
	"context"
	"github.com/rezaAmiri123/microservice/payments/internal/app/commands"
	"github.com/rezaAmiri123/microservice/payments/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
)

//type Application struct {
//	Commands Commands
//	Queries  Queries
//}
//
//type Queries struct{}
//
//type Commands struct {
//	CreateInvoice    *commands.CreateInvoiceHandler
//	AuthorizePayment *commands.AuthorizePaymentHandler
//	ConfirmPayment   *commands.ConfirmPaymentHandler
//	PayInvoice       *commands.PayInvoiceHandler
//	CancelInvoice    *commands.CancelInvoiceHandler
//}

type (
	App interface {
		Commands
		Queries
	}
	Commands interface {
		CreateInvoice(ctx context.Context, cmd commands.CreateInvoice) error
		AuthorizePayment(ctx context.Context, cmd commands.AuthorizePayment) error
		ConfirmPayment(ctx context.Context, cmd commands.ConfirmPayment) error
		PayInvoice(ctx context.Context, cmd commands.PayInvoice) error
		CancelInvoice(ctx context.Context, cmd commands.CancelInvoice) error
	}
	Queries interface {
	}
	Application struct {
		appCommands
		appQueries
	}
	appCommands struct {
		commands.CreateInvoiceHandler
		commands.AuthorizePaymentHandler
		commands.ConfirmPaymentHandler
		commands.PayInvoiceHandler
		commands.CancelInvoiceHandler
	}
	appQueries struct {
	}
)

var _ App = (*Application)(nil)

func New(
	invoices domain.InvoiceRepository,
	payments domain.PaymentRepository,
	publisher ddd.EventPublisher[ddd.Event],
	log logger.Logger,
) Application {
	return Application{
		appCommands: appCommands{
			CreateInvoiceHandler:    commands.NewCreateInvoiceHandler(invoices, log),
			AuthorizePaymentHandler: commands.NewAuthorizePaymentHandler(payments, log),
			ConfirmPaymentHandler:   commands.NewConfirmPaymentHandler(payments, log),
			PayInvoiceHandler:       commands.NewPayInvoiceHandler(invoices, publisher, log),
			CancelInvoiceHandler:    commands.NewCancelInvoiceHandler(invoices, log),
		},
		appQueries: appQueries{},
	}
}
