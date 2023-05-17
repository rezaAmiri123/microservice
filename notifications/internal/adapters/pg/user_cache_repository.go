package pg

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rezaAmiri123/microservice/notifications/internal/app"
	"github.com/rezaAmiri123/microservice/notifications/internal/models"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/stackus/errors"
)

type UserCacheRepository struct {
	tableName string
	db        postgres.DB
	fallback  app.UserRepository
}

var _ app.UserCacheRepository = (*UserCacheRepository)(nil)

func NewUserCacheRepository(tableName string, db postgres.DB, fallback app.UserRepository) UserCacheRepository {
	return UserCacheRepository{
		tableName: tableName,
		db:        db,
		fallback:  fallback,
	}
}

func (r UserCacheRepository) Add(ctx context.Context, userID, username string) error {
	const query = "INSERT INTO %s (id, username) VALUES ($1, $2) ON CONFLICT DO NOTHING"

	_, err := r.db.ExecContext(ctx, r.table(query), userID, username)

	return err
}

func (r UserCacheRepository) Find(ctx context.Context, userID string) (*models.User, error) {
	const query = `SELECT username FROM %s WHERE id = $1 LIMIT 1`

	user := &models.User{
		ID: userID,
	}

	err := r.db.QueryRowContext(ctx, r.table(query), userID).Scan(
		&user.Username,
	)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "scanning user")
		}
		user, err = r.fallback.Find(ctx, userID)
		if err != nil {
			return nil, errors.Wrap(err, "user callback failed")
		}
		// attempt to add it to the cache
		return user, r.Add(ctx, user.ID, user.Username)
	}

	return user, nil
}

func (r UserCacheRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
