package api

import (
	"time"
)

type CreateAccountRequest struct {
	OwnerId  string `json:"owner_id"`
	Balance  int64  `json:"balance" validate:"required"`
	Currency string `json:"currency" validate:"required"`
}

type CreateAccountResponse struct {
	AccountId string    `json:"account_id" validate:"required"`
	OwnerId   string    `json:"owner_id" validate:"required"`
	Balance   int64     `json:"balance" validate:"required"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateTransferRequest struct {
	OwnerId       string `json:"owner_id"`
	FromAccountId string `json:"from_account_id" validate:"required"`
	ToAccountId   string `json:"to_account_id" validate:"required"`
	Amount        int64  `json:"amount" validate:"required"`
}
type ListResponse struct {
	TotalCount int64 `json:"total_count"`
	TotalPages int64 `json:"total_pages"`
	Page       int64 `json:"page"`
	Size       int64 `json:"size"`
	HasMore    bool  `json:"has_more"`
}
type FilterRequest struct {
	Page  int64  `json:"page"`
	Size  int64  `json:"size"`
	Order string `json:"order"`
}

type TransferResponse struct {
	TransferID    string    `json:"transfer_id" validate:"required"`
	FromAccountId string    `json:"from_account_id" validate:"required,min=6,max=30"`
	ToAccountId   string    `json:"to_account_id" validate:"required,min=3,max=250,email"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
type CreateTransferResponse struct {
	TransferResponse
}

type GetTransfersRequest struct {
	FilterRequest
}

type GetTransfersResponse struct {
	ListResponse
	Transfers []*TransferResponse `json:"transfers"`
}
