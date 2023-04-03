package domain

import (
	"context"
)

type UserRepository interface {
	Authorize(ctx context.Context, userID string) error
}
