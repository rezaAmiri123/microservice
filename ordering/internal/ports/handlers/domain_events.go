package handlers

import (
	"context"
	"fmt"
	"github.com/rezaAmiri123/microservice/ordering/internal/domain"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
)

type domainHandlers[T ddd.Event] struct {
	publisher am.EventPublisher
}

func NewDomainEventHandlers(publisher am.EventPublisher) ddd.EventHandler[ddd.Event] {
	return domainHandlers[ddd.Event]{publisher: publisher}
}

func RegisterDomainEventHandlers(subscriber ddd.EventSubscriber[ddd.Event], handlers ddd.EventHandler[ddd.Event]) {

	subscriber.Subscribe(handlers,
		domain.OrderCreatedEvent,
		domain.OrderRejectedEvent,
		domain.OrderApprovedEvent,
		domain.OrderReadiedEvent,
		domain.OrderCanceledEvent,
		domain.OrderCompletedEvent,
	)
}

func (h domainHandlers[T]) HandleEvent(ctx context.Context, event T) error {
	switch event.EventName() {
	case domain.OrderCreatedEvent:
		return h.onOrderCreated(ctx, event)
	case domain.OrderReadiedEvent:
		return h.onOrderReadied(ctx, event)
	case domain.OrderCompletedEvent:
		return h.onOrderCompleted(ctx, event)
	case domain.OrderCanceledEvent:
		return h.onOrderCancel(ctx, event)
	}
	return nil
}

func (h domainHandlers[T]) onOrderReadied(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.Order)
	fmt.Println("user id: ", payload.UserID)
	return h.publisher.Publish(ctx, orderingpb.OrderAggregateChannel,
		ddd.NewEvent(orderingpb.OrderReadiedEvent, &orderingpb.OrderReadied{
			Id:        payload.ID(),
			UserId:    payload.UserID,
			PaymentId: payload.PaymentID,
			Total:     payload.GetTotal(),
		}),
	)
}

func (h domainHandlers[T]) onOrderCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.Order)
	items := make([]*orderingpb.OrderCreated_Item, len(payload.Items))
	for i, item := range payload.Items {
		items[i] = &orderingpb.OrderCreated_Item{
			ProductId: item.ProductID,
			StoreId:   item.StoreID,
			Price:     item.Price,
			Quantity:  int32(item.Quantity),
		}
	}

	return h.publisher.Publish(ctx, orderingpb.OrderAggregateChannel,
		ddd.NewEvent(orderingpb.OrderCreatedEvent, &orderingpb.OrderCreated{
			Id:         payload.ID(),
			UserId:     payload.UserID,
			PaymentId:  payload.PaymentID,
			ShoppingId: payload.ShoppingID,
			Items:      items,
		}),
	)
}

func (h domainHandlers[T]) onOrderCompleted(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.Order)
	return h.publisher.Publish(ctx, orderingpb.OrderAggregateChannel,
		ddd.NewEvent(orderingpb.OrderCompletedEvent, &orderingpb.OrderCompleted{
			Id:        payload.ID(),
			UserId:    payload.UserID,
			InvoiceId: payload.InvoiceID,
		}),
	)
}

func (h domainHandlers[T]) onOrderCancel(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.Order)
	return h.publisher.Publish(ctx, orderingpb.OrderAggregateChannel,
		ddd.NewEvent(orderingpb.OrderCanceledEvent, &orderingpb.OrderCanceled{
			Id:        payload.ID(),
			UserId:    payload.UserID,
			PaymentId: payload.PaymentID,
		}),
	)
}
