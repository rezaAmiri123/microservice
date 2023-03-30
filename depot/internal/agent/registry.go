package agent

import (
	"github.com/rezaAmiri123/microservice/depot/depotpb"
	"github.com/rezaAmiri123/microservice/depot/internal/constants"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/stores/storespb"
)

func (a *Agent) setupRegistry() error {
	reg := registry.New()
	if err := storespb.Registrations(reg); err != nil {
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
