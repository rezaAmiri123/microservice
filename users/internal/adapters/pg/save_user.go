package pg

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
)

func (r *PGUserRepository) Save(ctx context.Context, user *domain.User) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGRepository.Save")
	defer span.Finish()

	const query = `INSERT INTO users (id, username, password, email, bio, image) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.DB.ExecContext(ctx, query,
		user.ID(),
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Bio,
		&user.Image,
	)
	//if err := r.DB.QueryRowxContext(
	//	ctx,
	//	query,
	//	user.ID(),
	//	&user.Username,
	//	&user.Password,
	//	&user.Email,
	//	&user.Bio,
	//	&user.Image,
	//).Err(); err != nil {
	//	//return fmt.Errorf("postgres connot create user: %w", err)
	//}

	return err

}
