package grpc

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/baskets/basketspb"
	"github.com/rezaAmiri123/microservice/baskets/internal/app/commands"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) CancelBasket(ctx context.Context, request *basketspb.CancelBasketRequest) (resp *basketspb.CancelBasketResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.CancelBasket(ctx, request)
}

func (s server) CancelBasket(ctx context.Context, request *basketspb.CancelBasketRequest) (*basketspb.CancelBasketResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.CancelBasket")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()

	err := s.cfg.App.CancelBasket(ctx, commands.CancelBasket{
		ID: request.GetId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to cancel basket: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to cancel basket: %s", err)
	}

	resp := &basketspb.CancelBasketResponse{}

	return resp, nil
}
