package commands

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/payments/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/stackus/errors"
)

type (
	ConfirmPayment struct {
		ID string
	}

	ConfirmPaymentHandler struct {
		payments domain.PaymentRepository
		//publisher ddd.EventPublisher[ddd.Event]
		logger logger.Logger
	}
)

func NewConfirmPaymentHandler(
	payments domain.PaymentRepository,
//publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) *ConfirmPaymentHandler {
	return &ConfirmPaymentHandler{
		payments: payments,
		//publisher: publisher,
		logger: logger,
	}
}

func (h ConfirmPaymentHandler) Handle(ctx context.Context, cmd ConfirmPayment) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ConfirmPaymentHandler.Handle")
	defer span.Finish()

	if payment, err := h.payments.Find(ctx, cmd.ID); err != nil || payment == nil {
		return errors.Wrap(errors.ErrNotFound, "payment cannot be confirmed")
	}
	
	return nil
}
