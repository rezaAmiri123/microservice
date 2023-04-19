package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/rezaAmiri123/microservice/baskets/basketspb"
	"github.com/rezaAmiri123/microservice/baskets/internal/app/commands"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func (s serverTx) StartBasket(ctx context.Context, request *basketspb.StartBasketRequest) (resp *basketspb.StartBasketResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.StartBasket(ctx, request)
}

func (s server) StartBasket(ctx context.Context, request *basketspb.StartBasketRequest) (*basketspb.StartBasketResponse, error) {
	span := trace.SpanFromContext(ctx)

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()
	id := uuid.New().String()

	span.SetAttributes(
		attribute.String("BasketID", id),
		attribute.String("UserID", request.GetUserId()),
	)

	err := s.cfg.App.StartBasket(ctx, commands.StartBasket{
		ID:     id,
		UserID: request.GetUserId(),
	})

	if err != nil {
		// s.cfg.Logger.Errorf("failed to start basket: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		//return nil, status.Errorf(codes.Internal, "failed to start basket: %s", err)
	}

	resp := &basketspb.StartBasketResponse{Id: id}

	return resp, nil
}
