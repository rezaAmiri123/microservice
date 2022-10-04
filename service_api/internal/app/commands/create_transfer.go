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

type CreateTransferHandler struct {
	client financeservice.FinanceServiceClient
	logger logger.Logger
}

func NewCreateTransferHandler(
	client financeservice.FinanceServiceClient,
	logger logger.Logger,
) *CreateTransferHandler {
	if client == nil {
		panic("rpc client is nil")
	}
	return &CreateTransferHandler{client: client, logger: logger}
}

func (h *CreateTransferHandler) Handle(ctx context.Context, req *api.CreateTransferRequest) (*api.CreateTransferResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateTransferHandler.Handle")
	defer span.Finish()

	ctx = tracing.InjectTextMapCarrierToGrpcMetaData(ctx, span.Context())

	rpcReq := &financeservice.CreateTransferRequest{
		FromAccountId: req.FromAccountId[:],
		ToAccountId:   req.ToAccountId[:],
		Amount:        req.Amount,
	}
	t, err := h.client.CreateTransfer(ctx, rpcReq)
	if err != nil {
		return &api.CreateTransferResponse{}, err
	}

	transfer := t.Transfer
	TransferID, _ := utils.ConvertBase64ToUUID(transfer.TransferId)
	FromAccountId, _ := utils.ConvertBase64ToUUID(transfer.FromAccountId)
	ToAccountId, _ := utils.ConvertBase64ToUUID(transfer.ToAccountId)
	res := &api.CreateTransferResponse{
		TransferID:    TransferID,
		FromAccountId: FromAccountId,
		ToAccountId:   ToAccountId,
		Amount:        transfer.Amount,
		CreatedAt:     transfer.CreatedAt.AsTime(),
		UpdatedAt:     transfer.UpdatedAt.AsTime(),
	}

	return res, nil
}
