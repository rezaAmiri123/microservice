package grpc

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/ordering/internal/app/commands"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) ReadyOrder(ctx context.Context, request *orderingpb.ReadyOrderRequest) (resp *orderingpb.ReadyOrderResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.ReadyOrder(ctx, request)
}

func (s server) ReadyOrder(ctx context.Context, request *orderingpb.ReadyOrderRequest) (*orderingpb.ReadyOrderResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.ReadyOrder")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()

	err := s.cfg.App.Commands.ReadyOrder.Handle(ctx, commands.ReadyOrder{
		ID: request.GetId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to create order: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to ready order: %s", err)
	}

	resp := &orderingpb.ReadyOrderResponse{}

	return resp, nil
}
