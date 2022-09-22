package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/rezaAmiri123/microservice/service_user/pkg/token"
)

type CreateUserParams struct {
	Username string `json:"username" validate:"required,min=6,max=30"`
	Password string `json:"password" validate:"required,min=8,max=15"`
	Email    string `json:"email" validate:"required,min=3,max=250,email"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

type CreateSessionParams struct {
	SessionID    uuid.UUID `json:"session_id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type LoginRequestParams struct {
	Username string `json:"username" validate:"required,min=6,max=30"`
	Password string `json:"password" validate:"required,min=8,max=15"`
}

type LoginResponseParams struct {
	AccessToken    string
	AccessPayload  *token.Payload
	RefreshToken   string
	RefreshPayload *token.Payload
	Session        *Session
	User           *User
}
