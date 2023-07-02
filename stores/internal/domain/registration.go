package domain

import (
	"github.com/rezaAmiri123/microservice/pkg/es"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/registry/serdes"
)

func Registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Store
	if err = serde.Register(Store{}, func(v any) error {
		store := v.(*Store)
		store.Aggregate = es.NewAggregate("", StoreAggregate)
		return nil
	}); err != nil {
		return
	}
	// store events
	if err = serde.Register(StoreCreated{}); err != nil {
		return
	}
	if err = serde.RegisterKey(StoreParticipationEnabledEvent, StoreParticipationToggled{}); err != nil {
		return
	}
	if err = serde.RegisterKey(StoreParticipationDisabledEvent, StoreParticipationToggled{}); err != nil {
		return
	}
	if err = serde.Register(StoreRebranded{}); err != nil {
		return
	}
	// store snapshots
	if err = serde.RegisterKey(StoreV1{}.SnapshotName(), StoreV1{}); err != nil {
		return
	}

	// Product
	if err = serde.Register(Product{}, func(v any) error {
		store := v.(*Product)
		store.Aggregate = es.NewAggregate("", ProductAggregate)
		return nil
	}); err != nil {
		return
	}
	// product events
	if err = serde.Register(ProductAdded{}); err != nil {
		return
	}
	if err = serde.Register(ProductRebranded{}); err != nil {
		return
	}
	if err = serde.RegisterKey(ProductPriceIncreasedEvent, ProductPriceChanged{}); err != nil {
		return
	}
	if err = serde.RegisterKey(ProductPriceDecreasedEvent, ProductPriceChanged{}); err != nil {
		return
	}
	if err = serde.Register(ProductRemoved{}); err != nil {
		return
	}
	// product snapshots
	if err = serde.RegisterKey(ProductV1{}.SnapshotName(), ProductV1{}); err != nil {
		return
	}

	return

}
