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
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *UserGRPCServer) Login(ctx context.Context, req *userService.LoginUserRequest) (*userService.LoginUserResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserGRPCServer.Login")
	defer span.Finish()

	s.cfg.Metric.LoginRequests.Inc()

	violations := validateLoginRequest(req)
	if violations != nil {
		return nil, pkgGrpc.InvalidArgumentError(violations)
	}

	arg := &user.LoginRequestParams{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}
	login, err := s.cfg.App.Commands.Login.Handle(ctx, arg)
	if err != nil {
		s.cfg.Logger.Errorf("failed to login: %s", err)
		s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to login: %s", err)
	}

	res := &userService.LoginUserResponse{
		User:                  UserToGrpc(login.User),
		SessionId:             login.Session.SessionID.String(),
		AccessToken:           login.AccessToken,
		RefreshToken:          login.RefreshToken,
		AccessTokenExpiresAt:  timestamppb.New(login.AccessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(login.RefreshPayload.ExpiredAt),
	}

	s.cfg.Metric.SuccessGrpcRequests.Inc()
	return res, nil
}

func validateLoginRequest(req *userService.LoginUserRequest) (violation []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateUsername(req.GetUsername()); err != nil {
		violation = append(violation, pkgGrpc.FieldViolation("username", err))
	}
	if err := validator.ValidatePassword(req.GetPassword()); err != nil {
		violation = append(violation, pkgGrpc.FieldViolation("password", err))
	}

	return
}
