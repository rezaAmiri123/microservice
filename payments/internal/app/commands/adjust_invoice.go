package commands

import (
	"context"
	"github.com/rezaAmiri123/microservice/payments/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/logger"
)

type (
	AdjustInvoice struct {
		ID     string
		Amount float64
	}

	AdjustInvoiceHandler struct {
		invoices domain.InvoiceRepository
		//publisher ddd.EventPublisher[ddd.Event]
		logger logger.Logger
	}
)

func NewAdjustInvoiceHandler(
	invoices domain.InvoiceRepository,
	//publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) AdjustInvoiceHandler {
	return AdjustInvoiceHandler{
		invoices: invoices,
		//publisher: publisher,
		logger: logger,
	}
}

func (h AdjustInvoiceHandler) AdjustInvoice(ctx context.Context, cmd AdjustInvoice) error {
	//span, ctx := opentracing.StartSpanFromContext(ctx, "PayInvoiceHandler.Handle")
	//defer span.Finish()

	invoice, err := h.invoices.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	invoice.Amount = cmd.Amount

	return h.invoices.Update(ctx, invoice)
}
