package commands

import (
	"context"
	"github.com/rezaAmiri123/microservice/payments/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/logger"
)

type (
	AuthorizePayment struct {
		ID     string
		UserID string
		Amount float64
	}

	AuthorizePaymentHandler struct {
		payments domain.PaymentRepository
		//publisher ddd.EventPublisher[ddd.Event]
		logger logger.Logger
	}
)

func NewAuthorizePaymentHandler(
	payments domain.PaymentRepository,
	//publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) AuthorizePaymentHandler {
	return AuthorizePaymentHandler{
		payments: payments,
		//publisher: publisher,
		logger: logger,
	}
}

func (h AuthorizePaymentHandler) AuthorizePayment(ctx context.Context, cmd AuthorizePayment) error {
	//span, ctx := opentracing.StartSpanFromContext(ctx, "AuthorizePaymentHandler.Handle")
	//defer span.Finish()

	return h.payments.Save(ctx, &domain.Payment{
		ID:     cmd.ID,
		UserID: cmd.UserID,
		Amount: cmd.Amount,
	})
}
