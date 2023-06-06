package commands

import (
	"context"
	"github.com/rezaAmiri123/microservice/ordering/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/stackus/errors"
)

type (
	CreateOrder struct {
		ID        string
		UserID    string
		PaymentID string
		Items     []domain.Item
	}

	CreateOrderHandler struct {
		orders    domain.OrderRepository
		publisher ddd.EventPublisher[ddd.Event]
		logger    logger.Logger
	}
)

func NewCreateOrderHandler(
	orders domain.OrderRepository,
	publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) CreateOrderHandler {
	return CreateOrderHandler{
		orders:    orders,
		publisher: publisher,
		logger:    logger,
	}
}

func (h CreateOrderHandler) CreateOrder(ctx context.Context, cmd CreateOrder) error {
	order, err := h.orders.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := order.CreateOrder(cmd.UserID, cmd.PaymentID, cmd.Items)
	if err != nil {
		return errors.Wrap(err, "create order command")
	}

	if err = h.orders.Save(ctx, order); err != nil {
		return errors.Wrap(err, "order creation")
	}

	err = h.publisher.Publish(ctx, event)

	return err
}
