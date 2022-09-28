package pg

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
)

const createAccount = `INSERT INTO accounts (owner_id, balance, currency) VALUES ($1, $2, $3) RETURNING *`

func (r *PGFinanceRepository) CreateAccount(ctx context.Context, arg *finance.CreateAccountParams) (*finance.Account, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGFinanceRepository.CreateUser")
	defer span.Finish()

	var a finance.Account
	if err := r.DB.QueryRowxContext(
		ctx,
		createAccount,
		&arg.OwnerID,
		&arg.Balance,
		&arg.Currency,
	).StructScan(&a); err != nil {
		return nil, fmt.Errorf("database connot create account: %w", err)
	}

	return &a, nil
}
