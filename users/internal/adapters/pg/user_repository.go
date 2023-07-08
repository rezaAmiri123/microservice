package pg

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
	"github.com/stackus/errors"
)

type UserRepository struct {
	tableName string
	db        postgres.DB
	//logger    logger.Logger
}

var _ domain.UserRepository = (*UserRepository)(nil)

func NewUserRepository(tableName string, db postgres.DB) UserRepository {
	return UserRepository{
		tableName: tableName,
		db:        db,
		//logger:    logger,
	}
}

func (r UserRepository) Find(ctx context.Context, userID string) (*domain.User, error) {
	const query = `SELECT username, email, enabled, bio, image FROM %s WHERE id = $1`

	user := domain.NewUser(userID)
	if err := r.db.QueryRowContext(ctx, r.table(query), userID).Scan(
		&user.Username,
		&user.Email,
		&user.Enabled,
		&user.Bio,
		&user.Image,
	); err != nil {
		if err == sql.ErrNoRows {
			err = errors.ErrNotFound.Err(err)
		}
		return nil, errors.Wrap(err, "database UserRepository.Find")
	}

	return user, nil
}

func (r UserRepository) Save(ctx context.Context, user *domain.User) (err error) {
	const query = `INSERT INTO %s (id, username, password, email, bio, image, enabled) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = r.db.ExecContext(ctx, r.table(query),
		user.ID(),
		user.Username,
		user.Password,
		user.Email,
		user.Bio,
		user.Image,
		user.Enabled,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				err = errors.ErrInvalidArgument.Err(pgErr)
			}
		}
		return errors.Wrap(err, "database UserRepository.Save")
	}

	return nil
}
func (r UserRepository) Update(ctx context.Context, user *domain.User) error {
	const query = `UPDATE %s 
						SET username = COALESCE(NULLIF($2, ''), username),
							email = COALESCE(NULLIF($3, ''), email),
							enabled = $4,
							bio = COALESCE(NULLIF($5, ''), bio),
							image = COALESCE(NULLIF($6, ''), image)
						WHERE id = $1`
	_, err := r.db.ExecContext(ctx, r.table(query),
		user.ID(),
		&user.Username,
		//&user.Password,
		&user.Email,
		&user.Enabled,
		&user.Bio,
		&user.Image,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				err = errors.ErrInvalidArgument.Err(pgErr)
			}
		}
		if err == sql.ErrNoRows {
			err = errors.ErrNotFound.Err(err)
		}
		return errors.Wrap(err, "database UserRepository.Update")
	}

	return nil
}

func (r UserRepository) table(query string) string {
	q := fmt.Sprintf(query, r.tableName)
	return q
}
