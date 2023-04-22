package commands

import (
	"context"
	"github.com/rezaAmiri123/microservice/depot/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/stackus/errors"
)

type (
	OrderItem struct {
		StoreID   string
		ProductID string
		Quantity  int
	}

	CreateShoppingList struct {
		ID      string
		OrderID string
		Items   []OrderItem
	}

	CreateShoppingListHandler struct {
		shoppingLists domain.ShoppingListRepository
		stores        domain.StoreRepository
		products      domain.ProductRepository
		publisher     ddd.EventPublisher[ddd.AggregateEvent]
		logger        logger.Logger
	}
)

func NewCreateShoppingListHandler(
	shoppingLists domain.ShoppingListRepository,
	stores domain.StoreRepository,
	products domain.ProductRepository,
	publisher ddd.EventPublisher[ddd.AggregateEvent],
	logger logger.Logger,
) CreateShoppingListHandler {
	return CreateShoppingListHandler{
		shoppingLists: shoppingLists,
		stores:        stores,
		products:      products,
		publisher:     publisher,
		logger:        logger,
	}
}

func (h CreateShoppingListHandler) CreateShoppingList(ctx context.Context, cmd CreateShoppingList) (err error) {
	//span, ctx := opentracing.StartSpanFromContext(ctx, "CreateShoppingListHandler.Handle")
	//defer span.Finish()

	list := domain.CreateShoppingList(cmd.ID, cmd.OrderID)
	for _, item := range cmd.Items {
		// horribly inefficient
		store, err := h.stores.Find(ctx, item.StoreID)
		if err != nil {
			return errors.Wrap(err, "building shopping list")
		}
		product, err := h.products.Find(ctx, item.ProductID)
		if err != nil {
			return errors.Wrap(err, "building shopping list")
		}

		err = list.AddItem(store, product, item.Quantity)
		if err != nil {
			return errors.Wrap(err, "building shopping list")
		}
	}
	if err := h.shoppingLists.Save(ctx, list); err != nil {
		return errors.Wrap(err, "scheduling shopping")
	}

	// publish domain events
	if err = h.publisher.Publish(ctx, list.Events()...); err != nil {
		return err
	}

	return nil
}
