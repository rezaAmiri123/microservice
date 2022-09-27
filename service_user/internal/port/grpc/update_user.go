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

func (s *UserGRPCServer) UpdateUser(ctx context.Context, req *userService.UpdateUserRequest) (*userService.UpdateUserResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserGRPCServer.UpdateUser")
	defer span.Finish()

	s.cfg.Metric.UpdateUserGrpcRequests.Inc()

	authPayload, err := s.AuthorizeUser(ctx)
	if err != nil {
		return nil, pkgGrpc.UnauthenticatedError(err)
	}

	violations := validateUpdateUserRequest(req)
	if violations != nil {
		return nil, pkgGrpc.InvalidArgumentError(violations)
	}

	arg := &user.UpdateUserParams{}
	if req.Username != nil {
		arg.Password = req.GetUsername()
	}
	if req.Password != nil {
		arg.Password = req.GetPassword()
	}
	if req.Email != nil {
		arg.Email = req.GetEmail()
	}
	if req.Bio != nil {
		arg.Bio = req.GetBio()
	}
	if req.Image != nil {
		arg.Image = req.GetImage()
	}

	u, err := s.cfg.App.Commands.UpdateUser.Handle(ctx, arg, authPayload.Username)
	if err != nil {
		s.cfg.Logger.Errorf("failed to update user: %s", err)
		s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to update user: %s", err)
	}

	res := &userService.UpdateUserResponse{
		User: UserToGrpc(u),
	}

	s.cfg.Metric.SuccessGrpcRequests.Inc()
	return res, nil
}

func validateUpdateUserRequest(req *userService.UpdateUserRequest) (violation []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateUsername(req.GetUsername()); err != nil {
		violation = append(violation, pkgGrpc.FieldViolation("username", err))
	}
	if err := validator.ValidatePassword(req.GetPassword()); err != nil {
		violation = append(violation, pkgGrpc.FieldViolation("password", err))
	}
	if err := validator.ValidateEmail(req.GetEmail()); err != nil {
		violation = append(violation, pkgGrpc.FieldViolation("email", err))
	}
	return violation
}
