package pg

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_user/internal/domain/user"
)

// const createUser = `INSERT INTO users (username, password, email, bio, image) VALUES ($1, $2, $3, $4, $5) RETURNING
//
//	(user_id, username, password, email, bio, image, created_at, updated_at)`
const getUser = `SELECT 
					user_id, username, password, email, bio, image, created_at, updated_at 
				FROM users
				WHERE username = $1`

func (r *PGUserRepository) GetUserByUsername(ctx context.Context, username string) (*user.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGUserRepository.GetUserByUsername")
	defer span.Finish()

	u := &user.User{}
	if err := r.DB.GetContext(ctx, u, getUser, username); err != nil {
		return nil, fmt.Errorf("postgres connot get user: %w", err)
	}
	return u, nil
}
