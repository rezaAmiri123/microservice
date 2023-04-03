package handlers

import (
	"context"
	"github.com/rezaAmiri123/microservice/baskets/basketspb"
	"github.com/rezaAmiri123/microservice/depot/depotpb"
	"github.com/rezaAmiri123/microservice/ordering/internal/app"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/registry"
)

type integrationHandlers[T ddd.Event] struct {
	app *app.Application
}

var _ ddd.EventHandler[ddd.Event] = (*integrationHandlers[ddd.Event])(nil)

func NewIntegrationEventHandlers(reg registry.Registry, app *app.Application, mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewEventHandler(reg, integrationHandlers[ddd.Event]{app: app}, mws...)
}

func RegisterIntegrationEventHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) (err error) {
	_, err = subscriber.Subscribe(basketspb.BasketAggregateChannel, handlers, am.MessageFilter{
		basketspb.BasketCheckedOutEvent,
	}, am.GroupName("ordering-baskets"))
	if err != nil {
		return err
	}

	_, err = subscriber.Subscribe(depotpb.ShoppingListAggregateChannel, handlers, am.MessageFilter{
		depotpb.ShoppingListCompletedEvent,
	}, am.GroupName("ordering-depot"))
	return err
}

func (h integrationHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	//switch event.EventName() {
	//case basketspb.BasketCheckedOutEvent:
	//	return h.onBasketCheckedOut(ctx, event)
	//case depotpb.ShoppingListCompletedEvent:
	//	return h.onShoppingListCompleted(ctx, event)
	//}

	return nil
}
