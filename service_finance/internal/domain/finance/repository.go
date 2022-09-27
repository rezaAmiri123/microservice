//go:generate mockgen -source repository.go -destination mock/repository.go -package finance_mock
package finance

import "context"

type Repository interface {
	CreateAccount(ctx context.Context, arg *CreateAccountParams) (*Account, error)
}
