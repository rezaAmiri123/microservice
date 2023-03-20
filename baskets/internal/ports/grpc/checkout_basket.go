package grpc

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/baskets/basketspb"
	"github.com/rezaAmiri123/microservice/baskets/internal/app/commands"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) CheckoutBasket(ctx context.Context, request *basketspb.CheckoutBasketRequest) (resp *basketspb.CheckoutBasketResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.CheckoutBasket(ctx, request)
}

func (s server) CheckoutBasket(ctx context.Context, request *basketspb.CheckoutBasketRequest) (*basketspb.CheckoutBasketResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.CheckoutBasket")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()

	err := s.cfg.App.Commands.CheckoutBasket.Handle(ctx, commands.CheckoutBasket{
		ID:        request.GetId(),
		PaymentID: request.GetPaymentId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to start basket: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to start basket: %s", err)
	}

	resp := &basketspb.CheckoutBasketResponse{}

	return resp, nil
}
