package commands

import (
	"context"
	"github.com/rezaAmiri123/microservice/baskets/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
)

type (
	AddItem struct {
		ID        string
		BasketID  string
		ProductID string
		Quantity  int
	}

	AddItemHandler struct {
		baskets   domain.BasketRepository
		stores    domain.StoreRepository
		products  domain.ProductRepository
		publisher ddd.EventPublisher[ddd.Event]
		logger    logger.Logger
	}
)

func NewAddItemHandler(
	baskets domain.BasketRepository,
	stores domain.StoreRepository,
	products domain.ProductRepository,
	publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) AddItemHandler {
	return AddItemHandler{
		baskets:   baskets,
		stores:    stores,
		products:  products,
		publisher: publisher,
		logger:    logger,
	}
}

func (h AddItemHandler) AddItem(ctx context.Context, cmd AddItem) error {
	//span, ctx := opentracing.StartSpanFromContext(ctx, "AddItemHandler.Handle")
	//defer span.Finish()

	basket, err := h.baskets.Load(ctx, cmd.BasketID)
	if err != nil {
		return err
	}
	product, err := h.products.Find(ctx, cmd.ProductID)
	if err != nil {
		return err
	}
	store, err := h.stores.Find(ctx, product.StoreID)
	if err != nil {
		return err
	}

	err = basket.AddItem(store, product, cmd.Quantity)
	if err != nil {
		return err
	}

	return h.baskets.Save(ctx, basket)
}
