package pg

import (
	"context"
	"fmt"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
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
	err := r.db.QueryRowContext(ctx, r.table(query), userID).Scan(
		&user.Username,
		&user.Email,
		&user.Enabled,
		&user.Bio,
		&user.Image,
	)

	return user, err
}

func (r UserRepository) Save(ctx context.Context, user *domain.User) error {
	const query = `INSERT INTO %s (id, username, password, email, bio, image, enabled) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, r.table(query),
		user.ID(),
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Bio,
		&user.Image,
		&user.Enabled,
	)
	return err
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
	return err
}

func (r UserRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
