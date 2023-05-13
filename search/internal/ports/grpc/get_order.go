package grpc

import (
	"context"
	"database/sql"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"github.com/rezaAmiri123/microservice/search/internal/app"
	"github.com/rezaAmiri123/microservice/search/internal/constants"
	"github.com/rezaAmiri123/microservice/search/searchpb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func (s serverTx) GetOrder(ctx context.Context, request *searchpb.GetOrderRequest) (resp *searchpb.GetOrderResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.GetOrder(ctx, request)
}

func (s server) GetOrder(ctx context.Context, request *searchpb.GetOrderRequest) (*searchpb.GetOrderResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("OrderID", request.GetId()),
	)

	order, err := s.cfg.App.GetOrder(ctx, app.GetOrder{
		OrderID: request.GetId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("GetOrder: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	resp := &searchpb.GetOrderResponse{
		Order: s.orderFromDomain(order),
	}

	return resp, nil
}
