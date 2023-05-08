package queries

import (
	"context"
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

func NewGetProductHandler(catalog domain.CatalogRepository, logger logger.Logger) GetProductHandler {
	return GetProductHandler{catalog: catalog, logger: logger}
}

func (h GetProductHandler) GetProduct(ctx context.Context, query GetProduct) (*domain.CatalogProduct, error) {
	return h.catalog.Find(ctx, query.ID)
}
