package agent

import (
	"github.com/rezaAmiri123/microservice/cosec/internal"
	"github.com/rezaAmiri123/microservice/cosec/internal/constants"
	"github.com/rezaAmiri123/microservice/cosec/internal/domain"
	"github.com/rezaAmiri123/microservice/depot/depotpb"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"github.com/rezaAmiri123/microservice/payments/paymentspb"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/registry/serdes"
	"github.com/rezaAmiri123/microservice/users/userspb"
)

func (a *Agent) setupRegistry() error {
	reg := registry.New()
	if err := registrations(reg); err != nil {
		return err
	}

	if err := orderingpb.Registrations(reg); err != nil {
		return err
	}

	if err := userspb.Registrations(reg); err != nil {
		return err
	}

	if err := depotpb.Registrations(reg); err != nil {
		return err
	}

	if err := paymentspb.Registrations(reg); err != nil {
		return err
	}

	a.container.AddSingleton(constants.RegistryKey, func(c di.Container) (any, error) {
		return reg, nil
	})
	return nil
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Saga data
	if err = serde.RegisterKey(internal.CreateOrderSagaName, domain.CreateOrderData{}); err != nil {
		return err
	}

	return nil
}
