package commands

import (
	"context"
	"github.com/rezaAmiri123/microservice/ordering/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/stackus/errors"
)

type (
	CancelOrder struct {
		ID string
	}

	CancelOrderHandler struct {
		orders    domain.OrderRepository
		publisher ddd.EventPublisher[ddd.Event]
		logger    logger.Logger
	}
)

func NewCancelOrderHandler(
	orders domain.OrderRepository,
	publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) CancelOrderHandler {
	return CancelOrderHandler{
		orders:    orders,
		publisher: publisher,
		logger:    logger,
	}
}

func (h CancelOrderHandler) CancelOrder(ctx context.Context, cmd CancelOrder) error {
	//span, ctx := opentracing.StartSpanFromContext(ctx, "CancelOrderHandler.Handle")
	//defer span.Finish()

	order, err := h.orders.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := order.Cancel()
	if err != nil {
		return errors.Wrap(err, "cancel order command")
	}

	// // TODO CH8 remove; handled in the saga
	// if err = h.shopping.Cancel(ctx, order.ShoppingID); err != nil {
	// 	return err
	// }

	if err = h.orders.Update(ctx, order); err != nil {
		return errors.Wrap(err, "complete order")
	}

	err = h.publisher.Publish(ctx, event)

	return err
}
