package grpc

import (
	"context"
	"github.com/opentracing/opentracing-go"
	pkgGrpc "github.com/rezaAmiri123/microservice/pkg/grpc"
	"github.com/rezaAmiri123/microservice/pkg/utils"
	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
	financeService "github.com/rezaAmiri123/microservice/service_finance/proto/grpc"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *FinanceGRPCServer) CreateTransfer(ctx context.Context, req *financeService.CreateTransferRequest) (*financeService.CreateTransferResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "FinanceGRPCServer.CreateTransfer")
	defer span.Finish()

	s.cfg.Metric.CreateTransferGrpcRequests.Inc()

	arg, violations := validateCreateTransferRequest(req)
	if violations != nil {
		return nil, pkgGrpc.InvalidArgumentError(violations)
	}

	t, err := s.cfg.App.Commands.CreateTransfer.Handle(ctx, arg)
	if err != nil {
		s.cfg.Logger.Errorf("failed to create transfer: %s", err)
		s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to create transfer: %s", err)
	}

	res := &financeService.CreateTransferResponse{
		Transfer: TransferToGrpc(&t.Transfer),
	}

	s.cfg.Metric.SuccessGrpcRequests.Inc()
	return res, nil
}

func validateCreateTransferRequest(req *financeService.CreateTransferRequest) (
	arg finance.TransferTxParams,
	violation []*errdetails.BadRequest_FieldViolation,
) {
	fromAccountID, err := utils.ConvertBase64ToUUID(req.GetFromAccountId())
	if err != nil {
		violation = append(violation, pkgGrpc.FieldViolation("from_account_id", err))
	}
	arg.FromAccountID = fromAccountID

	toAccountID, err := utils.ConvertBase64ToUUID(req.GetToAccountId())
	if err != nil {
		violation = append(violation, pkgGrpc.FieldViolation("to_account_id", err))
	}
	arg.ToAccountID = toAccountID

	arg.Amount = req.GetAmount()
	return
}
