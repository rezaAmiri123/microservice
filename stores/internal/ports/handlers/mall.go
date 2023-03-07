package handlers

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/stores/internal/constants"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
)

type mallHandlers[T ddd.Event] struct {
	mall domain.MallRepository
}

var _ ddd.EventHandler[ddd.Event] = (*mallHandlers[ddd.Event])(nil)

func NewMallHandlers(mall domain.MallRepository) ddd.EventHandler[ddd.Event] {
	return mallHandlers[ddd.Event]{
		mall: mall,
	}
}

func RegisterMallHandlers(subscriber ddd.EventSubscriber[ddd.Event], handlers ddd.EventHandler[ddd.Event]) {
	subscriber.Subscribe(handlers,
		domain.StoreCreatedEvent,
	)
}

func RegisterMallHandlersTx(container di.Container) {
	handlers := ddd.EventHandlerFunc[ddd.Event](func(ctx context.Context, event ddd.Event) error {
		mallHandlers := di.Get(ctx, constants.MallHandlersKey).(ddd.EventHandler[ddd.Event])

		return mallHandlers.HandleEvent(ctx, event)
	})

	subscriber := container.Get(constants.DomainDispatcherKey).(*ddd.EventDispatcher[ddd.Event])

	RegisterMallHandlers(subscriber, handlers)
}
func (h mallHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	switch event.EventName() {
	case domain.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	}
	return nil
}

func (h mallHandlers[T]) onStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.Store)
	return h.mall.AddStore(ctx, payload.ID(), payload.Name, payload.Location)
}
