package pg

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
)

const getUser = `SELECT 
					username, email, bio, image, created_at, updated_at 
				FROM users
				WHERE id = $1`

func (r *PGUserRepository) Find(ctx context.Context, id string) (*domain.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGRepository.Find")
	defer span.Finish()

	u := domain.NewUser(id)
	err := r.DB.QueryRowContext(ctx, getUser, id).Scan(
		&u.Username, &u.Email, &u.Bio, &u.Enabled,
	)
	//if err := r.DB.GetContext(ctx, u, getUser, id); err != nil {
	//	return nil, fmt.Errorf("postgres connot get user: %w", err)
	//}
	return u, err

}
