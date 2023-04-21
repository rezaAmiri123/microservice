package commands

import (
	"context"
	"github.com/rezaAmiri123/microservice/ordering/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/stackus/errors"
)

type (
	CompleteOrder struct {
		ID        string
		InvoiceID string
	}

	CompleteOrderHandler struct {
		orders    domain.OrderRepository
		publisher ddd.EventPublisher[ddd.Event]
		logger    logger.Logger
	}
)

func NewCompleteOrderHandler(
	orders domain.OrderRepository,
	publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) CompleteOrderHandler {
	return CompleteOrderHandler{
		orders:    orders,
		publisher: publisher,
		logger:    logger,
	}
}

func (h CompleteOrderHandler) CompleteOrder(ctx context.Context, cmd CompleteOrder) error {
	//span, ctx := opentracing.StartSpanFromContext(ctx, "CompleteOrderHandler.Handle")
	//defer span.Finish()

	order, err := h.orders.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := order.Complete(cmd.InvoiceID)
	if err != nil {
		return errors.Wrap(err, "complete order command")
	}

	if err = h.orders.Update(ctx, order); err != nil {
		return errors.Wrap(err, "complete order")
	}

	err = h.publisher.Publish(ctx, event)

	return err
}
