package queries

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
)

type GetProduct struct {
	ID string
}

type GetProductHandler struct {
	catalog domain.CatalogRepository
	logger  logger.Logger
}

func NewGetProductHandler(catalog domain.CatalogRepository, logger logger.Logger) *GetProductHandler {
	return &GetProductHandler{catalog: catalog, logger: logger}
}

func (h GetProductHandler) Handle(ctx context.Context, query GetProduct) (*domain.CatalogProduct, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetProductHandler.Handle")
	defer span.Finish()

	return h.catalog.Find(ctx, query.ID)
}
