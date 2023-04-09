package handlers

import (
	"context"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"github.com/rezaAmiri123/microservice/payments/internal/app"
	"github.com/rezaAmiri123/microservice/payments/internal/app/commands"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/registry"
)

type integrationHandlers[T ddd.Event] struct {
	app *app.Application
}

var _ ddd.EventHandler[ddd.Event] = (*integrationHandlers[ddd.Event])(nil)

func NewIntegrationEventHandlers(reg registry.Registry, app *app.Application, mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewEventHandler(reg, integrationHandlers[ddd.Event]{
		app: app,
	}, mws...)
}

func RegisterIntegrationEventHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) error {
	_, err := subscriber.Subscribe(orderingpb.OrderAggregateChannel, handlers, am.MessageFilter{
		orderingpb.OrderReadiedEvent,
		orderingpb.OrderCanceledEvent,
	}, am.GroupName("payment-orders"))

	return err
}

func (h integrationHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case orderingpb.OrderReadiedEvent:
		return h.onOrderReadied(ctx, event)
	case orderingpb.OrderCanceledEvent:
		return h.onOrderCanceled(ctx, event)
	}

	return nil
}

func (h integrationHandlers[T]) onOrderCanceled(ctx context.Context, event T) error {
	payload := event.Payload().(*orderingpb.OrderCanceled)
	return h.app.Commands.CancelInvoice.Handle(ctx, commands.CancelInvoice{
		ID: payload.GetId(),
	})
}

func (h integrationHandlers[T]) onOrderReadied(ctx context.Context, event T) error {
	payload := event.Payload().(*orderingpb.OrderReadied)
	return h.app.Commands.CreateInvoice.Handle(ctx, commands.CreateInvoice{
		ID:        payload.GetId(),
		OrderID:   payload.GetId(),
		PaymentID: payload.GetPaymentId(),
		Amount:    payload.GetTotal(),
	})
}
