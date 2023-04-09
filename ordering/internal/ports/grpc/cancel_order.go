package grpc

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/ordering/internal/app/commands"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) CancelOrder(ctx context.Context, request *orderingpb.CancelOrderRequest) (resp *orderingpb.CancelOrderResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.CancelOrder(ctx, request)
}

func (s server) CancelOrder(ctx context.Context, request *orderingpb.CancelOrderRequest) (*orderingpb.CancelOrderResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.CancelOrder")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()

	err := s.cfg.App.Commands.CancelOrder.Handle(ctx, commands.CancelOrder{
		ID: request.GetId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to cancel order: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to cancel order: %s", err)
	}

	resp := &orderingpb.CancelOrderResponse{}

	return resp, nil
}
