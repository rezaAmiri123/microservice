package queries

import (
	"context"
	"github.com/rezaAmiri123/microservice/ordering/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/stackus/errors"
)

type (
	GetOrder struct {
		ID string
	}

	GetOrderHandler struct {
		orders domain.OrderRepository
		//publisher ddd.EventPublisher[ddd.Event]
		logger logger.Logger
	}
)

func NewGetOrderHandler(
	orders domain.OrderRepository,
	//publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) GetOrderHandler {
	return GetOrderHandler{
		orders: orders,
		//publisher: publisher,
		logger: logger,
	}
}

func (h GetOrderHandler) GetOrder(ctx context.Context, query GetOrder) (*domain.Order, error) {
	order, err := h.orders.Load(ctx, query.ID)

	return order, errors.Wrap(err, "get order query")
}
