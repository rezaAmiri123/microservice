package app

import (
	"context"
	"github.com/rezaAmiri123/microservice/notifications/internal/models"
)

type UserRepository interface {
	Find(ctx context.Context, userID string) (*models.User, error)
}
