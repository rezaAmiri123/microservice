package queries

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_api/internal/domain/api"
	financeservice "github.com/rezaAmiri123/microservice/service_finance/proto/grpc"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/tracing"
)

type GetTransfersHandler struct {
	client financeservice.FinanceServiceClient
	logger logger.Logger
}

func NewGetTransfersHandler(
	client financeservice.FinanceServiceClient,
	logger logger.Logger,
) *GetTransfersHandler {
	if client == nil {
		panic("rpc client is nil")
	}
	return &GetTransfersHandler{client: client, logger: logger}
}

func (h *GetTransfersHandler) Handle(ctx context.Context, req *api.GetTransfersRequest) (*api.GetTransfersResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetTransfersHandler.Handle")
	defer span.Finish()

	ctx = tracing.InjectTextMapCarrierToGrpcMetaData(ctx, span.Context())

	rpcReq := &financeservice.ListTransferRequest{
		Page:  req.Page,
		Size:  req.Size,
		Order: req.Order,
	}

	list, err := h.client.ListTransfer(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	res := &api.GetTransfersResponse{}
	res.TotalCount = list.TotalCount
	res.TotalPages = list.TotalPages
	res.Page = list.Page
	res.Size = list.Size
	res.HasMore = list.HasMore

	var transfersList = make([]*api.TransferResponse, 0, len(list.Transfers))
	for _, transfer := range list.Transfers {
		transfersList = append(transfersList, GrpcToTransfer(transfer))
	}

	res.Transfers = transfersList

	return res, nil
}

func GrpcToTransfer(req *financeservice.Transfer) *api.TransferResponse {
	res := &api.TransferResponse{}
	res.TransferID = req.TransferId
	res.FromAccountId = req.FromAccountId
	res.ToAccountId = req.ToAccountId
	res.Amount = req.Amount
	res.CreatedAt = req.CreatedAt.AsTime()
	res.UpdatedAt = req.UpdatedAt.AsTime()
	return res
}
