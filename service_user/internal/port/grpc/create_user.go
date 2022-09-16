package grpc

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_user/internal/domain/user"
	userService "github.com/rezaAmiri123/microservice/service_user/proto/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserGRPCServer) CreateUser(ctx context.Context, req *userService.CreateUserRequest) (*userService.CreateUserResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserGRPCServer.CreateUser")
	defer span.Finish()

	s.cfg.Metric.CreateUserGrpcRequests.Inc()

	arg := &user.CreateUserParams{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
		Bio:      req.GetBio(),
		Image:    req.GetImage(),
	}
	u, err := s.cfg.App.Commands.CreateUser.Handle(ctx, arg)
	if err != nil {
		s.cfg.Logger.Errorf("failed to create user: %s", err)
		s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	res := &userService.CreateUserResponse{
		User: UserToGrpc(u),
	}

	s.cfg.Metric.SuccessGrpcRequests.Inc()
	return res, nil
}
