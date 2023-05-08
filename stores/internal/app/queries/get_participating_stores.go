package queries

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
)

type GetParticipatingStores struct{}

type GetParticipatingStoresHandler struct {
	mall   domain.MallRepository
	logger logger.Logger
}

func NewGetParticipatingStoresHandler(mall domain.MallRepository, logger logger.Logger) GetParticipatingStoresHandler {
	return GetParticipatingStoresHandler{mall: mall, logger: logger}
}

func (h GetParticipatingStoresHandler) GetParticipatingStores(ctx context.Context, _ GetParticipatingStores) ([]*domain.MallStore, error) {
	return h.mall.AllParticipating(ctx)
}
