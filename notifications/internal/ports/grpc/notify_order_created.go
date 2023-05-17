package grpc

import (
	"context"
	"github.com/rezaAmiri123/microservice/notifications/internal/app"
	"github.com/rezaAmiri123/microservice/notifications/notificationspb"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func (s server) NotifyOrderCreated(ctx context.Context, request *notificationspb.NotifyOrderCreatedRequest) (*notificationspb.NotifyOrderCreatedResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("OrderID", request.GetOrderId()),
		attribute.String("UserID", request.GetUserId()),
	)

	err := s.cfg.App.NotifyOrderCreated(ctx, app.OrderCreated{
		OrderID: request.GetOrderId(),
		UserID:  request.GetUserId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("NotifyOrderCreated: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &notificationspb.NotifyOrderCreatedResponse{}, err
}
