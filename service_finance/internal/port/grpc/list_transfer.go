package grpc

import (
	"context"

	"github.com/opentracing/opentracing-go"
	pkgGrpc "github.com/rezaAmiri123/microservice/pkg/grpc"
	"github.com/rezaAmiri123/microservice/pkg/pagnation"
	"github.com/rezaAmiri123/microservice/service_finance/internal/domain/finance"
	financeService "github.com/rezaAmiri123/microservice/service_finance/proto/grpc"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *FinanceGRPCServer) ListTransfer(ctx context.Context, req *financeService.ListTransferRequest) (*financeService.ListTransferResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "FinanceGRPCServer.ListTransfer")
	defer span.Finish()

	s.cfg.Metric.ListTransferGrpcRequests.Inc()

	arg, violations := s.validateListTransferRequest(ctx, req)
	if violations != nil {
		return nil, pkgGrpc.InvalidArgumentError(violations)
	}

	t, err := s.cfg.App.Queries.GetTransfers.Handle(ctx, arg)
	if err != nil {
		s.cfg.Logger.Errorf("failed to get transfers: %s", err)
		s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to get transfers: %s", err)
	}

	res := TransferListToGrpc(t)

	s.cfg.Metric.SuccessGrpcRequests.Inc()
	return res, nil
}

func (s *FinanceGRPCServer) validateListTransferRequest(ctx context.Context, req *financeService.ListTransferRequest) (
	arg finance.ListTransferParams,
	violation []*errdetails.BadRequest_FieldViolation,
) {
	page := pagnation.NewPaginationQuery(int(req.GetSize()), int(req.GetPage()), req.GetOrder())
	arg.Paginate = page

	// TODO make filter based on user
	return
}
