package grpc

import (
	"context"

	"github.com/opentracing/opentracing-go"
	pkgGrpc "github.com/rezaAmiri123/microservice/pkg/grpc"
	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
	"github.com/rezaAmiri123/microservice/service_finance/internal/validator"
	financeService "github.com/rezaAmiri123/microservice/service_finance/proto/grpc"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *FinanceGRPCServer) CreateAccount(ctx context.Context, req *financeService.CreateAccountRequest) (*financeService.CreateAccountResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "FinanceGRPCServer.CreateAccount")
	defer span.Finish()

	s.cfg.Metric.CreateAccountGrpcRequests.Inc()

	violations := validateCreateAccountRequest(req)
	if violations != nil {
		return nil, pkgGrpc.InvalidArgumentError(violations)
	}

	arg := &finance.CreateAccountParams{
		Balance:  req.GetBalance(),
		Currency: req.GetCurrency(),
	}
	copy(arg.OwnerID[:], req.GetOwnerId())

	a, err := s.cfg.App.Commands.CreateAccount.Handle(ctx, arg)
	if err != nil {
		s.cfg.Logger.Errorf("failed to create account: %s", err)
		s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to create account: %s", err)
	}

	res := &financeService.CreateAccountResponse{
		Account: AccountToGrpc(a),
	}

	s.cfg.Metric.SuccessGrpcRequests.Inc()
	return res, nil
}

func validateCreateAccountRequest(req *financeService.CreateAccountRequest) (violation []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateCurrency(req.GetCurrency()); err != nil {
		violation = append(violation, pkgGrpc.FieldViolation("currency", err))
	}
	if err := validator.ValidateBalance(req.GetBalance()); err != nil {
		violation = append(violation, pkgGrpc.FieldViolation("balance", err))
	}
	return
}
