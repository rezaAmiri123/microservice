package domain

import "context"

type FakeOrderRepository struct {
	orders map[string]*Order
}

func NewFakeOrderRepository() *FakeOrderRepository {
	return &FakeOrderRepository{orders: make(map[string]*Order)}
}

var _ OrderRepository = (*FakeOrderRepository)(nil)

func (r *FakeOrderRepository) Load(ctx context.Context, orderID string) (*Order, error) {
	if order, exists := r.orders[orderID]; exists {
		return order, nil
	}
	return NewOrder(orderID), nil
}

func (r *FakeOrderRepository) Save(ctx context.Context, order *Order) error {
	for _, event := range order.Events() {
		if err := order.ApplyEvent(event); err != nil {
			return err
		}
	}
	r.orders[order.ID()] = order
	return nil
}

func (r *FakeOrderRepository) Reset(orders ...*Order) {
	r.orders = make(map[string]*Order)

	for _, order := range orders {
		r.orders[order.ID()] = order
	}
}
