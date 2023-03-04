package handlers

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
	"github.com/rezaAmiri123/microservice/users/userspb"
)

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.AggregateEvent], handlers ddd.EventHandler[ddd.AggregateEvent]) {
	subscriber.Subscribe(handlers,
		domain.UserRegisteredEvent,
		domain.UserEnabledEvent,
		domain.UserDisabledEvent,
	)
}

type domainHandlers[T ddd.AggregateEvent] struct {
	publisher am.EventPublisher
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*domainHandlers[ddd.AggregateEvent])(nil)

func NewDomainEventHandlers(publisher am.EventPublisher) ddd.EventHandler[ddd.AggregateEvent] {
	return &domainHandlers[ddd.AggregateEvent]{
		publisher: publisher,
	}
}

func (h domainHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.UserRegisteredEvent:
		return h.onUserRegistered(ctx, event)
	}
	return nil
}
func (h domainHandlers[T]) onUserRegistered(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.UserRegistered)
	return h.publisher.Publish(ctx, userspb.UserAggregateChannel,
		ddd.NewEvent(userspb.UserRegisteredEvent, &userspb.UserRegistered{
			Id:       payload.User.ID(),
			Username: payload.User.Username,
		}),
	)
}
