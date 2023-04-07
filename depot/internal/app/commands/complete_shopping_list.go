package commands

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/depot/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/stackus/errors"
)

type (
	CompleteShoppingList struct {
		ID string
	}

	CompleteShoppingListHandler struct {
		shoppingLists domain.ShoppingListRepository
		publisher     ddd.EventPublisher[ddd.AggregateEvent]
		logger        logger.Logger
	}
)

func NewCompleteShoppingListHandler(
	shoppingLists domain.ShoppingListRepository,
	publisher ddd.EventPublisher[ddd.AggregateEvent],
	logger logger.Logger,
) *CompleteShoppingListHandler {
	return &CompleteShoppingListHandler{
		shoppingLists: shoppingLists,
		publisher:     publisher,
		logger:        logger,
	}
}

func (h CompleteShoppingListHandler) Handle(ctx context.Context, cmd CompleteShoppingList) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CompleteShoppingListHandler.Handle")
	defer span.Finish()

	list, err := h.shoppingLists.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = list.Complete(); err != nil {
		return err
	}
	if err := h.shoppingLists.Update(ctx, list); err != nil {
		return errors.Wrap(err, "completing shopping")
	}

	// publish domain events
	if err = h.publisher.Publish(ctx, list.Events()...); err != nil {
		return err
	}

	return nil
}
