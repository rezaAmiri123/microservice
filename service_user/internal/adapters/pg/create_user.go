package pg

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_user/internal/domain/user"
)

const createUser = `INSERT INTO users (username, password, email, bio, image) VALUES ($1, $2, $3, $4, $5) RETURNING *`

// const createUser = `INSERT INTO users (username, password, email, bio, image) VALUES ($1, $2, $3, $4, $5) RETURNING
// 						(user_id, username, password, email, bio, image, created_at, updated_at)`

// const getUser = `SELECT
// 					user_id, username, password, email, bio, image, created_at, updated_at
// 				FROM users
// 				WHERE username = $1`

func (r *PGUserRepository) CreateUser(ctx context.Context, arg *user.CreateUserParams) (*user.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGUserRepository.CreateUser")
	defer span.Finish()

	var u user.User
	if err := r.DB.QueryRowxContext(
		ctx,
		createUser,
		&arg.Username,
		&arg.Password,
		&arg.Email,
		&arg.Bio,
		&arg.Image,
		// ).Scan(&res); err != nil {
	).StructScan(&u); err != nil {
		return nil, fmt.Errorf("postgres connot create user: %w", err)
	}

	return &u, nil
}
