package domain

import "github.com/rezaAmiri123/microservice/pkg/registry"

const (
	StoreCreatedEvent               = "stores.StoreCreated"
	StoreParticipationEnabledEvent  = "stores.StoreParticipationEnabled"
	StoreParticipationDisabledEvent = "stores.StoreParticipationDisabled"
	StoreRebrandedEvent             = "stores.StoreRebranded"
)

type StoreCreated struct {
	Name     string
	Location string
}

var _ registry.Registrable = (*StoreCreated)(nil)

func (StoreCreated) Key() string { return StoreCreatedEvent }

type StoreParticipationToggled struct {
	Participating bool
}

type StoreRebranded struct {
	Name string
}

var _ registry.Registrable = (*StoreRebranded)(nil)

func (StoreRebranded) Key() string { return StoreRebrandedEvent }
