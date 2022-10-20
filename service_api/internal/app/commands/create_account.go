package commands

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/pkg/tracing"
	"github.com/rezaAmiri123/microservice/service_api/internal/domain/api"
	financeservice "github.com/rezaAmiri123/microservice/service_finance/proto/grpc"
)

type CreateAccountHandler struct {
	client financeservice.FinanceServiceClient
	logger logger.Logger
}

func NewCreateAccountHandler(
	client financeservice.FinanceServiceClient,
	logger logger.Logger,
) *CreateAccountHandler {
	if client == nil {
		panic("rpc client is nil")
	}
	return &CreateAccountHandler{client: client, logger: logger}
}

func (h *CreateAccountHandler) Handle(ctx context.Context, req *api.CreateAccountRequest) (*api.CreateAccountResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateAccountHandler.Handle")
	defer span.Finish()

	ctx = tracing.InjectTextMapCarrierToGrpcMetaData(ctx, span.Context())
	// utils.ConvertBase64ToUUID(req.OwnerId)
	// var x uuid.UUID
	// ownerID, _ := uuid.Parse(req.OwnerId)
	// fmt.Println("**************************************************")
	// fmt.Println(ownerID)
	// fmt.Println(ownerID.MarshalBinary())
	rpcReq := &financeservice.CreateAccountRequest{
		OwnerId:  req.OwnerId,
		Balance:  req.Balance,
		Currency: req.Currency,
	}
	t, err := h.client.CreateAccount(ctx, rpcReq)
	if err != nil {
		return &api.CreateAccountResponse{}, err
	}

	account := t.Account
	res := &api.CreateAccountResponse{
		OwnerId:   account.OwnerId,
		AccountId: account.AccountId,
		Balance:   account.Balance,
		Currency:  account.Currency,
		CreatedAt: account.CreatedAt.AsTime(),
		UpdatedAt: account.UpdatedAt.AsTime(),
	}

	return res, nil
}
