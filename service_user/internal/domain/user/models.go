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

type Session struct {
	SessionID    uuid.UUID `json:"session_id"  db:"session_id"`
	Username     string    `json:"username" db:"username"`
	RefreshToken string    `json:"refresh_token" db:"refresh_token"`
	UserAgent    string    `json:"user_agent" db:"user_agent"`
	ClientIp     string    `json:"client_ip" db:"client_ip"`
	IsBlocked    bool      `json:"is_blocked" db:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at" db:"expires_at"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
