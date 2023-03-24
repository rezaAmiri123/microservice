package handlers

import (
	"context"
	"github.com/rezaAmiri123/microservice/payments/internal/domain"
	"github.com/rezaAmiri123/microservice/payments/paymentspb"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
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
		domain.InvoicePaidEvent,
	)
}

func (h domainHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	switch event.EventName() {
	case domain.InvoicePaidEvent:
		return h.onInvoicePaid(ctx, event)
	}

	return nil
}

func (h domainHandlers[T]) onInvoicePaid(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*domain.InvoicePaid)
	return h.publisher.Publish(ctx, paymentspb.InvoiceAggregateChannel,
		ddd.NewEvent(paymentspb.InvoicePaidEvent, &paymentspb.InvoicePaid{
			Id:      payload.ID,
			OrderId: payload.OrderID,
		}),
	)
}
