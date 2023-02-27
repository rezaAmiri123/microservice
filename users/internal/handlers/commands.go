package handlers

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/users/internal/app"
	"github.com/rezaAmiri123/microservice/users/userspb"
)

type commandHandlers struct {
	app *app.Application
}

func NewCommandHandlers(app *app.Application) ddd.CommandHandler[ddd.Command] {
	return commandHandlers{
		app: app,
	}
}

func RegisterCommandHandlers(subscriber am.CommandSubscriber, handlers ddd.CommandHandler[ddd.Command]) error {
	cmdMsgHandler := am.CommandMessageHandlerFunc(func(ctx context.Context, cmdMsg am.IncomingCommandMessage) (ddd.Reply, error) {
		return handlers.HandleCommand(ctx, cmdMsg)
	})

	return subscriber.Subscribe(userspb.CommandChannel, cmdMsgHandler, am.MessageFilter{
		userspb.AuthorizeUserCommand,
	}, am.GroupName("user-commands"))
}

func (h commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	switch cmd.CommandName() {
	case userspb.AuthorizeUserCommand:
		return h.doAuthorizeUser(ctx, cmd)
	}
	return nil, nil
}

func (h commandHandlers) doAuthorizeUser(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	//payload := cmd.Payload().(*userspb.AuthorizeUser)
	//return nil, h.app.AuthorizeCustomer(ctx, application.AuthorizeCustomer{ID: payload.GetId()})
	return nil, nil
}
