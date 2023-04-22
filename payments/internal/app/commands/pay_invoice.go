package commands

import (
	"context"
	"github.com/rezaAmiri123/microservice/payments/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/stackus/errors"
)

type (
	PayInvoice struct {
		ID string
	}

	PayInvoiceHandler struct {
		invoices  domain.InvoiceRepository
		publisher ddd.EventPublisher[ddd.Event]
		logger    logger.Logger
	}
)

func NewPayInvoiceHandler(
	invoices domain.InvoiceRepository,
	publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) PayInvoiceHandler {
	return PayInvoiceHandler{
		invoices:  invoices,
		publisher: publisher,
		logger:    logger,
	}
}

func (h PayInvoiceHandler) PayInvoice(ctx context.Context, cmd PayInvoice) error {
	//span, ctx := opentracing.StartSpanFromContext(ctx, "PayInvoiceHandler.Handle")
	//defer span.Finish()

	invoice, err := h.invoices.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if invoice.Status != domain.InvoiceIsPending {
		return errors.Wrap(errors.ErrBadRequest, "invoice cannot be paid for")
	}

	invoice.Status = domain.InvoiceIsPaid

	// Before or after the invoice is saved we still risk something failing which
	// will leave the state change only partially complete
	if err = h.publisher.Publish(ctx, ddd.NewEvent(domain.InvoicePaidEvent, &domain.InvoicePaid{
		ID:      invoice.ID,
		OrderID: invoice.OrderID,
	})); err != nil {
		return err
	}
	return h.invoices.Update(ctx, invoice)
}
