package commands

import (
	"context"
	"github.com/rezaAmiri123/microservice/depot/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/stackus/errors"
)

type (
	InitiateShopping struct {
		ID string
	}

	InitiateShoppingHandler struct {
		shoppingLists domain.ShoppingListRepository
		publisher     ddd.EventPublisher[ddd.AggregateEvent]
		logger        logger.Logger
	}
)

func NewInitiateShoppingHandler(
	shoppingLists domain.ShoppingListRepository,
	publisher ddd.EventPublisher[ddd.AggregateEvent],
	logger logger.Logger,
) InitiateShoppingHandler {
	return InitiateShoppingHandler{
		shoppingLists: shoppingLists,
		publisher:     publisher,
		logger:        logger,
	}
}

func (h InitiateShoppingHandler) InitiateShopping(ctx context.Context, cmd InitiateShopping) error {
	//span, ctx := opentracing.StartSpanFromContext(ctx, "InitiateShoppingHandler.Handle")
	//defer span.Finish()

	list, err := h.shoppingLists.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = list.Initiate(); err != nil {
		return err
	}
	if err := h.shoppingLists.Update(ctx, list); err != nil {
		return errors.Wrap(err, "scheduling shopping")
	}

	// publish domain events
	if err = h.publisher.Publish(ctx, list.Events()...); err != nil {
		return err
	}

	return nil
}
