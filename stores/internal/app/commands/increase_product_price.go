package commands

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
)

type IncreaseProductPrice struct {
	ID    string
	Price float64
}

type IncreaseProductPriceHandler struct {
	products  domain.ProductRepository
	publisher ddd.EventPublisher[ddd.Event]
	logger    logger.Logger
}

func NewIncreaseProductPriceHandler(
	products domain.ProductRepository,
	publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) IncreaseProductPriceHandler {
	return IncreaseProductPriceHandler{
		products:  products,
		publisher: publisher,
		logger:    logger,
	}
}
func (h IncreaseProductPriceHandler) IncreaseProductPrice(ctx context.Context, cmd IncreaseProductPrice) error {
	product, err := h.products.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := product.IncreasePrice(cmd.Price)
	if err != nil {
		return err
	}
	err = h.products.Save(ctx, product)
	if err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
