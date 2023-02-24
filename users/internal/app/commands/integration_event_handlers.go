package commands

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
	"github.com/rezaAmiri123/microservice/users/userspb"
)

type IntegrationEventHandlers[T ddd.AggregateEvent] struct {
	publisher am.MessagePublisher[ddd.Event]
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*IntegrationEventHandlers[ddd.AggregateEvent])(nil)

func NewIntegrationEventHandlers(publisher am.MessagePublisher[ddd.Event]) *IntegrationEventHandlers[ddd.AggregateEvent] {
	return &IntegrationEventHandlers[ddd.AggregateEvent]{
		publisher: publisher,
	}
}

func (h IntegrationEventHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.UserRegisteredEvent:
		return h.onUserRegistered(ctx, event)
	}
	return nil
}
func (h IntegrationEventHandlers[T]) onUserRegistered(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.UserRegistered)
	return h.publisher.Publish(ctx, userspb.UserAggregateChannel,
		ddd.NewEvent(userspb.UserRegisteredEvent, &userspb.UserRegistered{
			Id:       payload.User.ID(),
			Username: payload.User.Username,
		}),
	)
}
