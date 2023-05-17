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

func (s server) NotifyOrderReady(ctx context.Context, request *notificationspb.NotifyOrderReadyRequest) (*notificationspb.NotifyOrderReadyResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("OrderID", request.GetOrderId()),
		attribute.String("UserID", request.GetUserId()),
	)

	err := s.cfg.App.NotifyOrderReady(ctx, app.OrderReady{
		OrderID: request.GetOrderId(),
		UserID:  request.GetUserId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("NotifyOrderReady: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &notificationspb.NotifyOrderReadyResponse{}, err
}
