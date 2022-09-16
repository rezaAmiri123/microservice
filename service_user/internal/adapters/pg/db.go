package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/service_user/internal/metrics"
)

func NewPGUserRepository(db *sqlx.DB, log logger.Logger, metric *metrics.UserServiceMetric) *PGUserRepository {
	return &PGUserRepository{DB: db, Logger: log, Metric: metric}
}

// News Repository
type PGUserRepository struct {
	DB     *sqlx.DB
	Logger logger.Logger
	Metric *metrics.UserServiceMetric
}
