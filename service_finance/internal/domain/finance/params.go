package finance

import (
	"github.com/google/uuid"
	"github.com/rezaAmiri123/microservice/pkg/pagnation"
)

type (
	// FilterParams struct {
	// 	Page  int64  `json:"page"`
	// 	Size  int64  `json:"size"`
	// 	Order string `json:"order"`
	// }

	ListResult struct {
		TotalCount int64 `json:"total_count"`
		TotalPages int64 `json:"total_pages"`
		Page       int64 `json:"page"`
		Size       int64 `json:"size"`
		HasMore    bool  `json:"has_more"`
	}

	CreateAccountParams struct {
		OwnerID  uuid.UUID `json:"owner_id"`
		Balance  int64     `json:"balance"`
		Currency string    `json:"currency"`
	}

	CreateTransferParams struct {
		FromAccountID uuid.UUID `json:"from_account_id"`
		ToAccountID   uuid.UUID `json:"to_account_id"`
		Amount        int64     `json:"amount"`
	}

	ListTransferParams struct {
		Paginate pagnation.Pagination
	}

	ListTransferResult struct {
		ListResult
		Transfers []*Transfer
	}

	CreateEntryParams struct {
		AccountID uuid.UUID `json:"account_id"`
		Amount    int64     `json:"amount"`
	}

	AddAccountBalanceParams struct {
		Amount    int64     `json:"amount"`
		AccountID uuid.UUID `json:"id"`
	}

	// TransferTxParams contains the input parameters of the transfer transaction
	TransferTxParams struct {
		FromAccountID uuid.UUID `json:"from_account_id"`
		ToAccountID   uuid.UUID `json:"to_account_id"`
		Amount        int64     `json:"amount"`
	}

	// TransferTxResult is the result of the transfer transaction
	TransferTxResult struct {
		Transfer    Transfer `json:"transfer"`
		FromAccount Account  `json:"from_account"`
		ToAccount   Account  `json:"to_account"`
		FromEntry   Entry    `json:"from_entry"`
		ToEntry     Entry    `json:"to_entry"`
	}
)
