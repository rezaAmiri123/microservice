package user

import (
	"time"

	"github.com/google/uuid"
)

// User is user model
type User struct {
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	Username  string    `json:"username" db:"username" validate:"required,min=6,max=30"`
	Password  string    `json:"password" db:"password" validate:"required,min=8,max=15"`
	Email     string    `json:"email" db:"email" validate:"required,min=3,max=250,email"`
	Bio       string    `json:"bio" db:"bio"`
	Image     string    `json:"image" db:"image"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
