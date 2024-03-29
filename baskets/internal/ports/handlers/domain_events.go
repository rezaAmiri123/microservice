package handlers

import (
	"context"
	"github.com/rezaAmiri123/microservice/baskets/basketspb"
	"github.com/rezaAmiri123/microservice/baskets/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
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
		domain.BasketStartedEvent,
		domain.BasketCanceledEvent,
		domain.BasketCheckedOutEvent,
	)
}

func (h domainHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		if err != nil {
			span.AddEvent(
				"Encountered an error handling domain event",
				trace.WithAttributes(errorsotel.ErrAttrs(err)...),
			)
		}
		span.AddEvent("Handled domain event", trace.WithAttributes(
			attribute.Int64("TookMS", time.Since(started).Microseconds()),
		))
	}(time.Now())

	span.AddEvent("Handled domain event", trace.WithAttributes(
		attribute.String("Event", event.EventName()),
	))
	switch event.EventName() {
	case domain.BasketStartedEvent:
		return h.onBasketStarted(ctx, event)
	case domain.BasketCheckedOutEvent:
		return h.onBasketCheckout(ctx, event)
	case domain.BasketCanceledEvent:
		return h.onBasketCanceled(ctx, event)
	}
	return nil
}

func (h domainHandlers[T]) onBasketCanceled(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.Basket)
	return h.publisher.Publish(ctx, basketspb.BasketAggregateChannel,
		ddd.NewEvent(basketspb.BasketCanceledEvent, &basketspb.BasketCanceled{
			Id: payload.ID(),
		}),
	)
}

func (h domainHandlers[T]) onBasketStarted(ctx context.Context, event ddd.Event) error {
	basket := event.Payload().(*domain.Basket)
	return h.publisher.Publish(ctx, basketspb.BasketAggregateChannel,
		ddd.NewEvent(basketspb.BasketStartedEvent, &basketspb.BasketStarted{
			Id:     basket.ID(),
			UserId: basket.UserID,
		}),
	)
}

func (h domainHandlers[T]) onBasketCheckout(ctx context.Context, event ddd.Event) error {
	basket := event.Payload().(*domain.Basket)
	items := make([]*basketspb.BasketCheckedOut_Item, 0, len(basket.Items))
	for _, item := range basket.Items {
		items = append(items, &basketspb.BasketCheckedOut_Item{
			StoreId:     item.StoreID,
			ProductId:   item.ProductID,
			StoreName:   item.StoreName,
			ProductName: item.ProductName,
			Price:       item.ProductPrice,
			Quantity:    int32(item.Quantity),
		})
	}

	return h.publisher.Publish(ctx, basketspb.BasketAggregateChannel,
		ddd.NewEvent(basketspb.BasketCheckedOutEvent, &basketspb.BasketCheckedOut{
			Id:        basket.ID(),
			UserId:    basket.UserID,
			PaymentId: basket.PaymentID,
			Items:     items,
		}),
	)
}
