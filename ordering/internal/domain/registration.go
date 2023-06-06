package domain

import (
	"github.com/rezaAmiri123/microservice/pkg/es"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/registry/serdes"
)

func Registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Order
	if err = serde.Register(Order{}, func(v any) error {
		order := v.(*Order)
		order.Aggregate = es.NewAggregate("", OrderAggregate)
		return nil
	}); err != nil {
		return err
	}
	// order events
	if err = serde.Register(OrderCreated{}); err != nil {
		return err
	}
	if err = serde.Register(OrderRejected{}); err != nil {
		return err
	}
	if err = serde.Register(OrderApproved{}); err != nil {
		return err
	}
	if err = serde.Register(OrderCanceled{}); err != nil {
		return err
	}
	if err = serde.Register(OrderReadied{}); err != nil {
		return err
	}
	if err = serde.Register(OrderCompleted{}); err != nil {
		return err
	}
	// order snapshots
	if err = serde.RegisterKey(OrderV1{}.SnapshotName(), OrderV1{}); err != nil {
		return err
	}

	return nil
}
