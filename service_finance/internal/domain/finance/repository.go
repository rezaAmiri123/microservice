//go:generate mockgen -source repository.go -destination mock/repository.go -package finance_mock
package finance

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	CreateAccount(ctx context.Context, arg *CreateAccountParams) (*Account, error)
	GetAccountByID(ctx context.Context, accountID uuid.UUID) (*Account, error)
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	ExecTx(ctx context.Context, fn func(Repository) error) error
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	ListTransfer(ctx context.Context, arg ListTransferParams) (ListTransferResult, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
}
