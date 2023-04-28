package domain

import "context"

type ProductRepository interface {
	Find(ctx context.Context, productID string) (*Product, error)
}

type ProductCacheRepository interface {
	Add(ctx context.Context, productID, storeID, name string) error
	Rebrand(ctx context.Context, productID, name string) error
	Remove(ctx context.Context, productID string) error
	ProductRepository
}
