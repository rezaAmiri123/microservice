package handlers

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
	"github.com/rezaAmiri123/microservice/users/userspb"
)

func RegisterDomainEventHandlers(eventHandlers ddd.EventHandler[ddd.AggregateEvent], domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent]) {
	domainSubscriber.Subscribe(eventHandlers,
		domain.UserRegisteredEvent,
		domain.UserEnabledEvent,
		domain.UserDisabledEvent,
	)
}

type DomainEventHandlers[T ddd.AggregateEvent] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*DomainEventHandlers[ddd.AggregateEvent])(nil)

func NewDomainEventHandlers(publisher am.MessagePublisher[ddd.Event]) *DomainEventHandlers[ddd.AggregateEvent] {
	return &DomainEventHandlers[ddd.AggregateEvent]{
		publisher: publisher,
	}
}

func (h DomainEventHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.UserRegisteredEvent:
		return h.onUserRegistered(ctx, event)
	}
	return nil
}
func (h DomainEventHandlers[T]) onUserRegistered(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.UserRegistered)
	return h.publisher.Publish(ctx, userspb.UserAggregateChannel,
		ddd.NewEvent(userspb.UserRegisteredEvent, &userspb.UserRegistered{
			Id:       payload.User.ID(),
			Username: payload.User.Username,
		}),
	)
}
