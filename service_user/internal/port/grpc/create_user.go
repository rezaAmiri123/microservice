package grpc

import (
	"context"

	"github.com/opentracing/opentracing-go"
	pkgGrpc "github.com/rezaAmiri123/microservice/pkg/grpc"
	"github.com/rezaAmiri123/microservice/service_user/internal/domain/user"
	"github.com/rezaAmiri123/microservice/service_user/internal/validator"
	userService "github.com/rezaAmiri123/microservice/service_user/proto/grpc"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserGRPCServer) CreateUser(ctx context.Context, req *userService.CreateUserRequest) (*userService.CreateUserResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserGRPCServer.CreateUser")
	defer span.Finish()

	s.cfg.Metric.CreateUserGrpcRequests.Inc()

	violations := validateCreateUserRequest(req)
	if violations != nil {
		return nil, pkgGrpc.InvalidArgumentError(violations)
	}

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

func validateCreateUserRequest(req *userService.CreateUserRequest) (violation []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateUsername(req.GetUsername()); err != nil {
		violation = append(violation, pkgGrpc.FieldViolation("username", err))
	}
	if err := validator.ValidatePassword(req.GetPassword()); err != nil {
		violation = append(violation, pkgGrpc.FieldViolation("password", err))
	}
	if err := validator.ValidateEmail(req.GetEmail()); err != nil {
		violation = append(violation, pkgGrpc.FieldViolation("email", err))
	}
	return
}
