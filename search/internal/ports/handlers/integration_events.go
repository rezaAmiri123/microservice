package handlers

import (
	"context"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/search/internal/domain"
	"github.com/rezaAmiri123/microservice/stores/storespb"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
)

type integrationHandlers[T ddd.Event] struct {
	orders   domain.OrderRepository
	products domain.ProductCacheRepository
	stores   domain.StoreCacheRepository
	users    domain.UserCacheRepository
}

var _ ddd.EventHandler[ddd.Event] = (*integrationHandlers[ddd.Event])(nil)

func NewIntegrationEventHandlers(
	reg registry.Registry,
	orders domain.OrderRepository,
	products domain.ProductCacheRepository,
	stores domain.StoreCacheRepository,
	users domain.UserCacheRepository,
	mws ...am.MessageHandlerMiddleware,
) am.MessageHandler {
	return am.NewEventHandler(reg, integrationHandlers[ddd.Event]{
		users:    users,
		orders:   orders,
		products: products,
		stores:   stores,
	}, mws...)
}

func RegisterIntegrationEventHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) (err error) {
	if _, err = subscriber.Subscribe(userspb.UserAggregateChannel, handlers, am.MessageFilter{
		userspb.UserRegisteredEvent,
	}, am.GroupName("search-users")); err != nil {
		return err
	}

	if _, err = subscriber.Subscribe(orderingpb.OrderAggregateChannel, handlers, am.MessageFilter{
		orderingpb.OrderCreatedEvent,
		orderingpb.OrderReadiedEvent,
		orderingpb.OrderCanceledEvent,
		orderingpb.OrderCompletedEvent,
	}, am.GroupName("search-orders")); err != nil {
		return err
	}

	if _, err = subscriber.Subscribe(storespb.ProductAggregateChannel, handlers, am.MessageFilter{
		storespb.ProductAddedEvent,
		storespb.ProductRebrandedEvent,
		storespb.ProductRemovedEvent,
	}, am.GroupName("search-products")); err != nil {
		return err
	}

	if _, err = subscriber.Subscribe(storespb.StoreAggregateChannel, handlers, am.MessageFilter{
		storespb.StoreCreatedEvent,
		storespb.StoreRebrandedEvent,
	}, am.GroupName("search-stores")); err != nil {
		return err
	}

	return
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

	// ordering
	case orderingpb.OrderCreatedEvent:
		return h.onOrderCreated(ctx, event)
	case orderingpb.OrderReadiedEvent:
		return h.onOrderReadied(ctx, event)
	case orderingpb.OrderCanceledEvent:
		return h.onOrderCanceled(ctx, event)
	case orderingpb.OrderCompletedEvent:
		return h.onOrderCompleted(ctx, event)

	// stores
	case storespb.ProductAddedEvent:
		return h.onProductAdded(ctx, event)
	case storespb.ProductRebrandedEvent:
		return h.onProductRebranded(ctx, event)
	case storespb.ProductRemovedEvent:
		return h.onProductRemoved(ctx, event)
	case storespb.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	case storespb.StoreRebrandedEvent:
		return h.onStoreRebranded(ctx, event)

	}

	return nil
}

func (h integrationHandlers[T]) onProductAdded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.ProductAdded)
	return h.products.Add(ctx, payload.GetId(), payload.GetStoreId(), payload.GetName())
}

func (h integrationHandlers[T]) onProductRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.ProductRebranded)
	return h.products.Rebrand(ctx, payload.GetId(), payload.GetName())
}

func (h integrationHandlers[T]) onProductRemoved(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.ProductRemoved)
	return h.products.Remove(ctx, payload.GetId())
}

func (h integrationHandlers[T]) onStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.StoreCreated)
	return h.stores.Add(ctx, payload.GetId(), payload.GetName())
}

func (h integrationHandlers[T]) onStoreRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.StoreRebranded)
	return h.stores.Rename(ctx, payload.GetId(), payload.GetName())
}

func (h integrationHandlers[T]) onUserRegistered(ctx context.Context, event T) error {
	payload := event.Payload().(*userspb.UserRegistered)
	return h.users.Add(ctx, payload.GetId(), payload.GetUsername())
}

func (h integrationHandlers[T]) onOrderCreated(ctx context.Context, event T) error {
	payload := event.Payload().(*orderingpb.OrderCreated)

	user, err := h.users.Find(ctx, payload.GetUserId())
	if err != nil {
		return err
	}

	var total float64
	items := make([]domain.Item, len(payload.GetItems()))
	seenStores := map[string]*domain.Store{}
	for i, item := range payload.GetItems() {
		product, err := h.products.Find(ctx, item.GetProductId())
		if err != nil {
			return err
		}
		var store *domain.Store
		var exists bool

		if store, exists = seenStores[product.StoreID]; !exists {
			store, err = h.stores.Find(ctx, product.StoreID)
			if err != nil {
				return err
			}
			seenStores[store.ID] = store
		}
		items[i] = domain.Item{
			ProductID:   product.ID,
			StoreID:     store.ID,
			ProductName: product.Name,
			StoreName:   store.Name,
			Price:       item.Price,
			Quantity:    int(item.Quantity),
		}
		total += float64(item.Quantity) * item.Price
	}
	order := &domain.Order{
		OrderID:  payload.GetId(),
		UserID:   user.ID,
		Username: user.Name,
		Items:    items,
		Total:    total,
		Status:   "New",
	}
	return h.orders.Add(ctx, order)
}

func (h integrationHandlers[T]) onOrderReadied(ctx context.Context, event T) error {
	payload := event.Payload().(*orderingpb.OrderReadied)
	return h.orders.UpdateStatus(ctx, payload.GetId(), "Ready For Pickup")
}
func (h integrationHandlers[T]) onOrderCanceled(ctx context.Context, event T) error {
	payload := event.Payload().(*orderingpb.OrderCanceled)
	return h.orders.UpdateStatus(ctx, payload.GetId(), "Canceled")
}
func (h integrationHandlers[T]) onOrderCompleted(ctx context.Context, event T) error {
	payload := event.Payload().(*orderingpb.OrderCanceled)
	return h.orders.UpdateStatus(ctx, payload.GetId(), "Completed")
}
