package commands

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/ordering/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/stackus/errors"
)

type (
	ApproveOrder struct {
		ID         string
		ShoppingID string
	}

	ApproveOrderHandler struct {
		orders    domain.OrderRepository
		publisher ddd.EventPublisher[ddd.Event]
		logger    logger.Logger
	}
)

func NewApproveOrderHandler(
	orders domain.OrderRepository,
	publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) *ApproveOrderHandler {
	return &ApproveOrderHandler{
		orders:    orders,
		publisher: publisher,
		logger:    logger,
	}
}

func (h ApproveOrderHandler) Handle(ctx context.Context, cmd ApproveOrder) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ApproveOrderHandler.Handle")
	defer span.Finish()

	order, err := h.orders.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := order.Approve(cmd.ShoppingID)
	if err != nil {
		return errors.Wrap(err, "approve order command")
	}

	if err = h.orders.Update(ctx, order); err != nil {
		return errors.Wrap(err, "approve order")
	}

	err = h.publisher.Publish(ctx, event)

	return err
}