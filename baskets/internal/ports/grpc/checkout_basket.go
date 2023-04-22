package grpc

import (
	"context"
	"github.com/rezaAmiri123/microservice/baskets/basketspb"
	"github.com/rezaAmiri123/microservice/baskets/internal/app/commands"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	grpcCodes "google.golang.org/grpc/codes"
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
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("BasketID", request.GetId()),
		attribute.String("PaymentID", request.GetPaymentId()),
	)

	err := s.cfg.App.CheckoutBasket(ctx, commands.CheckoutBasket{
		ID:        request.GetId(),
		PaymentID: request.GetPaymentId(),
	})

	if err != nil {
		//s.cfg.Logger.Errorf("failed to start basket: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, status.Errorf(grpcCodes.Internal, "failed to start basket: %s", err)
	}

	resp := &basketspb.CheckoutBasketResponse{}

	return resp, nil
}
