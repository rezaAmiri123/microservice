package userspb

import (
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/registry/serdes"
)

const (
	UserAggregateChannel = "mallbots.users.events.User"

	UserRegisteredEvent = "usersapi.UserRegistered"
	UserEnabledEvent    = "usersapi.UserEnabled"
	UserDisabledEvent   = "usersapi.UserDisabled"
)

func Registrations(reg registry.Registry) error {
	serde := serdes.NewProtoSerde(reg)

	// Store events
	if err := serde.Register(&UserRegistered{}); err != nil {
		return err
	}
	if err := serde.Register(&UserEnabled{}); err != nil {
		return err
	}
	if err := serde.Register(&UserDisabled{}); err != nil {
		return err
	}
	return nil
}

func (*UserRegistered) Key() string { return UserRegisteredEvent }
func (*UserEnabled) Key() string    { return UserEnabledEvent }
func (*UserDisabled) Key() string   { return UserDisabledEvent }
