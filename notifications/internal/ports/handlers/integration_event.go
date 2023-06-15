package handlers

import (
	"context"
	"github.com/rezaAmiri123/microservice/notifications/internal/app"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
)

type integrationHandlers[T ddd.Event] struct {
	app   app.App
	users app.UserCacheRepository
}

var _ ddd.EventHandler[ddd.Event] = (*integrationHandlers[ddd.Event])(nil)

func NewIntegrationEventHandlers(
	reg registry.Registry,
	app app.App,
	users app.UserCacheRepository,
	mws ...am.MessageHandlerMiddleware,
) am.MessageHandler {
	return am.NewEventHandler(reg, integrationHandlers[ddd.Event]{
		app:   app,
		users: users,
	}, mws...)
}

func RegisterIntegrationEventHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) (err error) {
	_, err = subscriber.Subscribe(userspb.UserAggregateChannel, handlers, am.MessageFilter{
		userspb.UserRegisteredEvent,
	}, am.GroupName("notification-users"))
	if err != nil {
		return err
	}

	_, err = subscriber.Subscribe(orderingpb.OrderAggregateChannel, handlers, am.MessageFilter{
		orderingpb.OrderCreatedEvent,
		orderingpb.OrderReadiedEvent,
		orderingpb.OrderCanceledEvent,
		orderingpb.OrderCompletedEvent,
	}, am.GroupName("notification-orders"))

	return err
}

func (h integrationHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		if err != nil {
			span.AddEvent(
				"Encountered an error handling integration event",
				trace.WithAttributes(errorsotel.ErrAttrs(err)...),
			)
		}
		span.AddEvent("Handled integration event", trace.WithAttributes(
			attribute.Int64("TookMS", time.Since(started).Milliseconds()),
		))
	}(time.Now())

	span.AddEvent("Handling integration event", trace.WithAttributes(
		attribute.String("Event", event.EventName()),
	))

	switch event.EventName() {
	case userspb.UserRegisteredEvent:
		return h.onUserRegistered(ctx, event)
	case orderingpb.OrderCreatedEvent:
		return h.onOrderCreated(ctx, event)
	case orderingpb.OrderReadiedEvent:
		return h.onOrderReadied(ctx, event)
	case orderingpb.OrderCanceledEvent:
		return h.onOrderCanceled(ctx, event)
	}

	return nil
}

func (h integrationHandlers[T]) onUserRegistered(ctx context.Context, event T) (err error) {
	payload := event.Payload().(*userspb.UserRegistered)
	return h.users.Add(ctx, payload.GetId(), payload.GetUsername())
}

func (h integrationHandlers[T]) onOrderCreated(ctx context.Context, event T) (err error) {
	payload := event.Payload().(*orderingpb.OrderCreated)
	return h.app.NotifyOrderCreated(ctx, app.OrderCreated{
		OrderID: payload.GetId(),
		UserID:  payload.GetUserId(),
	})
}

func (h integrationHandlers[T]) onOrderReadied(ctx context.Context, event T) (err error) {
	payload := event.Payload().(*orderingpb.OrderReadied)
	return h.app.NotifyOrderReady(ctx, app.OrderReady{
		OrderID: payload.GetId(),
		UserID:  payload.GetUserId(),
	})
}

func (h integrationHandlers[T]) onOrderCanceled(ctx context.Context, event T) (err error) {
	payload := event.Payload().(*orderingpb.OrderCanceled)
	return h.app.NotifyOrderCanceled(ctx, app.OrderCanceled{
		OrderID: payload.GetId(),
		UserID:  payload.GetUserId(),
	})
}
