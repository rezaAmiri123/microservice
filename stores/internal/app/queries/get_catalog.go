package queries

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
)

type GetCatalog struct {
	StoreID string
}

type GetCatalogHandler struct {
	catalog domain.CatalogRepository
	logger  logger.Logger
}

func NewGetCatalogHandler(catalog domain.CatalogRepository, logger logger.Logger) GetCatalogHandler {
	return GetCatalogHandler{catalog: catalog, logger: logger}
}

func (h GetCatalogHandler) GetCatalog(ctx context.Context, query GetCatalog) ([]*domain.CatalogProduct, error) {
	return h.catalog.GetCatalog(ctx, query.StoreID)
}
