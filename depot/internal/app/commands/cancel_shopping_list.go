package commands

import (
	"context"
	"github.com/rezaAmiri123/microservice/depot/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
)

type (
	CancelShoppingList struct {
		ID string
	}

	CancelShoppingListHandler struct {
		shoppingLists domain.ShoppingListRepository
		publisher     ddd.EventPublisher[ddd.AggregateEvent]
		logger        logger.Logger
	}
)

func NewCancelShoppingListHandler(
	shoppingLists domain.ShoppingListRepository,
	publisher ddd.EventPublisher[ddd.AggregateEvent],
	logger logger.Logger,
) CancelShoppingListHandler {
	return CancelShoppingListHandler{
		shoppingLists: shoppingLists,
		publisher:     publisher,
		logger:        logger,
	}
}

func (h CancelShoppingListHandler) CancelShoppingList(ctx context.Context, cmd CancelShoppingList) error {
	list, err := h.shoppingLists.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = list.Cancel(); err != nil {
		return err
	}
	if err = h.shoppingLists.Update(ctx, list); err != nil {
		return err
	}

	// publish domain events
	if err = h.publisher.Publish(ctx, list.Events()...); err != nil {
		return err
	}

	return nil
}
