package grpc

import (
	"context"
	"github.com/rezaAmiri123/microservice/ordering/internal/app/queries"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func (s serverTx) GetOrder(ctx context.Context, request *orderingpb.GetOrderRequest) (resp *orderingpb.GetOrderResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.GetOrder(ctx, request)
}

func (s server) GetOrder(ctx context.Context, request *orderingpb.GetOrderRequest) (*orderingpb.GetOrderResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("OrderID", request.GetId()),
	)

	order, err := s.cfg.App.GetOrder(ctx, queries.GetOrder{
		ID: request.GetId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("GetOrder: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	resp := &orderingpb.GetOrderResponse{
		Order: s.orderFromDomain(order),
	}

	return resp, nil
}
