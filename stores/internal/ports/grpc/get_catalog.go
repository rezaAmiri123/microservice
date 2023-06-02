package grpc

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"github.com/rezaAmiri123/microservice/stores/internal/app/queries"
	"github.com/rezaAmiri123/microservice/stores/storespb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func (s server) GetCatalog(ctx context.Context, request *storespb.GetCatalogRequest) (*storespb.GetCatalogResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("StoreID", request.GetStoreId()),
	)

	products, err := s.cfg.App.GetCatalog(ctx, queries.GetCatalog{
		StoreID: request.GetStoreId(),
	})
	if err != nil {
		s.cfg.Logger.Errorf("GetCatalog: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	protoProducts := make([]*storespb.Product, len(products))
	for i, product := range products {
		protoProducts[i] = s.productFromDomain(product)
	}
	resp := &storespb.GetCatalogResponse{
		Products: protoProducts,
	}

	return resp, nil
}
