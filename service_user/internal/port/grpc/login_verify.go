package grpc

import (
	"context"

	"github.com/opentracing/opentracing-go"
	userService "github.com/rezaAmiri123/microservice/service_user/proto/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *UserGRPCServer) LoginVerify(ctx context.Context, req *userService.LoginVerifyUserRequest) (*userService.LoginVerifyUserResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserGRPCServer.LoginVerify")
	defer span.Finish()

	s.cfg.Metric.LoginVerifyRequests.Inc()

	payload, err := s.cfg.Maker.VerifyToken(req.GetToken())
	if err != nil {
		s.cfg.Logger.Errorf("invalid access token: %s", err)
		s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.InvalidArgument, "invalid access token: %s", err)
	}

	res := &userService.LoginVerifyUserResponse{
		Id:        payload.ID.String(),
		Username:  payload.Username,
		UserId:    payload.UserID,
		IssuedAt:  timestamppb.New(payload.IssuedAt),
		ExpiredAt: timestamppb.New(payload.ExpiredAt),
	}

	s.cfg.Metric.SuccessGrpcRequests.Inc()
	return res, nil
}
