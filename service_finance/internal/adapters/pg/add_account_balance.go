package pg

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
)

const updateAccountBalance = `UPDATE accounts
								SET balance = balance + $1
							  WHERE id = $2 RETURNING *`

func (r PGFinanceRepository) AddAccountBalance(ctx context.Context, arg finance.AddAccountBalanceParams) (finance.Account, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGFinanceRepository.AddAccountBalance")
	defer span.Finish()
	var res finance.Account
	if err := r.GetDB().QueryRowxContext(
		ctx,
		updateAccountBalance,
		&arg.Amount,
		&arg.AccountID,
	).StructScan(&res); err != nil {
		return finance.Account{}, fmt.Errorf("database connot add balance: %w", err)
	}

	return res, nil
}
