package commands

import (
	"context"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
)

type CreateTransferHandler struct {
	logger logger.Logger
	repo   finance.Repository
}

func NewCreateTransferHandler(repo finance.Repository, logger logger.Logger) *CreateTransferHandler {
	if repo == nil {
		panic("repo is nil")
	}
	return &CreateTransferHandler{repo: repo, logger: logger}
}

func (h CreateTransferHandler) Handle(ctx context.Context, arg finance.TransferTxParams) (finance.TransferTxResult, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateTransferHandler.Handle")
	defer span.Finish()

	var result finance.TransferTxResult

	err := h.repo.ExecTx(ctx, func(r finance.Repository) error {
		var err error
		result.Transfer, err = r.CreateTransfer(ctx, finance.CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = r.CreateEntry(ctx, finance.CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = r.CreateEntry(ctx, finance.CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.FromAccountID.String() < arg.ToAccountID.String() {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, r, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, r, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}

		return err
	})

	return result, err

}

func addMoney(
	ctx context.Context,
	r finance.Repository,
	accountID1 uuid.UUID,
	amount1 int64,
	accountID2 uuid.UUID,
	amount2 int64,
) (account1 finance.Account, account2 finance.Account, err error) {
	account1, err = r.AddAccountBalance(ctx, finance.AddAccountBalanceParams{
		AccountID: accountID1,
		Amount:    amount1,
	})
	if err != nil {
		return
	}
	account2, err = r.AddAccountBalance(ctx, finance.AddAccountBalanceParams{
		AccountID: accountID2,
		Amount:    amount2,
	})
	return
}
