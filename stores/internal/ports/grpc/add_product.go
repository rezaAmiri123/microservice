package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"github.com/rezaAmiri123/microservice/stores/internal/app/commands"
	"github.com/rezaAmiri123/microservice/stores/storespb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func (s server) AddProduct(ctx context.Context, request *storespb.AddProductRequest) (*storespb.AddProductResponse, error) {
	span := trace.SpanFromContext(ctx)

	id := uuid.New().String()

	span.SetAttributes(
		attribute.String("ProductID", id),
	)

	err := s.cfg.App.AddProduct(ctx, commands.AddProduct{
		ID:          id,
		StoreID:     request.GetStoreId(),
		Name:        request.GetName(),
		Description: request.GetDescription(),
		SKU:         request.GetSku(),
		Price:       request.GetPrice(),
	})
	if err != nil {
		s.cfg.Logger.Errorf("AddProduct: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	resp := &storespb.AddProductResponse{Id: id}

	return resp, nil
}
