package app

import "context"

type UserCacheRepository interface {
	Add(ctx context.Context, userID, Username string) error
	UserRepository
}
