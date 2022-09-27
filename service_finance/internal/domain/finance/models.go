package finance

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	AccountID uuid.UUID `json:"account_id" db:"account_id"`
	OwnerID   uuid.UUID `json:"owner_id" db:"owner_id"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Entry struct {
	EntryID   uuid.UUID `json:"entry_id" db:"entry_id"`
	AccountID uuid.UUID `json:"account_id" db:"account_id"`
	// can be negative or positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Transfer struct {
	TransferID    uuid.UUID `json:"transfer_id" db:"transfer_id"`
	FromAccountID uuid.UUID `json:"from_account_id" db:"from_account_id"`
	ToAccountID   uuid.UUID `json:"to_account_id" db:"to_account_id"`
	// must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
