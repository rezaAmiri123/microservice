package commands

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/utils"
	"github.com/rezaAmiri123/microservice/service_api/internal/domain/api"
	financeservice "github.com/rezaAmiri123/microservice/service_finance/proto/grpc"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/tracing"
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

	rpcReq := &financeservice.CreateAccountRequest{
		OwnerId:  req.OwnerId[:],
		Balance:  req.Balance,
		Currency: req.Currency,
	}
	t, err := h.client.CreateAccount(ctx, rpcReq)
	if err != nil {
		return &api.CreateAccountResponse{}, err
	}

	account := t.Account
	accountID, _ := utils.ConvertBase64ToUUID(account.AccountId)
	ownerID, _ := utils.ConvertBase64ToUUID(account.OwnerId)
	res := &api.CreateAccountResponse{
		OwnerId:   ownerID,
		AccountId: accountID,
		Balance:   account.Balance,
		Currency:  account.Currency,
		CreatedAt: account.CreatedAt.AsTime(),
		UpdatedAt: account.UpdatedAt.AsTime(),
	}

	return res, nil
}
