package basketspb

import (
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/registry/serdes"
)

const (
	BasketAggregateChannel = "mallbots.baskets.events.Basket"

	BasketStartedEvent    = "basketsapi.BasketStarted"
	BasketCanceledEvent   = "basketsapi.BasketCanceled"
	BasketCheckedOutEvent = "basketsapi.BasketCheckedOut"
)

func Registration(reg registry.Registry) error {
	return RegistrationWithSerde(serdes.NewProtoSerde(reg))
}

func RegistrationWithSerde(serde registry.Serde) error {
	// basket events
	if err := serde.Register(&BasketStarted{}); err != nil {
		return err
	}
	if err := serde.Register(&BasketCanceled{}); err != nil {
		return err
	}
	if err := serde.Register(&BasketCheckedOut{}); err != nil {
		return err
	}

	return nil
}

// implemented registry.Registrable
func (*BasketStarted) Key() string    { return BasketStartedEvent }
func (*BasketCanceled) Key() string   { return BasketCanceledEvent }
func (*BasketCheckedOut) Key() string { return BasketCheckedOutEvent }
