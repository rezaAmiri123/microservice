package message

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type CreateEmailParams struct {
	UserID  uuid.UUID      `json:"user_id"`
	From    string         `json:"from"`
	To      pq.StringArray `json:"to"`
	Subject string         `json:"subject"`
	Body    string         `json:"body"`
}
