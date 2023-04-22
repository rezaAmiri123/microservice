package grpc

import (
	"context"
	"github.com/rezaAmiri123/microservice/ordering/internal/app/commands"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	grpcCode "google.golang.org/grpc/codes"
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
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("OrderID", request.GetId()),
	)

	err := s.cfg.App.CancelOrder(ctx, commands.CancelOrder{
		ID: request.GetId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to cancel order: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, status.Errorf(grpcCode.Internal, "failed to cancel order: %s", err)
	}

	resp := &orderingpb.CancelOrderResponse{}

	return resp, nil
}
