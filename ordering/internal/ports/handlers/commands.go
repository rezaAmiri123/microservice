package handlers

import (
	"context"
	"github.com/rezaAmiri123/microservice/ordering/internal/app"
	"github.com/rezaAmiri123/microservice/ordering/internal/app/commands"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/registry"
)

type commandHandlers struct {
	app app.Application
}

func NewCommandHandlers(reg registry.Registry, app app.Application, replyPublisher am.ReplyPublisher, mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewCommandHandler(reg, replyPublisher, commandHandlers{app: app}, mws...)
}

func RegisterCommandHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) error {
	_, err := subscriber.Subscribe(orderingpb.CommandChannel, handlers, am.MessageFilter{
		orderingpb.RejectOrderCommand,
		orderingpb.ApproveOrderCommand,
	}, am.GroupName("ordering-commands"))

	return err
}

func (h commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (reply ddd.Reply, err error) {
	switch cmd.CommandName() {
	//case orderingpb.RejectOrderCommand:
	//	return h.doRejectOrder(ctx, cmd)
	case orderingpb.ApproveOrderCommand:
		return h.doApproveOrder(ctx, cmd)
	}

	return nil, nil

}

func (h commandHandlers) doApproveOrder(ctx context.Context, cmd ddd.Command) (reply ddd.Reply, err error) {
	payload := cmd.Payload().(*orderingpb.ApproveOrder)

	return nil, h.app.Commands.ApproveOrder.Handle(ctx, commands.ApproveOrder{
		ID:         payload.GetId(),
		ShoppingID: payload.GetShoppingId(),
	})

}
