package pg

import (
	"context"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
)

const getUser = `SELECT 
					username, email, Enabled, bio, image 
				FROM users
				WHERE id = $1`

func (r *PGUserRepository) Find(ctx context.Context, id string) (*domain.User, error) {
	//span, ctx := opentracing.StartSpanFromContext(ctx, "PGRepository.Find")
	//defer span.Finish()

	const query = `SELECT username, email, Enabled, bio, image FROM users WHERE id = $1`

	u := domain.NewUser(id)
	if err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&u.Username,
		&u.Email,
		&u.Enabled,
		&u.Bio,
		&u.Image,
	); err != nil {
		//fmt.Println("pg error after QueryRowContext, ", err.Error())
		return nil, err
	}
	//if err := r.DB.GetContext(ctx, u, getUser, id); err != nil {
	//	return nil, fmt.Errorf("postgres connot get user: %w", err)
	//}
	return u, nil

}
