package user

import (
	"time"

	"github.com/google/uuid"
)

// User is user model
type User struct {
	UserID    uuid.UUID `json:"user_id"`
	Username  string    `json:"username" validate:"required,min=6,max=30"`
	Password  string    `json:"password" validate:"required,min=8,max=15"`
	Email     string    `json:"email" validate:"required,min=3,max=250,email"`
	Bio       string    `json:"bio"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
