package domain

import "context"

type OrderRepository interface {
	Add(ctx context.Context, order *Order) error
	UpdateStatus(ctx context.Context, orderID, status string) error
	Search(ctx context.Context, search SearchOrders) ([]*Order, error)
	Get(ctx context.Context, orderID string) (*Order, error)
}
