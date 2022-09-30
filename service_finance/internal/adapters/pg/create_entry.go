package pg

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
)

const createEntry = `INSERT INTO entries (account_id, amount) VALUES ($1, $2) RETURNING *`

func (r PGFinanceRepository) CreateEntry(ctx context.Context, arg finance.CreateEntryParams) (finance.Entry, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGFinanceRepository.CreateEntry")
	defer span.Finish()

	var e finance.Entry
	if err := r.GetDB().QueryRowxContext(
		ctx,
		createEntry,
		&arg.AccountID,
		&arg.Amount,
	).StructScan(&e); err != nil {
		return finance.Entry{}, fmt.Errorf("database connot create entry: %w", err)
	}

	return e, nil
}
