package pg

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
)

const getAccount = `SELECT 
					account_id, owner_id, balance, currency, created_at, updated_at
				FROM accounts
					WHERE account_id = $1`

func (r PGFinanceRepository) GetAccountByID(ctx context.Context, accountID uuid.UUID) (*finance.Account, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGFinanceRepository.GetAccountByID")
	defer span.Finish()

	a := &finance.Account{}
	if err := r.GetDB().GetContext(ctx, a, getAccount, accountID); err != nil {
		return nil, fmt.Errorf("database connot get account: %w", err)
	}
	return a, nil
}
