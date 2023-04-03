package domain

import (
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/es"
	"github.com/stackus/errors"
)

const OrderAggregate = "ordering.Order"

var (
	ErrOrderAlreadyCreated    = errors.Wrap(errors.ErrBadRequest, "the order cannot be recreated")
	ErrOrderHasNoItems        = errors.Wrap(errors.ErrBadRequest, "the order has no items")
	ErrOrderCannotBeCancelled = errors.Wrap(errors.ErrBadRequest, "the order cannot be cancelled")
	ErrUserIDCannotBeBlank    = errors.Wrap(errors.ErrBadRequest, "the customer id cannot be blank")
	ErrPaymentIDCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the payment id cannot be blank")
)

type Order struct {
	es.Aggregate
	UserID     string
	PaymentID  string
	InvoiceID  string
	ShoppingID string
	Items      []Item
	Status     OrderStatus
}

var _ interface {
	es.EventApplier
	es.Snapshotter
} = (*Order)(nil)

func NewOrder(id string) *Order {
	return &Order{
		Aggregate: es.NewAggregate(id, OrderAggregate),
	}
}

func (Order) Key() string { return OrderAggregate }

func (o *Order) CreateOrder(userID, paymentID string, items []Item) (ddd.Event, error) {
	if o.Status != OrderUnknown {
		return nil, ErrOrderAlreadyCreated
	}

	if len(items) == 0 {
		return nil, ErrOrderHasNoItems
	}

	if userID == "" {
		return nil, ErrUserIDCannotBeBlank
	}

	if paymentID == "" {
		return nil, ErrPaymentIDCannotBeBlank
	}

	o.AddEvent(OrderCreatedEvent, &OrderCreated{
		UserID:    userID,
		PaymentID: paymentID,
		Items:     items,
	})
	o.UserID = userID
	o.PaymentID = paymentID
	o.Items = items
	o.Status = OrderIsPending

	return ddd.NewEvent(OrderCreatedEvent, o), nil
}

func (o *Order) Reject() (ddd.Event, error) {
	// validate status

	o.AddEvent(OrderRejectedEvent, &OrderRejected{})

	return ddd.NewEvent(OrderRejectedEvent, o), nil
}

func (o *Order) Approve(shoppingID string) (ddd.Event, error) {
	// validate status

	o.AddEvent(OrderApprovedEvent, &OrderApproved{
		ShoppingID: shoppingID,
	})

	return ddd.NewEvent(OrderApprovedEvent, o), nil
}

func (o *Order) Cancel() (ddd.Event, error) {
	if o.Status != OrderIsPending {
		return nil, ErrOrderCannotBeCancelled
	}

	o.AddEvent(OrderCanceledEvent, &OrderCanceled{
		UserID:    o.UserID,
		PaymentID: o.PaymentID,
	})
	return ddd.NewEvent(OrderCanceledEvent, o), nil
}

func (o *Order) Ready() (ddd.Event, error) {
	// validate status

	o.AddEvent(OrderReadiedEvent, &OrderReadied{
		UserID:    o.UserID,
		PaymentID: o.PaymentID,
		Total:     o.GetTotal(),
	})

	return ddd.NewEvent(OrderReadiedEvent, o), nil
}

func (o *Order) Complete(invoiceID string) (ddd.Event, error) {
	// validate invoice exists

	// validate status

	o.AddEvent(OrderCompletedEvent, &OrderCompleted{
		UserID:    o.UserID,
		InvoiceID: invoiceID,
	})

	return ddd.NewEvent(OrderCompletedEvent, o), nil
}

func (o Order) GetTotal() float64 {
	var total float64

	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}

	return total
}
func (o *Order) ApplyEvent(event ddd.Event) error {

	switch payload := event.Payload().(type) {
	case *OrderCreated:
		o.UserID = payload.UserID
		o.PaymentID = payload.PaymentID
		o.Items = payload.Items
		o.Status = OrderIsPending

	case *OrderRejected:
		o.Status = OrderIsRejected

	case *OrderApproved:
		o.ShoppingID = payload.ShoppingID
		o.Status = OrderIsApproved

	case *OrderCanceled:
		o.Status = OrderIsCancelled

	case *OrderReadied:
		o.Status = OrderIsReady

	case *OrderCompleted:
		o.InvoiceID = payload.InvoiceID
		o.Status = OrderIsCompleted

	default:
		return errors.ErrInternal.Msgf("%T received the event %s with unexpected payload %T", o, event.EventName(), payload)
	}

	return nil
}
func (o *Order) ApplySnapshot(snapshot es.Snapshot) error {
	switch ss := snapshot.(type) {
	case *OrderV1:
		o.UserID = ss.UserID
		o.PaymentID = ss.PaymentID
		o.InvoiceID = ss.InvoiceID
		o.ShoppingID = ss.ShoppingID
		o.Items = ss.Items
		o.Status = ss.Status

	default:
		return errors.ErrInternal.Msgf("%T received the unexpected snapshot %T", o, snapshot)
	}

	return nil
}

func (o *Order) ToSnapshot() es.Snapshot {
	return &OrderV1{
		UserID:     o.UserID,
		PaymentID:  o.PaymentID,
		InvoiceID:  o.InvoiceID,
		ShoppingID: o.ShoppingID,
		Items:      o.Items,
		Status:     o.Status,
	}
}
