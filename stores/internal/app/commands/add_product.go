package commands

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
	"github.com/stackus/errors"
)

type AddProduct struct {
	ID          string
	StoreID     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

type AddProductHandler struct {
	products  domain.ProductRepository
	publisher ddd.EventPublisher[ddd.Event]
	logger    logger.Logger
}

func NewAddProductHandler(products domain.ProductRepository, publisher ddd.EventPublisher[ddd.Event], logger logger.Logger) *AddProductHandler {
	return &AddProductHandler{
		products:  products,
		publisher: publisher,
		logger:    logger,
	}
}
func (h AddProductHandler) Handle(ctx context.Context, cmd AddProduct) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AddProductHandler.Handle")
	defer span.Finish()

	product, err := h.products.Load(ctx, cmd.ID)
	if err != nil {
		return errors.Wrap(err, "error adding product")
	}

	event, err := product.InitProduct(cmd.ID, cmd.StoreID, cmd.Name, cmd.Description, cmd.SKU, cmd.Price)
	if err != nil {
		return errors.Wrap(err, "initializing product")
	}
	err = h.products.Save(ctx, product)
	if err != nil {
		return errors.Wrap(err, "error adding product")
	}

	if err = h.publisher.Publish(ctx, event); err != nil {
		return errors.Wrap(err, "publishing domain event")
	}

	return nil
}
