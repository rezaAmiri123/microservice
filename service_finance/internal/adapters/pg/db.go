package pg

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
	"github.com/rezaAmiri123/microservice/service_finance/internal/metrics"
)

func NewPGFinanceRepository(db *sqlx.DB, log logger.Logger, metric *metrics.FinanceServiceMetric) *PGFinanceRepository {
	return &PGFinanceRepository{DB: db, Logger: log, Metric: metric}
}

var _ finance.Repository = (*PGFinanceRepository)(nil)

// News Repository
type PGFinanceRepository struct {
	DB     *sqlx.DB
	Tx     *sqlx.Tx
	Logger logger.Logger
	Metric *metrics.FinanceServiceMetric
}

func (r *PGFinanceRepository) GetDB() DBTX {
	if r.Tx != nil {
		return r.Tx
	}
	return r.DB
}

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	sqlx.ExtContext
}
