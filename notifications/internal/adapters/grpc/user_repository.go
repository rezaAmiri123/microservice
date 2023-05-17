package grpc

import (
	"context"
	"github.com/rezaAmiri123/microservice/notifications/internal/app"
	"github.com/rezaAmiri123/microservice/notifications/internal/models"
	pkgGrpc "github.com/rezaAmiri123/microservice/pkg/grpc"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"google.golang.org/grpc"
)

type UserRepository struct {
	endpoint string
	logger   logger.Logger
}

var _ app.UserRepository = (*UserRepository)(nil)

func NewUserRepository(endpoint string, logger logger.Logger) UserRepository {
	return UserRepository{
		endpoint: endpoint,
		logger:   logger,
	}
}

func (r UserRepository) Find(ctx context.Context, userID string) (user *models.User, err error) {
	var conn *grpc.ClientConn
	conn, err = r.dial(ctx)
	if err != nil {
		return nil, err
	}

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	resp, err := userspb.NewUserServiceClient(conn).GetUser(ctx, &userspb.GetUserRequest{
		Id: userID,
	})
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:       resp.GetUser().GetUserUuid(),
		Username: resp.GetUser().GetUsername(),
	}, nil
}

func (r UserRepository) dial(ctx context.Context) (*grpc.ClientConn, error) {
	return pkgGrpc.NewContextGrpcClient(ctx, r.endpoint, nil, r.logger)
}
