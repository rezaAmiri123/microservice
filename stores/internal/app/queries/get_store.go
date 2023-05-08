package queries

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
)

type GetStore struct {
	ID string
}

type GetStoreHandler struct {
	mall   domain.MallRepository
	logger logger.Logger
}

func NewGetStoreHandler(mall domain.MallRepository, logger logger.Logger) GetStoreHandler {
	return GetStoreHandler{mall: mall, logger: logger}
}

func (h GetStoreHandler) GetStore(ctx context.Context, query GetStore) (*domain.MallStore, error) {
	return h.mall.Find(ctx, query.ID)
}
