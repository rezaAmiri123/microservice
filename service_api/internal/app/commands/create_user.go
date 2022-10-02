package commands

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_api/internal/domain/api"
	userservice "github.com/rezaAmiri123/microservice/service_user/proto/grpc"
	"github.com/rezaAmiri123/test-microservice/api_service/dto"
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
) CreateUserHandler {
	if client == nil {
		panic("article client is nil")
	}
	return CreateUserHandler{client: client, logger: logger}
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
	list, err := h.client.CreateUser(ctx, rpcReq)
	if err != nil {
		return &api.CreateUserResponse{}, err
	}

	res := &dto.GetEmailsResponse{}
	res.TotalCount = list.GetTotalCount()
	res.TotalPages = list.GetTotalPages()
	res.Page = list.GetPage()
	res.Size = list.GetSize()
	res.HasMore = list.GetHasMore()

	emailList := make([]*dto.EmailResponse, 0, len(list.GetEmails()))
	for _, a := range list.GetEmails() {
		emailList = append(emailList, dto.EmailResponseFromGrpc(a))
	}
	res.Emails = emailList

	return res, nil
}
