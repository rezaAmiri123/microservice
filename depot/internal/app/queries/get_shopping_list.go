package queries

import (
	"context"
	"github.com/rezaAmiri123/microservice/depot/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/logger"
)

type (
	GetShoppingList struct {
		ID string
	}

	GetShoppingListHandler struct {
		shoppingLists domain.ShoppingListRepository
		//publisher     ddd.EventPublisher[ddd.AggregateEvent]
		logger logger.Logger
	}
)

func NewGetShoppingListHandler(
	shoppingLists domain.ShoppingListRepository,
	//publisher ddd.EventPublisher[ddd.AggregateEvent],
	logger logger.Logger,
) GetShoppingListHandler {
	return GetShoppingListHandler{
		shoppingLists: shoppingLists,
		//publisher:     publisher,
		logger: logger,
	}
}

func (h GetShoppingListHandler) GetShoppingList(ctx context.Context, query GetShoppingList) (*domain.ShoppingList, error) {
	return h.shoppingLists.Find(ctx, query.ID)
}
