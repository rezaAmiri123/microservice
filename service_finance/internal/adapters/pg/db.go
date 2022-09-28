package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/service_finance/internal/metrics"
)

func NewPGFinanceRepository(db *sqlx.DB, log logger.Logger, metric *metrics.FinanceServiceMetric) *PGFinanceRepository {
	return &PGFinanceRepository{DB: db, Logger: log, Metric: metric}
}

// News Repository
type PGFinanceRepository struct {
	DB     *sqlx.DB
	Logger logger.Logger
	Metric *metrics.FinanceServiceMetric
}
