package grpc

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/ordering/internal/app/commands"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) CompleteOrder(ctx context.Context, request *orderingpb.CompleteOrderRequest) (resp *orderingpb.CompleteOrderResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.CompleteOrder(ctx, request)
}

func (s server) CompleteOrder(ctx context.Context, request *orderingpb.CompleteOrderRequest) (*orderingpb.CompleteOrderResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.CompleteOrder")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()

	err := s.cfg.App.Commands.CompleteOrder.Handle(ctx, commands.CompleteOrder{
		ID:        request.GetId(),
		InvoiceID: request.GetInvoiceId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to complete order: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to complete order: %s", err)
	}

	resp := &orderingpb.CompleteOrderResponse{}

	return resp, nil
}
