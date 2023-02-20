package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
)

func NewPGUserRepository(db *sqlx.DB, log logger.Logger) *PGUserRepository {
	return &PGUserRepository{DB: db, Logger: log}
}

// News Repository
type PGUserRepository struct {
	DB     *sqlx.DB
	Logger logger.Logger
	//Metric *metrics.UserServiceMetric
}

var _ domain.UserRepository = (*PGUserRepository)(nil)
