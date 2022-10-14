package message

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Email struct {
	EmailID   uuid.UUID      `json:"email_id" db:"email_id"`
	UserID    uuid.UUID      `json:"user_id" db:"user_id"`
	From      string         `json:"from" db:"from_email"`
	To        pq.StringArray `json:"to" db:"to_email"`
	Subject   string         `json:"subject" db:"subject"`
	Body      string         `json:"body" db:"body"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
}
