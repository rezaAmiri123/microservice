package commands

import (
	"context"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/token"
	userservice "github.com/rezaAmiri123/microservice/service_user/proto/grpc"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/tracing"
)

type LoginVerifyHandler struct {
	client userservice.UserServiceClient
	logger logger.Logger
}

func NewLoginVerifyHandler(
	client userservice.UserServiceClient,
	logger logger.Logger,
) *LoginVerifyHandler {
	if client == nil {
		panic("rpc client is nil")
	}
	return &LoginVerifyHandler{client: client, logger: logger}
}

func (h *LoginVerifyHandler) Handle(ctx context.Context, t string) (*token.Payload, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LoginVerifyHandler.Handle")
	defer span.Finish()

	ctx = tracing.InjectTextMapCarrierToGrpcMetaData(ctx, span.Context())

	rpcReq := &userservice.LoginVerifyUserRequest{
		Token: t,
	}

	l, err := h.client.LoginVerify(ctx, rpcReq)
	if err != nil {
		return &token.Payload{}, err
	}

	// ID, err := utils.ConvertBase64ToUUID(l.Id)
	// ID, err := ConvertBase64ToUUID(l.Id)
	// if err != nil {
	// 	return &token.Payload{}, err
	// }
	ID, _ := uuid.Parse(l.Id)
	res := &token.Payload{
		ID:        ID,
		UserID:    l.UserId,
		Username:  l.Username,
		IssuedAt:  l.IssuedAt.AsTime(),
		ExpiredAt: l.ExpiredAt.AsTime(),
	}

	return res, nil
}
