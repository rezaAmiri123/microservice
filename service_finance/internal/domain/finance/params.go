package finance

import "github.com/google/uuid"

type CreateAccountParams struct {
	OwnerID  uuid.UUID `json:"owner_id"`
	Balance  int64     `json:"balance"`
	Currency string    `json:"currency"`
}
