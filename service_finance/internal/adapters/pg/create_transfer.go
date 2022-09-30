package pg

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
)

const createTransfer = `INSERT INTO transfers (from_account_id, to_account_id, amount) VALUES ($1, $2, $3) RETURNING *`

func (r PGFinanceRepository) CreateTransfer(ctx context.Context, arg finance.CreateTransferParams) (finance.Transfer, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGFinanceRepository.CreateTransfer")
	defer span.Finish()

	var t finance.Transfer
	if err := r.GetDB().QueryRowxContext(
		ctx,
		createTransfer,
		&arg.FromAccountID,
		&arg.ToAccountID,
		&arg.Amount,
	).StructScan(&t); err != nil {
		return finance.Transfer{}, fmt.Errorf("database connot create transfer: %w", err)
	}

	return t, nil
}
