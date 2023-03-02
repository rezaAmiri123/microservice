package pg

import (
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
)

func NewPGUserRepository(db postgres.DB, log logger.Logger) *PGUserRepository {
	return &PGUserRepository{DB: db, Logger: log}
}

// News Repository
type PGUserRepository struct {
	DB     postgres.DB
	Logger logger.Logger
	//Metric *metrics.UserServiceMetric
}

var _ domain.UserRepository = (*PGUserRepository)(nil)
