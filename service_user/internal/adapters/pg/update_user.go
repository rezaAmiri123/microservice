package pg

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_user/internal/domain/user"
)

const updateUser = `UPDATE users 
						SET username = COALESCE(NULLIF($2, ''), username),
							password = COALESCE(NULLIF($3, ''), password), 
							email = COALESCE(NULLIF($4, ''), email), 
							bio = COALESCE(NULLIF($5, ''), bio),
							image = COALESCE(NULLIF($6, ''), image)
						WHERE username = $1
						RETURNING *`

func (r *PGUserRepository) UpdateUser(ctx context.Context, arg *user.UpdateUserParams, username string) (*user.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGUserRepository.UpdateUser")
	defer span.Finish()

	var u user.User
	if err := r.DB.QueryRowxContext(
		ctx,
		updateUser,
		username,
		&arg.Username,
		&arg.Password,
		&arg.Email,
		&arg.Bio,
		&arg.Image,
		// ).Scan(&res); err != nil {
	).StructScan(&u); err != nil {
		return nil, fmt.Errorf("database connot update user: %w", err)
	}

	return &u, nil
}
