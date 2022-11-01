package queries

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
)

type GetTransfersHandler struct {
	logger logger.Logger
	repo   finance.Repository
}

func NewGetTransfersHandler(repo finance.Repository, logger logger.Logger) *GetTransfersHandler {
	if repo == nil {
		panic("userRepo is nil")
	}
	return &GetTransfersHandler{repo: repo, logger: logger}
}

func (h GetTransfersHandler) Handle(ctx context.Context, arg finance.ListTransferParams) (*finance.ListTransferResult, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetTransfersHandler.Handle")
	defer span.Finish()

	list, err := h.repo.ListTransfer(ctx, arg)

	if err != nil {
		h.logger.Errorf("connot get transfer list: %v", err)
		return nil, fmt.Errorf("connot get transfer list: %w", err)
	}

	return list, nil
}
