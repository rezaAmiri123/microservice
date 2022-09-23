//go:generate mockgen -source repository.go -destination mock/repository.go -package user_mock
package user

import "context"

type Repository interface {
	CreateUser(ctx context.Context, arg *CreateUserParams) (*User, error)
	UpdateUser(ctx context.Context, arg *UpdateUserParams, username string) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	CreateSession(ctx context.Context, arg *CreateSessionParams) (*Session, error)
	//Update(ctx context.Context, user *User)error
	//GetByEmail(ctx context.Context, email string)(*User,error)
}
