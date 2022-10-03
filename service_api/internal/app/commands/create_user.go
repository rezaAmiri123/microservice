package commands

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/utils"
	"github.com/rezaAmiri123/microservice/service_api/internal/domain/api"
	userservice "github.com/rezaAmiri123/microservice/service_user/proto/grpc"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/tracing"
)

type CreateUserHandler struct {
	client userservice.UserServiceClient
	logger logger.Logger
}

func NewCreateUserHandler(
	client userservice.UserServiceClient,
	logger logger.Logger,
) *CreateUserHandler {
	if client == nil {
		panic("rpc client is nil")
	}
	return &CreateUserHandler{client: client, logger: logger}
}

func (h *CreateUserHandler) Handle(ctx context.Context, req *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateUserHandler.Handle")
	defer span.Finish()

	ctx = tracing.InjectTextMapCarrierToGrpcMetaData(ctx, span.Context())

	rpcReq := &userservice.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Bio:      req.Bio,
		Image:    req.Image,
	}
	u, err := h.client.CreateUser(ctx, rpcReq)
	if err != nil {
		return &api.CreateUserResponse{}, err
	}
	fmt.Println(u.User)
	user := u.User
	userID, err := utils.ConvertBase64ToUUID(user.UserUuid)
	if err != nil {
		return &api.CreateUserResponse{}, err
	}

	res := &api.CreateUserResponse{
		UserID:   userID,
		Username: user.Username,
		Email:    user.Email,
		Bio:      user.Bio,
		Image:    user.Image,
	}

	return res, nil
}
