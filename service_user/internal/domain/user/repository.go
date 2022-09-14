//go:generate mockgen -source repository.go -destination mock/repository.go -package user_mock
package user

import "context"

type Repository interface {
	CreateUser(ctx context.Context, arg *CreateUserParams) (*User, error)
	//Update(ctx context.Context, user *User)error
	//GetByEmail(ctx context.Context, email string)(*User,error)
	//GetByUsername(ctx context.Context, username string) (*User, error)
}
