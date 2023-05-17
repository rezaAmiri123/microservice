package app

import (
	"context"
	"fmt"
)

type (
	OrderCreated struct {
		OrderID string
		UserID  string
	}
	OrderCanceled struct {
		OrderID string
		UserID  string
	}
	OrderReady struct {
		OrderID string
		UserID  string
	}

	App interface {
		NotifyOrderCreated(ctx context.Context, cmd OrderCreated) error
		NotifyOrderCanceled(ctx context.Context, cmd OrderCanceled) error
		NotifyOrderReady(ctx context.Context, cmd OrderReady) error
	}

	Application struct {
		users UserRepository
	}
)

var _ App = (*Application)(nil)

func New(users UserRepository) Application {
	return Application{
		users: users,
	}
}

func (a Application) NotifyOrderCreated(ctx context.Context, cmd OrderCreated) error {
	return fmt.Errorf("not implemented")

}

func (a Application) NotifyOrderCanceled(ctx context.Context, cmd OrderCanceled) error {
	return fmt.Errorf("not implemented")
}

func (a Application) NotifyOrderReady(ctx context.Context, cmd OrderReady) error {
	return fmt.Errorf("not implemented")
}
