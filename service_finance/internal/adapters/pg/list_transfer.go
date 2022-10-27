package pg

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
)

const (
	listTransfer          = `INSERT INTO transfers (from_account_id, to_account_id, amount) VALUES ($1, $2, $3) RETURNING *`
	getTotalTransferCount = `SELECT COUNT(transfer_id) FROM transfers`
)

func (r PGFinanceRepository) ListTransfer(ctx context.Context, arg finance.ListTransferParams) (finance.ListTransferResult, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGFinanceRepository.ListTransfer")
	defer span.Finish()

	var totalCount int64
	if err := r.GetDB().GetContext(ctx, &totalCount, getTotalTransferCount); err != nil {
		return finance.ListTransferResult{}, fmt.Errorf("database connot get transfer count: %w", err)
	}

	var result finance.ListTransferResult
	if totalCount == 0 {
		result.TotalCount = totalCount

		return result{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
			Page:       pq.GetPage(),
			Size:       pq.GetSize(),
			HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
			News:       make([]*models.News, 0),
		}, nil
	}

	var t finance.Transfer
	if err := r.GetDB().QueryRowxContext(
		ctx,
		listTransfer,
		&arg.FromAccountID,
		&arg.ToAccountID,
		&arg.Amount,
	).StructScan(&t); err != nil {
		return finance.Transfer{}, fmt.Errorf("database connot create transfer: %w", err)
	}

	return t, nil
}
