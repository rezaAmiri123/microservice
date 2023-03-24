package handlers

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rezaAmiri123/microservice/payments/internal/app"
	"github.com/rezaAmiri123/microservice/payments/paymentspb"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/registry"
)

type commandHandlers struct {
	app app.Application
}

func NewCommandHandlers(reg registry.Registry, app app.Application, replyPublisher am.ReplyPublisher, mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewCommandHandler(reg, replyPublisher, commandHandlers{
		app: app,
	})
}

func RegisterCommandHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) error {
	_, err := subscriber.Subscribe(paymentspb.CommandChannel, handlers, am.MessageFilter{
		paymentspb.ConfirmPaymentCommand,
	}, am.GroupName("payment-commands"))
	return err
}

func (h commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (reply ddd.Reply, err error) {
	switch cmd.CommandName() {
	case paymentspb.ConfirmPaymentCommand:
		return h.doConfirmPayment(ctx, cmd)
	}

	return nil, nil

}
func (h commandHandlers) doConfirmPayment(ctx context.Context, cmd ddd.Command) (reply ddd.Reply, err error) {
	payload := cmd.Payload().(*paymentspb.ConfirmPayment)
	_ = payload.Id
	//return nil, h.app.Commands.ConfirmPayment(ctx, application.ConfirmPayment{ID: payload.GetId()})
	return nil, errors.New("not implemented")
}
