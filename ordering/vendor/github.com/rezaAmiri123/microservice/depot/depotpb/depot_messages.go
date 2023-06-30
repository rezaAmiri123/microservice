package depotpb

import (
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/registry/serdes"
)

const (
	ShoppingListAggregateChannel = "mallbots.depot.events.ShoppingList"

	ShoppingListCompletedEvent = "depotapi.ShoppingListCompleted"

	CommandChannel = "mallbots.depot.commands"

	CreateShoppingListCommand = "depotapi.CreateShoppingListCommand"
	CancelShoppingListCommand = "depotapi.CancelShoppingListCommand"
	InitiateShoppingCommand   = "depotapi.InitiateShoppingCommand"

	CreatedShoppingListReply = "depotapi.CreatedShoppingListReply"
)

func Registrations(reg registry.Registry) (err error) {
	return RegistrationsWithSerde(serdes.NewProtoSerde(reg))
}


func RegistrationsWithSerde(serde registry.Serde) (err error) {
	if err = serde.Register(&ShoppingListCompleted{}); err != nil {
		return
	}

	if err = serde.Register(&CreateShoppingList{}); err != nil {
		return err
	}
	if err = serde.Register(&CancelShoppingList{}); err != nil {
		return err
	}
	if err = serde.Register(&InitiateShopping{}); err != nil {
		return err
	}

	if err = serde.Register(&CreatedShoppingList{}); err != nil {
		return err
	}

	return nil
}

// Events
func (*ShoppingListCompleted) Key() string { return ShoppingListCompletedEvent }

// Commands
func (*CreateShoppingList) Key() string { return CreateShoppingListCommand }
func (*CancelShoppingList) Key() string { return CancelShoppingListCommand }
func (*InitiateShopping) Key() string   { return InitiateShoppingCommand }

// Replies
func (*CreatedShoppingList) Key() string { return CreatedShoppingListReply }
