package domain

import (
	"context"
)

type BasketRepository interface {
	Load(ctx context.Context, basketID string) (*Basket, error)
	Save(ctx context.Context, basket *Basket) error
	//UpdateItems(ctx context.Context, basket *Basket) error
	//Update(ctx context.Context, basket *Basket) error
}
