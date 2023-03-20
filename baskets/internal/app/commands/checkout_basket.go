package commands

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/baskets/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/stackus/errors"
)

type (
	CheckoutBasket struct {
		ID        string
		PaymentID string
	}

	CheckoutBasketHandler struct {
		baskets   domain.BasketRepository
		publisher ddd.EventPublisher[ddd.Event]
		logger    logger.Logger
	}
)

func NewCheckoutBasketHandler(
	baskets domain.BasketRepository,
	publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) *CheckoutBasketHandler {
	return &CheckoutBasketHandler{
		baskets:   baskets,
		publisher: publisher,
		logger:    logger,
	}
}

func (h CheckoutBasketHandler) Handle(ctx context.Context, cmd CheckoutBasket) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CheckoutBasketHandler.Handle")
	defer span.Finish()
	basket, err := h.baskets.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := basket.Checkout(cmd.PaymentID)
	if err != nil {
		return errors.Wrap(err, "baskets checkout")
	}

	if err = h.baskets.Update(ctx, basket); err != nil {
		return errors.Wrap(err, "basket checkout")
	}

	return h.publisher.Publish(ctx, event)

}
