package grpc

import (
	"context"
	pkgGrpc "github.com/rezaAmiri123/microservice/pkg/grpc"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/search/internal/domain"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"google.golang.org/grpc"
)

type UserRepository struct {
	endpoint string
	logger   logger.Logger
}

var _ domain.UserRepository = (*UserRepository)(nil)

func NewUserRepository(endpoint string, logger logger.Logger) UserRepository {
	return UserRepository{
		endpoint: endpoint,
		logger:   logger,
	}
}

func (r UserRepository) Find(ctx context.Context, userID string) (user *domain.User, err error) {
	var conn *grpc.ClientConn
	conn, err = r.dial(ctx)

	if err != nil {
		return nil, err
	}

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	resp, err := userspb.NewUserServiceClient(conn).
		GetUser(ctx, &userspb.GetUserRequest{Id: userID})

	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:   resp.GetUser().GetUserUuid(),
		Name: resp.GetUser().GetUsername(),
	}, nil
}

func (r UserRepository) dial(ctx context.Context) (*grpc.ClientConn, error) {
	return pkgGrpc.NewContextGrpcClient(ctx, r.endpoint, nil, r.logger)
}
