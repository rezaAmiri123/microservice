package queries

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
)

type GetStores struct{}

type GetStoresHandler struct {
	mall   domain.MallRepository
	logger logger.Logger
}

func NewGetStoresHandler(mall domain.MallRepository, logger logger.Logger) GetStoresHandler {
	return GetStoresHandler{mall: mall, logger: logger}
}

func (h GetStoresHandler) GetStores(ctx context.Context, _ GetStores) ([]*domain.MallStore, error) {
	return h.mall.All(ctx)
}
