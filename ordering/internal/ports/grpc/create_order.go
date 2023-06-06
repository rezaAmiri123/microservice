package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/rezaAmiri123/microservice/ordering/internal/app/commands"
	"github.com/rezaAmiri123/microservice/ordering/internal/domain"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	grpcCode "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s server) CreateOrder(ctx context.Context, request *orderingpb.CreateOrderRequest) (*orderingpb.CreateOrderResponse, error) {
	span := trace.SpanFromContext(ctx)

	id := uuid.New().String()

	span.SetAttributes(
		attribute.String("OrderID", id),
		attribute.String("UserID", request.GetUserId()),
		attribute.String("PaymentID", request.GetPaymentId()),
	)

	items := make([]domain.Item, len(request.GetItems()))
	for i, item := range request.GetItems() {
		items[i] = s.itemToDomain(item)
	}

	err := s.cfg.App.CreateOrder(ctx, commands.CreateOrder{
		ID:        id,
		UserID:    request.GetUserId(),
		PaymentID: request.GetPaymentId(),
		Items:     items,
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to create order: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, status.Errorf(grpcCode.Internal, "failed to create create order: %s", err)
	}

	resp := &orderingpb.CreateOrderResponse{Id: id}

	return resp, nil
}
