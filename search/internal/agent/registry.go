package agent

import (
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/search/internal/constants"
	"github.com/rezaAmiri123/microservice/stores/storespb"
	"github.com/rezaAmiri123/microservice/users/userspb"
)

func (a *Agent) setupRegistry() error {
	reg := registry.New()
	//if err := domain.Registrations(reg); err != nil {
	//	return err
	//}
	if err := userspb.Registrations(reg); err != nil {
		return err
	}
	if err := storespb.Registrations(reg); err != nil {
		return err
	}
	if err := orderingpb.Registrations(reg); err != nil {
		return err
	}

	a.container.AddSingleton(constants.RegistryKey, func(c di.Container) (any, error) {
		return reg, nil
	})
	return nil
}
