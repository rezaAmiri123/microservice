package grpc

import (
	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
	financeService "github.com/rezaAmiri123/microservice/service_finance/proto/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func AccountToGrpc(f *finance.Account) *financeService.Account {
	res := &financeService.Account{}
	res.AccountId = f.AccountID[:]
	res.OwnerId = f.OwnerID[:]
	res.Balance = f.Balance
	res.Currency = f.Currency
	res.CreatedAt = timestamppb.New(f.CreatedAt)
	res.UpdatedAt = timestamppb.New(f.UpdatedAt)
	return res
}

func TransferToGrpc(f *finance.Transfer) *financeService.Transfer {
	res := &financeService.Transfer{}
	res.TransferId = f.TransferID[:]
	res.FromAccountId = f.FromAccountID[:]
	res.ToAccountId = f.ToAccountID[:]
	res.Amount = f.Amount
	res.CreatedAt = timestamppb.New(f.CreatedAt)
	res.UpdatedAt = timestamppb.New(f.UpdatedAt)
	return res
}
