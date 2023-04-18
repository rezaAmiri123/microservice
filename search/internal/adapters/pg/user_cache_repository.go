package pg

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/search/internal/domain"
	"github.com/stackus/errors"
)

type UserCacheRepository struct {
	tableName string
	db        postgres.DB
	fallback  domain.UserRepository
}

var _ domain.UserCacheRepository = (*UserCacheRepository)(nil)

func NewUserCacheRepository(tableName string, db postgres.DB, fallback domain.UserRepository) UserCacheRepository {
	return UserCacheRepository{
		tableName: tableName,
		db:        db,
		fallback:  fallback,
	}
}

func (r UserCacheRepository) Add(ctx context.Context, userID, name string) error {
	const query = "INSERT INTO %s (id, NAME) VALUES ($1, $2)"

	_, err := r.db.ExecContext(ctx, r.table(query), userID, name)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return nil
			}
		}
	}
	return err
}

func (r UserCacheRepository) Find(ctx context.Context, userID string) (*domain.User, error) {
	const query = `SELECT name FROM %s WHERE id = $1 LIMIT 1`

	user := &domain.User{
		ID: userID,
	}

	err := r.db.QueryRowContext(ctx, r.table(query), userID).Scan(
		&user.Name,
	)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "scanning user")
		}
		user, err = r.fallback.Find(ctx, userID)
		if err != nil {
			return nil, errors.Wrap(err, "user fallback failed")
		}
		// attempt to add it to the cache
		return user, r.Add(ctx, user.ID, user.Name)
	}
	return user, nil
}

func (r UserCacheRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
