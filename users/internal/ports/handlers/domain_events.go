package handlers

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"github.com/stackus/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
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
	logger    logger.Logger
}

var _ ddd.EventHandler[ddd.AggregateEvent] = (*domainHandlers[ddd.AggregateEvent])(nil)

func NewDomainEventHandlers(publisher am.EventPublisher, logger logger.Logger) ddd.EventHandler[ddd.AggregateEvent] {
	return &domainHandlers[ddd.AggregateEvent]{
		publisher: publisher,
		logger:    logger,
	}
}

func (h domainHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		if err != nil {
			span.AddEvent(
				"Encountered an error handling domain event",
				trace.WithAttributes(errorsotel.ErrAttrs(err)...),
			)
			h.logger.Error(errors.Wrap(err, "handler domainHandlers.HandleEvent"))
		}
		span.AddEvent("Handled domain event", trace.WithAttributes(
			attribute.Int64("TookMS", time.Since(started).Milliseconds()),
		))
	}(time.Now())

	span.AddEvent("Handling domain event", trace.WithAttributes(
		attribute.String("Event", event.EventName()),
	))

	switch event.EventName() {
	case domain.UserRegisteredEvent:
		return h.onUserRegistered(ctx, event)
	case domain.UserEnabledEvent:
		return h.onUserEnabled(ctx, event)
	case domain.UserDisabledEvent:
		return h.onUserDisabled(ctx, event)
	}
	return nil
}
func (h domainHandlers[T]) onUserRegistered(ctx context.Context, event ddd.AggregateEvent) error {
	payload := event.Payload().(*domain.UserRegistered)
	err := h.publisher.Publish(ctx, userspb.UserAggregateChannel,
		ddd.NewEvent(userspb.UserRegisteredEvent, &userspb.UserRegistered{
			Id:       payload.User.ID(),
			Username: payload.User.Username,
		}),
	)
	if err == nil {
		h.logger.Debugf("domain handler: register user %s", payload.User.ID())
	}
	return err
}

func (h domainHandlers[T]) onUserEnabled(ctx context.Context, event ddd.AggregateEvent) error {
	err := h.publisher.Publish(ctx, userspb.UserAggregateChannel,
		ddd.NewEvent(userspb.UserEnabledEvent, &userspb.UserEnabled{
			Id: event.AggregateID(),
		}),
	)
	if err == nil {
		h.logger.Debugf("domain handler: enable user %s", event.AggregateID())
	}
	return err
}

func (h domainHandlers[T]) onUserDisabled(ctx context.Context, event ddd.AggregateEvent) error {
	err := h.publisher.Publish(ctx, userspb.UserAggregateChannel,
		ddd.NewEvent(userspb.UserDisabledEvent, &userspb.UserDisabled{
			Id: event.AggregateID(),
		}),
	)
	if err == nil {
		h.logger.Debugf("domain handler: disable user %s", event.AggregateID())
	}
	return err
}
