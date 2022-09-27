package commands

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
)

type CreateAccountHandler struct {
	logger logger.Logger
	repo   finance.Repository
}

func NewCreateAccountHandler(repo finance.Repository, logger logger.Logger) *CreateAccountHandler {
	if repo == nil {
		panic("userRepo is nil")
	}
	return &CreateAccountHandler{repo: repo, logger: logger}
}

func (h CreateAccountHandler) Handle(ctx context.Context, arg *finance.CreateAccountParams) (*finance.Account, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateAccountHandler.Handle")
	defer span.Finish()

	account, err := h.repo.CreateAccount(ctx, arg)
	if err != nil {
		h.logger.Errorf("connot create user: %v", err)
		return &finance.Account{}, fmt.Errorf("connot create account: %w", err)
	}

	return account, nil
}
