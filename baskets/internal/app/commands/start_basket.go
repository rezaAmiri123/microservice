package commands

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rezaAmiri123/microservice/baskets/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
)

type (
	StartBasket struct {
		ID     string
		UserID string
	}

	StartBasketHandler struct {
		baskets   domain.BasketRepository
		publisher ddd.EventPublisher[ddd.Event]
		logger    logger.Logger
	}
)

func NewStartBasketHandler(baskets domain.BasketRepository, publisher ddd.EventPublisher[ddd.Event], logger logger.Logger) *StartBasketHandler {
	return &StartBasketHandler{
		baskets:   baskets,
		publisher: publisher,
		logger:    logger,
	}
}

func (h StartBasketHandler) Handle(ctx context.Context, cmd StartBasket) error {
	//span, ctx := opentracing.StartSpanFromContext(ctx, "StartBasketHandler.Handle")
	//defer span.Finish()

	basket, err := h.baskets.Load(ctx, cmd.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			basket = domain.NewBasket(cmd.ID)
		} else {
			return err
		}

	}
	event, err := basket.Start(cmd.UserID)
	if err != nil {
		return err
	}

	err = h.baskets.Save(ctx, basket)
	if err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
