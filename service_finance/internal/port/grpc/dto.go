package grpc

import (
	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
	financeService "github.com/rezaAmiri123/microservice/service_finance/proto/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func AccountToGrpc(f *finance.Account) *financeService.Account {
	res := &financeService.Account{}
	res.AccountId = f.AccountID.String()
	res.OwnerId = f.OwnerID.String()
	res.Balance = f.Balance
	res.Currency = f.Currency
	res.CreatedAt = timestamppb.New(f.CreatedAt)
	res.UpdatedAt = timestamppb.New(f.UpdatedAt)
	return res
}

func TransferToGrpc(f *finance.Transfer) *financeService.Transfer {
	res := &financeService.Transfer{}
	res.TransferId = f.TransferID.String()
	res.FromAccountId = f.FromAccountID.String()
	res.ToAccountId = f.ToAccountID.String()
	res.Amount = f.Amount
	res.CreatedAt = timestamppb.New(f.CreatedAt)
	res.UpdatedAt = timestamppb.New(f.UpdatedAt)
	return res
}

func TransferListToGrpc(f *finance.ListTransferResult) *financeService.ListTransferResponse {
	res := &financeService.ListTransferResponse{}
	res.TotalCount = f.TotalCount
	res.TotalPages = f.TotalPages
	res.Page = f.Page
	res.Size = f.Size
	res.HasMore = f.HasMore

	list := make([]*financeService.Transfer, 0, len(f.Transfers))
	for _, transfer := range f.Transfers {
		list = append(list, TransferToGrpc(transfer))
	}
	res.Transfers = list
	return res
}
