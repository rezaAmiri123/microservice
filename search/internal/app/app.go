package app

import (
	"context"
	"fmt"
	"github.com/rezaAmiri123/microservice/search/internal/domain"
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

func (a *Application) SearchOrders(ctx context.Context, cmd domain.SearchOrders) ([]*domain.Order, error) {
	return nil, fmt.Errorf("not implemented")
}
func (a *Application) GetOrder(ctx context.Context, cmd GetOrder) (*domain.Order, error) {
	return nil, fmt.Errorf("not implemented")
}
