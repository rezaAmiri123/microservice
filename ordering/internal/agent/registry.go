package agent

import (
	"github.com/rezaAmiri123/microservice/baskets/basketspb"
	"github.com/rezaAmiri123/microservice/depot/depotpb"
	"github.com/rezaAmiri123/microservice/ordering/internal/constants"
	"github.com/rezaAmiri123/microservice/ordering/internal/domain"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/es"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/registry/serdes"
)

func (a *Agent) setupRegistry() error {
	reg := registry.New()
	if err := registrations(reg); err != nil {
		return err
	}
	if err := orderingpb.Registrations(reg); err != nil {
		return err
	}
	if err := basketspb.Registration(reg); err != nil {
		return err
	}
	if err := depotpb.Registrations(reg); err != nil {
		return err
	}
	a.container.AddSingleton(constants.RegistryKey, func(c di.Container) (any, error) {
		return reg, nil
	})
	return nil
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Order
	if err = serde.Register(domain.Order{}, func(v any) error {
		order := v.(*domain.Order)
		order.Aggregate = es.NewAggregate("", domain.OrderAggregate)
		return nil
	}); err != nil {
		return err
	}
	// order events
	if err = serde.Register(domain.OrderCreated{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderRejected{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderApproved{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderCanceled{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderReadied{}); err != nil {
		return err
	}
	if err = serde.Register(domain.OrderCompleted{}); err != nil {
		return err
	}
	// order snapshots
	if err = serde.RegisterKey(domain.OrderV1{}.SnapshotName(), domain.OrderV1{}); err != nil {
		return err
	}

	return nil
}
