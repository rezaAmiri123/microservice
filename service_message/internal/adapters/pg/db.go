package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/service_message/internal/metrics"
)

func NewPGMessageRepository(db *sqlx.DB, log logger.Logger, metric *metrics.MessageServiceMetric) *PGMessageRepository {
	return &PGMessageRepository{DB: db, Logger: log, Metric: metric}
}

// News Repository
type PGMessageRepository struct {
	DB     *sqlx.DB
	Logger logger.Logger
	Metric *metrics.MessageServiceMetric
}
