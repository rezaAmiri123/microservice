package api

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=6,max=30"`
	Password string `json:"password" validate:"required,min=8,max=15"`
	Email    string `json:"email" validate:"required,min=3,max=250,email"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

type CreateUserResponse struct {
	UserID   uuid.UUID `json:"user-id" validate:"required"`
	Username string    `json:"username" validate:"required,min=6,max=30"`
	Email    string    `json:"email" validate:"required,min=3,max=250,email"`
	Bio      string    `json:"bio"`
	Image    string    `json:"image"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=6,max=30"`
	Password string `json:"password" validate:"required,min=8,max=15"`
}

type LoginResponse struct {
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}
