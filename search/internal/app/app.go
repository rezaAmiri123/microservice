package app

import (
	"context"
	"github.com/rezaAmiri123/microservice/search/internal/domain"
	"github.com/stackus/errors"
)

type (
	GetOrder struct {
		OrderID string
	}

	App interface {
		SearchOrders(ctx context.Context, cmd domain.SearchOrders) ([]*domain.Order, error)
		GetOrder(ctx context.Context, cmd GetOrder) (*domain.Order, error)
	}

	Application struct {
		orders domain.OrderRepository
	}
)

var _ App = (*Application)(nil)

func New(orders domain.OrderRepository) Application {
	return Application{
		orders: orders,
	}
}

func (a Application) SearchOrders(ctx context.Context, query domain.SearchOrders) ([]*domain.Order, error) {
	orders, err := a.orders.Search(ctx, query)
	return orders, errors.Wrap(err, "search order query")
}
func (a Application) GetOrder(ctx context.Context, query GetOrder) (*domain.Order, error) {
	order, err := a.orders.Get(ctx, query.OrderID)
	return order, errors.Wrap(err, "get order query")
}
