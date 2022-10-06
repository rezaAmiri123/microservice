package commands

import (
	"context"
	"encoding/base64"
	"fmt"

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

func ConvertBase64ToUUID(value []byte) (uuid.UUID, error) {
	strValue := string(value)
	data, err := base64.StdEncoding.DecodeString(strValue)
	fmt.Println("*****************************************")
	fmt.Println(err.Error())
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("cannot decode: %w", err)
	}
	valueByte, err := uuid.FromBytes(data)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("cannot decode: %w", err)
	}
	return valueByte, nil
}

func (h *LoginVerifyHandler) Handle(ctx context.Context, t string) (*token.Payload, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LoginVerifyHandler.Handle")
	defer span.Finish()

	ctx = tracing.InjectTextMapCarrierToGrpcMetaData(ctx, span.Context())

	rpcReq := &userservice.LoginVerifyUserRequest{
		Token: t,
	}
	fmt.Println("***********************************************8")
	fmt.Println("")
	l, err := h.client.LoginVerify(ctx, rpcReq)
	if err != nil {
		return &token.Payload{}, err
	}

	// ID, err := utils.ConvertBase64ToUUID(l.Id)
	// ID, err := ConvertBase64ToUUID(l.Id)
	// if err != nil {
	// 	return &token.Payload{}, err
	// }
	ID, _ := uuid.ParseBytes(l.Id)
	res := &token.Payload{
		ID:        ID,
		Username:  l.Username,
		IssuedAt:  l.IssuedAt.AsTime(),
		ExpiredAt: l.ExpiredAt.AsTime(),
	}

	return res, nil
}
