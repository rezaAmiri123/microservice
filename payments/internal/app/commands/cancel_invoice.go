package commands

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/payments/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/stackus/errors"
)

type (
	CancelInvoice struct {
		ID string
	}

	CancelInvoiceHandler struct {
		invoices domain.InvoiceRepository
		//publisher ddd.EventPublisher[ddd.Event]
		logger logger.Logger
	}
)

func NewCancelInvoiceHandler(
	invoices domain.InvoiceRepository,
	//publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) *CancelInvoiceHandler {
	return &CancelInvoiceHandler{
		invoices: invoices,
		//publisher: publisher,
		logger: logger,
	}
}

func (h CancelInvoiceHandler) Handle(ctx context.Context, cmd CancelInvoice) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CancelInvoiceHandler.Handle")
	defer span.Finish()

	invoice, err := h.invoices.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if invoice.Status != domain.InvoiceIsPending {
		return errors.Wrap(errors.ErrBadRequest, "invoice cannot be canceled")
	}

	invoice.Status = domain.InvoiceIsCanceled

	return h.invoices.Update(ctx, invoice)
}
