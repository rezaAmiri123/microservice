package pg

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
)

const updateUser = `UPDATE users 
						SET username = COALESCE(NULLIF($2, ''), username),
-- 							password = COALESCE(NULLIF($3, ''), password), 
							email = COALESCE(NULLIF($4, ''), email), 
							bio = COALESCE(NULLIF($5, ''), bio),
							image = COALESCE(NULLIF($6, ''), image)
						WHERE id = $1
						RETURNING *`

func (r *PGUserRepository) Update(ctx context.Context, user *domain.User) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGRepository.Update")
	defer span.Finish()

	if err := r.DB.QueryRowxContext(
		ctx,
		updateUser,
		user.ID(),
		&user.Username,
		//&user.Password,
		&user.Email,
		&user.Bio,
		&user.Image,
	).Err(); err != nil {
		return fmt.Errorf("postgres connot create user: %w", err)
	}

	return nil

}
