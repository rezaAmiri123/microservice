package domain

import "context"

type StoreRepository interface {
	Find(ctx context.Context, storeID string) (*Store, error)
}

type StoreCacheRepository interface {
	Add(ctx context.Context, storeID, name string) error
	Rename(ctx context.Context, storeID, name string) error
	StoreRepository
}
