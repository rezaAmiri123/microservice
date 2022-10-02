package api

import "github.com/google/uuid"

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
