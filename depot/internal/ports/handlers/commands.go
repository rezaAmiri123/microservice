package handlers

import (
	"context"
	"github.com/google/uuid"
	"github.com/rezaAmiri123/microservice/depot/depotpb"
	"github.com/rezaAmiri123/microservice/depot/internal/app"
	"github.com/rezaAmiri123/microservice/depot/internal/app/commands"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/registry"
)

type commandHandlers struct {
	app *app.Application
}

func NewCommandHandlers(reg registry.Registry, app *app.Application, replyPublisher am.ReplyPublisher, mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewCommandHandler(reg, replyPublisher, commandHandlers{
		app: app,
	}, mws...)
}

func RegisterCommandHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) error {
	_, err := subscriber.Subscribe(depotpb.CommandChannel, handlers, am.MessageFilter{
		depotpb.CreateShoppingListCommand,
		depotpb.CancelShoppingListCommand,
		depotpb.InitiateShoppingCommand,
	}, am.GroupName("depot-commands"))

	return err
}

func (h commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (reply ddd.Reply, err error) {
	switch cmd.CommandName() {
	case depotpb.CreateShoppingListCommand:
		return h.doCreateShoppingList(ctx, cmd)
	case depotpb.InitiateShoppingCommand:
		return h.doInitiateShopping(ctx, cmd)
	}
	return nil, nil
}

func (h commandHandlers) doCreateShoppingList(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*depotpb.CreateShoppingList)

	id := uuid.New().String()
	items := make([]commands.OrderItem, 0, len(payload.GetItems()))
	for _, item := range payload.GetItems() {
		items = append(items, commands.OrderItem{
			StoreID:   item.GetStoreId(),
			ProductID: item.GetProductId(),
			Quantity:  int(item.GetQuantity()),
		})
	}

	err := h.app.Commands.CreateShoppingList.Handle(ctx, commands.CreateShoppingList{
		ID:      id,
		OrderID: payload.GetOrderId(),
		Items:   items,
	})

	return ddd.NewReply(depotpb.CreatedShoppingListReply, &depotpb.CreatedShoppingList{
		Id: id,
	}), err

}
func (h commandHandlers) doInitiateShopping(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*depotpb.InitiateShopping)
	err := h.app.Commands.InitiateShopping.Handle(ctx, commands.InitiateShopping{
		ID: payload.GetId(),
	})

	// returning nil returns a simple Success or Failure reply; err being nil determines which
	return nil, err
}
