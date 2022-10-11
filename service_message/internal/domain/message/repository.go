//go:generate mockgen -source repository.go -destination mock/repository.go -package mock
package message

import "context"

type Repository interface {
	CreateEmail(ctx context.Context, arg *CreateEmailParams) (*Email, error)
}
