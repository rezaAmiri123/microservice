package pg

import (
	"context"
	"fmt"

	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
)

func (r PGFinanceRepository) TransferTx(ctx context.Context, arg finance.TransferTxParams) (finance.TransferTxResult, error) {
	return finance.TransferTxResult{}, fmt.Errorf("Not Emplemented")
}
