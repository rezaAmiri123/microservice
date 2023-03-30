package handlers

import (
	"context"
	"github.com/rezaAmiri123/microservice/depot/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/stores/storespb"
)

type integrationHandlers[T ddd.Event] struct {
	stores   domain.StoreCacheRepository
	products domain.ProductCacheRepository
}

var _ ddd.EventHandler[ddd.Event] = (*integrationHandlers[ddd.Event])(nil)

// TODO It should have accepted app instate of repository directly
func NewIntegrationEventHandlers(
	reg registry.Registry,
	stores domain.StoreCacheRepository,
	products domain.ProductCacheRepository,
	mws ...am.MessageHandlerMiddleware,
) am.MessageHandler {
	return am.NewEventHandler(reg, integrationHandlers[ddd.Event]{
		stores:   stores,
		products: products,
	}, mws...)
}

func RegisterIntegrationEventHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) (err error) {
	_, err = subscriber.Subscribe(storespb.StoreAggregateChannel, handlers, am.MessageFilter{
		storespb.StoreCreatedEvent,
		storespb.StoreRebrandedEvent,
	}, am.GroupName("depot-stores"))
	if err != nil {
		return err
	}

	_, err = subscriber.Subscribe(storespb.ProductAggregateChannel, handlers, am.MessageFilter{
		storespb.ProductAddedEvent,
		storespb.ProductRebrandedEvent,
		storespb.ProductRemovedEvent,
	}, am.GroupName("depot-products"))
	return err
}

func (h integrationHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	switch event.EventName() {
	case storespb.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	case storespb.ProductAddedEvent:
		return h.onProductAdded(ctx, event)
	}
	return nil
}

func (h integrationHandlers[T]) onStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.StoreCreated)
	return h.stores.Add(ctx, payload.GetId(), payload.GetName(), payload.GetLocation())
}

func (h integrationHandlers[T]) onStoreRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.StoreRebranded)
	return h.stores.Rename(ctx, payload.GetId(), payload.GetName())
}

func (h integrationHandlers[T]) onProductAdded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.ProductAdded)
	return h.products.Add(ctx, payload.GetId(), payload.GetStoreId(), payload.GetName())
}

func (h integrationHandlers[T]) onProductRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.ProductRebranded)
	return h.products.Rebrand(ctx, payload.GetId(), payload.GetName())
}

func (h integrationHandlers[T]) onProductRemoved(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.ProductRemoved)
	return h.products.Remove(ctx, payload.GetId())
}
