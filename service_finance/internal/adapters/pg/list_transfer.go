package pg

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
)

const (
	// listTransfer = `SELECT transfer_id, from_account_id, to_account_id, amount, updated_at, created_at
	// 					FROM transfers
	// 					ORDER BY COALESCE(NULLIF($3, ''), created_at) OFFSET $1 LIMIT $2`
	listTransfer = `SELECT t.transfer_id, t.from_account_id, t.to_account_id, t.amount, t.updated_at, t.created_at 
						FROM transfers AS t
						INNER JOIN accounts AS a ON a.account_id = t.from_account_id
						WHERE a.owner_id = $1
						ORDER BY $4 OFFSET $2 LIMIT $3`

	getTotalTransferCount = `SELECT COUNT(transfer_id) FROM transfers AS t
								INNER JOIN accounts AS a ON a.account_id = t.from_account_id
								WHERE a.owner_id = $1`
)

func (r PGFinanceRepository) ListTransfer(ctx context.Context, arg finance.ListTransferParams) (*finance.ListTransferResult, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGFinanceRepository.ListTransfer")
	defer span.Finish()

	var totalCount int
	if err := r.GetDB().GetContext(ctx, &totalCount, getTotalTransferCount, arg.OwnerID); err != nil {
		return nil, fmt.Errorf("database connot get transfer count: %w", err)
	}

	result := &finance.ListTransferResult{}
	if totalCount == 0 {
		result.TotalCount = int64(totalCount)
		result.TotalPages = int64(arg.Paginate.GetTotalPages(totalCount))
		result.Page = int64(arg.Paginate.GetPage())
		result.Size = int64(arg.Paginate.GetSize())
		result.HasMore = arg.Paginate.GetHasMore(totalCount)
		result.Transfers = make([]*finance.Transfer, 0)
		return result, nil
	}

	var transfersList = make([]*finance.Transfer, 0, arg.Paginate.GetSize())
	rows, err := r.GetDB().QueryxContext(
		ctx,
		listTransfer,
		arg.OwnerID,
		arg.Paginate.GetOffset(),
		arg.Paginate.GetLimit(),
		arg.Paginate.GetOrderBy())
	if err != nil {
		return nil, fmt.Errorf("database connot get transfer list: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		n := &finance.Transfer{}
		if err = rows.StructScan(n); err != nil {
			return nil, fmt.Errorf("database connot get transfers: %w", err)
		}
		transfersList = append(transfersList, n)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("database connot get transfers rows : %w", err)
	}

	result.TotalCount = int64(totalCount)
	result.TotalPages = int64(arg.Paginate.GetTotalPages(totalCount))
	result.Page = int64(arg.Paginate.GetPage())
	result.Size = int64(arg.Paginate.GetSize())
	result.HasMore = arg.Paginate.GetHasMore(totalCount)

	result.Transfers = transfersList
	return result, nil
}
