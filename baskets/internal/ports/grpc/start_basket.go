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
	grpcCodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, status.Errorf(grpcCodes.Internal, "failed to start basket: %s", err)
	}

	resp := &basketspb.StartBasketResponse{Id: id}

	return resp, nil
}
