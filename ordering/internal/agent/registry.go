package agent

import (
	"github.com/rezaAmiri123/microservice/baskets/basketspb"
	"github.com/rezaAmiri123/microservice/depot/depotpb"
	"github.com/rezaAmiri123/microservice/ordering/internal/constants"
	"github.com/rezaAmiri123/microservice/ordering/internal/domain"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/registry"
)

func (a *Agent) setupRegistry() error {
	reg := registry.New()
	if err := domain.Registrations(reg); err != nil {
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
