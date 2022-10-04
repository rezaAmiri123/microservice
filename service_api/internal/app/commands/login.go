package commands

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_api/internal/domain/api"
	userservice "github.com/rezaAmiri123/microservice/service_user/proto/grpc"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/tracing"
)

type LoginHandler struct {
	client userservice.UserServiceClient
	logger logger.Logger
}

func NewLoginHandler(
	client userservice.UserServiceClient,
	logger logger.Logger,
) *LoginHandler {
	if client == nil {
		panic("rpc client is nil")
	}
	return &LoginHandler{client: client, logger: logger}
}

func (h *LoginHandler) Handle(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LoginHandler.Handle")
	defer span.Finish()

	ctx = tracing.InjectTextMapCarrierToGrpcMetaData(ctx, span.Context())

	rpcReq := &userservice.LoginUserRequest{
		Username: req.Username,
		Password: req.Password,
	}
	l, err := h.client.Login(ctx, rpcReq)
	if err != nil {
		return &api.LoginResponse{}, err
	}

	res := &api.LoginResponse{
		AccessToken:           l.AccessToken,
		RefreshToken:          l.RefreshToken,
		AccessTokenExpiresAt:  l.AccessTokenExpiresAt.AsTime(),
		RefreshTokenExpiresAt: l.RefreshTokenExpiresAt.AsTime(),
	}

	return res, nil
}
