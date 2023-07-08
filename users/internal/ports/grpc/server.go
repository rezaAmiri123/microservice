package grpc

import (
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/users/internal/app"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type (
	Config struct {
		App    app.App
		Logger logger.Logger
	}
	server struct {
		cfg *Config
		userspb.UnimplementedUserServiceServer
	}
)

var _ userspb.UserServiceServer = (*server)(nil)

//func NewServer(app app.App, logger logger.Logger) *server {
//	cfg := &Config{
//		App:    app,
//		Logger: logger,
//	}
//	return &server{cfg: cfg}
//}

func (s server) userFromDomain(user *domain.User) *userspb.User {
	return &userspb.User{
		UserUuid:  user.ID(),
		Username:  user.Username,
		Email:     user.Email,
		Enabled:   user.Enabled,
		Bio:       user.Bio,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
