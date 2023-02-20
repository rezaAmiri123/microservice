package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/users/internal/app/commands"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) RegisterUser(ctx context.Context, req *userspb.RegisterUserRequest) (*userspb.RegisterUserResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.RegisterUser")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()
	id := uuid.New().String()
	err := s.cfg.App.Commands.RegisterUser.Handle(ctx, commands.RegisterUser{
		ID:       id,
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
	})
	if err != nil {
		s.cfg.Logger.Errorf("failed to register user: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}
	return &userspb.RegisterUserResponse{Id: id}, nil
}
