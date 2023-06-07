package queries

import (
	"context"
	"github.com/rezaAmiri123/microservice/baskets/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/logger"
)

type (
	GetBasket struct {
		ID string
	}

	GetBasketHandler struct {
		baskets domain.BasketRepository
		logger  logger.Logger
	}
)

func NewGetBasketHandler(baskets domain.BasketRepository, logger logger.Logger) GetBasketHandler {
	return GetBasketHandler{
		baskets: baskets,
		logger:  logger,
	}
}

func (h GetBasketHandler) GetBasket(ctx context.Context, query GetBasket) (*domain.Basket, error) {
	return h.baskets.Load(ctx, query.ID)
}
