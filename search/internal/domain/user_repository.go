package domain

import "context"

type UserRepository interface {
	Find(ctx context.Context, userID string) (*User, error)
}

type UserCacheRepository interface {
	UserRepository
	Add(ctx context.Context, userID, name string) error
}
