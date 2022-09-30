package pg

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
)

// const createEntry = `INSERT INTO entries (account_id, amount) VALUES ($1, $2) RETURNING *`

func (r PGFinanceRepository) ExecTx(ctx context.Context, fn func(finance.Repository) error) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGFinanceRepository.ExecTx")
	defer span.Finish()

	tx, err := r.DB.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	txRepo := PGFinanceRepository{
		DB:     r.DB,
		Tx:     tx,
		Logger: r.Logger,
		Metric: r.Metric,
	}

	err = fn(txRepo)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// type DBTX interface {
// 	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
// 	PrepareContext(context.Context, string) (*sql.Stmt, error)
// 	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
// 	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
// }

// func New(tx DBTX, pgRepo PGFinanceRepository) finance.Repository {
// 	pgRepo.DB.BeginTxx() = tx
// 	return PGFinanceRepository{
// 		DB: db,
// 	}
// }

// type Queries struct {
// 	db DBTX
// }

// func (q *Queries) WithTx(tx *sql.Tx) *Queries {
// 	return &Queries{
// 		db: tx,
// 	}
// }
