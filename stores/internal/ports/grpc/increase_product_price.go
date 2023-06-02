package grpc

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"github.com/rezaAmiri123/microservice/stores/internal/app/commands"
	"github.com/rezaAmiri123/microservice/stores/storespb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func (s server) IncreaseProductPrice(ctx context.Context, request *storespb.IncreaseProductPriceRequest) (*storespb.IncreaseProductPriceResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("ProductID", request.GetId()),
	)

	err := s.cfg.App.IncreaseProductPrice(ctx, commands.IncreaseProductPrice{
		ID:    request.GetId(),
		Price: request.GetPrice(),
	})
	if err != nil {
		s.cfg.Logger.Errorf("IncreaseProductPrice: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	resp := &storespb.IncreaseProductPriceResponse{}

	return resp, nil
}
