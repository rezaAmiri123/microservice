package commands

import (
	"context"
	"github.com/rezaAmiri123/microservice/depot/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/stackus/errors"
)

type (
	AssignShoppingList struct {
		ID    string
		BotID string
	}

	AssignShoppingListHandler struct {
		shoppingLists domain.ShoppingListRepository
		publisher     ddd.EventPublisher[ddd.AggregateEvent]
		logger        logger.Logger
	}
)

func NewAssignShoppingListHandler(
	shoppingLists domain.ShoppingListRepository,
	publisher ddd.EventPublisher[ddd.AggregateEvent],
	logger logger.Logger,
) AssignShoppingListHandler {
	return AssignShoppingListHandler{
		shoppingLists: shoppingLists,
		publisher:     publisher,
		logger:        logger,
	}
}

func (h AssignShoppingListHandler) AssignShoppingList(ctx context.Context, cmd AssignShoppingList) error {
	//span, ctx := opentracing.StartSpanFromContext(ctx, "AssignShoppingListHandler.Handle")
	//defer span.Finish()

	list, err := h.shoppingLists.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = list.Assign(cmd.BotID); err != nil {
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
