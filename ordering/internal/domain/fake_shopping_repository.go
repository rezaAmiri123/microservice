package domain

import "context"

type FakeShoppingRepository struct {
}

func NewFakeShoppingRepository() *FakeShoppingRepository {
	return &FakeShoppingRepository{}
}

var _ ShoppingRepository = (*FakeShoppingRepository)(nil)

func (r *FakeShoppingRepository) Create(ctx context.Context, orderID string, items []Item) (string, error) {
	panic("impliment me")
}

func (r *FakeShoppingRepository) Cancel(ctx context.Context, shoppingID string) error {
	panic("impliment me")
}

func (r *FakeShoppingRepository) Reset() {
	//panic("impliment me")
}
