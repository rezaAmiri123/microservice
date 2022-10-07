package queries

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
)

type GetAccountByIDHandler struct {
	logger logger.Logger
	repo   finance.Repository
}

func NewGetAccountByIDHandler(repo finance.Repository, logger logger.Logger) *GetAccountByIDHandler {
	if repo == nil {
		panic("userRepo is nil")
	}
	return &GetAccountByIDHandler{repo: repo, logger: logger}
}

func (h GetAccountByIDHandler) Handle(ctx context.Context, accountID uuid.UUID) (*finance.Account, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetAccountByIDHandler.Handle")
	defer span.Finish()

	account, err := h.repo.GetAccountByID(ctx, accountID)
	if err != nil {
		h.logger.Errorf("connot get account: %v", err)
		return &finance.Account{}, fmt.Errorf("connot get account: %w", err)
	}

	return account, nil
}
