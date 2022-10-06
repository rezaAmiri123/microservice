package api

import (
	"time"

	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	OwnerId  string `json:"owner_id" validate:"required"`
	Balance  int64  `json:"balance" validate:"required"`
	Currency string `json:"currency" validate:"required"`
}

type CreateAccountResponse struct {
	AccountId string    `json:"account_id" validate:"required"`
	OwnerId   string    `json:"owner_id" validate:"required`
	Balance   int64     `json:"balance" validate:"required"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateTransferRequest struct {
	FromAccountId uuid.UUID `json:"from_account_id" validate:"required"`
	ToAccountId   uuid.UUID `json:"to_account_id" validate:"required"`
	Amount        int64     `json:"amount" validate:"required"`
}

type CreateTransferResponse struct {
	TransferID    uuid.UUID `json:"transfer_id" validate:"required"`
	FromAccountId uuid.UUID `json:"from_account_id" validate:"required,min=6,max=30"`
	ToAccountId   uuid.UUID `json:"to_account_id" validate:"required,min=3,max=250,email"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
