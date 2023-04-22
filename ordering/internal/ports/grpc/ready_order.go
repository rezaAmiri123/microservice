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

func (s serverTx) ReadyOrder(ctx context.Context, request *orderingpb.ReadyOrderRequest) (resp *orderingpb.ReadyOrderResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.ReadyOrder(ctx, request)
}

func (s server) ReadyOrder(ctx context.Context, request *orderingpb.ReadyOrderRequest) (*orderingpb.ReadyOrderResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("OrderID", request.GetId()),
	)

	err := s.cfg.App.ReadyOrder(ctx, commands.ReadyOrder{
		ID: request.GetId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to create order: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, status.Errorf(grpcCode.Internal, "failed to ready order: %s", err)
	}

	resp := &orderingpb.ReadyOrderResponse{}

	return resp, nil
}
