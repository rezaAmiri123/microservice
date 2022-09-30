package grpc

import (
	"context"

	"github.com/opentracing/opentracing-go"
	pkgGrpc "github.com/rezaAmiri123/microservice/pkg/grpc"
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

	violations := validateCreateTransferRequest(req)
	if violations != nil {
		return nil, pkgGrpc.InvalidArgumentError(violations)
	}

	arg := finance.TransferTxParams{Amount: req.Amount}
	copy(arg.FromAccountID[:], req.GetFromAccountId())
	copy(arg.ToAccountID[:], req.GetToAccountId())

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

func validateCreateTransferRequest(req *financeService.CreateTransferRequest) (violation []*errdetails.BadRequest_FieldViolation) {
	// if err := validator.ValidateCurrency(req.GetCurrency()); err != nil {
	// 	violation = append(violation, pkgGrpc.FieldViolation("currency", err))
	// }
	// if err := validator.ValidateBalance(req.GetBalance()); err != nil {
	// 	violation = append(violation, pkgGrpc.FieldViolation("balance", err))
	// }
	return
}
