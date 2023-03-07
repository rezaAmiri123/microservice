package handlers

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
	"github.com/rezaAmiri123/microservice/stores/storespb"
)

type domainHandlers[T ddd.Event] struct {
	publisher am.EventPublisher
}

var _ ddd.EventHandler[ddd.Event] = (*domainHandlers[ddd.Event])(nil)

func NewDomainEventHandlers(publisher am.EventPublisher) ddd.EventHandler[ddd.Event] {
	return &domainHandlers[ddd.Event]{
		publisher: publisher,
	}
}

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.Event], handlers ddd.EventHandler[ddd.Event]) {
	subscriber.Subscribe(handlers,
		domain.StoreCreatedEvent,
		domain.ProductAddedEvent,
	)
}

func (h domainHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	switch event.EventName() {
	case domain.StoreCreatedEvent:
		h.onStoreCreated(ctx, event)
	case domain.ProductAddedEvent:
		h.onProductAdded(ctx, event)
	}
	return nil
}

func (h domainHandlers[T]) onStoreCreated(ctx context.Context, event ddd.Event) error {
	store := event.Payload().(*domain.Store)
	return h.publisher.Publish(ctx, storespb.StoreAggregateChannel,
		ddd.NewEvent(storespb.StoreCreatedEvent, &storespb.StoreCreated{
			Id:       store.ID(),
			Name:     store.Name,
			Location: store.Location,
		}),
	)
}

func (h domainHandlers[T]) onProductAdded(ctx context.Context, event ddd.Event) error {
	product := event.Payload().(*domain.Product)
	return h.publisher.Publish(ctx, storespb.ProductAggregateChannel,
		ddd.NewEvent(storespb.ProductAddedEvent, &storespb.ProductAdded{
			Id:          product.ID(),
			StoreId:     product.StoreID,
			Name:        product.Name,
			Description: product.Description,
			Sku:         product.SKU,
			Price:       product.Price,
		}),
	)
}
