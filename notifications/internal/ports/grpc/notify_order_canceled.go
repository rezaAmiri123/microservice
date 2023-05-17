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

func (s server) NotifyOrderCanceled(ctx context.Context, request *notificationspb.NotifyOrderCanceledRequest) (*notificationspb.NotifyOrderCanceledResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("OrderID", request.GetOrderId()),
		attribute.String("UserID", request.GetUserId()),
	)

	err := s.cfg.App.NotifyOrderCanceled(ctx, app.OrderCanceled{
		OrderID: request.GetOrderId(),
		UserID:  request.GetUserId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("NotifyOrderCanceled: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &notificationspb.NotifyOrderCanceledResponse{}, err
}
