package commands

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
)

type RemoveProduct struct {
	ID string
}

type RemoveProductHandler struct {
	products  domain.ProductRepository
	publisher ddd.EventPublisher[ddd.Event]
	logger    logger.Logger
}

func NewRemoveProductHandler(
	products domain.ProductRepository,
	publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) RemoveProductHandler {
	return RemoveProductHandler{
		products:  products,
		publisher: publisher,
		logger:    logger,
	}
}
func (h RemoveProductHandler) RemoveProduct(ctx context.Context, cmd RemoveProduct) error {
	product, err := h.products.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := product.Remove()
	if err != nil {
		return err
	}
	err = h.products.Save(ctx, product)
	if err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
