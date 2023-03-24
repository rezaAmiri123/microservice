package commands

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/payments/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
)

type (
	CreateInvoice struct {
		ID        string
		OrderID   string
		PaymentID string
		Amount    float64
	}

	CreateInvoiceHandler struct {
		invoices  domain.InvoiceRepository
		payments  domain.PaymentRepository
		publisher ddd.EventPublisher[ddd.Event]
		logger    logger.Logger
	}
)

func NewCreateInvoiceHandler(
	invoices domain.InvoiceRepository,
	//payments domain.PaymentRepository,
	//publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) *CreateInvoiceHandler {
	return &CreateInvoiceHandler{
		invoices: invoices,
		//payments:  payments,
		//publisher: publisher,
		logger: logger,
	}
}

func (h CreateInvoiceHandler) Handle(ctx context.Context, cmd CreateInvoice) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateInvoiceHandler.Handle")
	defer span.Finish()

	return h.invoices.Save(ctx, &domain.Invoice{
		ID:      cmd.ID,
		OrderID: cmd.OrderID,
		Amount:  cmd.Amount,
		Status:  domain.InvoiceIsPending,
	})
}
