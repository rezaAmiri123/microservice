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

func (s serverTx) CompleteOrder(ctx context.Context, request *orderingpb.CompleteOrderRequest) (resp *orderingpb.CompleteOrderResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.CompleteOrder(ctx, request)
}

func (s server) CompleteOrder(ctx context.Context, request *orderingpb.CompleteOrderRequest) (*orderingpb.CompleteOrderResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("OrderID", request.GetId()),
		attribute.String("InvoiceID", request.GetInvoiceId()),
	)

	err := s.cfg.App.CompleteOrder(ctx, commands.CompleteOrder{
		ID:        request.GetId(),
		InvoiceID: request.GetInvoiceId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to complete order: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, status.Errorf(grpcCode.Internal, "failed to complete order: %s", err)
	}

	resp := &orderingpb.CompleteOrderResponse{}

	return resp, nil
}
