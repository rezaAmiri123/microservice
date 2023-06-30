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

	CommandChannel = "mallbots.users.commands"

	AuthorizeUserCommand = "usersapi.AuthorizeUser"
)

func Registrations(reg registry.Registry) error {
	serde := serdes.NewProtoSerde(reg)

	// User events
	if err := serde.Register(&UserRegistered{}); err != nil {
		return err
	}
	if err := serde.Register(&UserEnabled{}); err != nil {
		return err
	}
	if err := serde.Register(&UserDisabled{}); err != nil {
		return err
	}

	// commands
	if err := serde.Register(&AuthorizeUser{}); err != nil {
		return err
	}
	return nil
}

func (*UserRegistered) Key() string { return UserRegisteredEvent }
func (*UserEnabled) Key() string    { return UserEnabledEvent }
func (*UserDisabled) Key() string   { return UserDisabledEvent }

func (*AuthorizeUser) Key() string { return AuthorizeUserCommand }
