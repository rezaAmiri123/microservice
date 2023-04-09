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
	CancelBasket struct {
		ID string
	}

	CancelBasketHandler struct {
		baskets   domain.BasketRepository
		publisher ddd.EventPublisher[ddd.Event]
		logger    logger.Logger
	}
)

func NewCancelBasketHandler(
	baskets domain.BasketRepository,
	publisher ddd.EventPublisher[ddd.Event],
	logger logger.Logger,
) *CancelBasketHandler {
	return &CancelBasketHandler{
		baskets:   baskets,
		publisher: publisher,
		logger:    logger,
	}
}

func (h CancelBasketHandler) Handle(ctx context.Context, cmd CancelBasket) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CancelBasketHandler.Handle")
	defer span.Finish()
	basket, err := h.baskets.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := basket.Cancel()
	if err != nil {
		return errors.Wrap(err, "basket cancel")
	}

	if err = h.baskets.Update(ctx, basket); err != nil {
		return errors.Wrap(err, "basket cancel")
	}

	return h.publisher.Publish(ctx, event)

}
