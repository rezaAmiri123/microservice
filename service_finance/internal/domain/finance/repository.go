//go:generate mockgen -source repository.go -destination mock/repository.go -package finance_mock
package finance

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	CreateAccount(ctx context.Context, arg *CreateAccountParams) (*Account, error)
	GetAccountByID(ctx context.Context, accountID uuid.UUID) (*Account, error)
}
