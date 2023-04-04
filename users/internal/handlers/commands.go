package handlers

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/users/internal/app"
	"github.com/rezaAmiri123/microservice/users/internal/app/queries"
	"github.com/rezaAmiri123/microservice/users/userspb"
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
	_, err := subscriber.Subscribe(userspb.CommandChannel, handlers, am.MessageFilter{
		userspb.AuthorizeUserCommand,
	}, am.GroupName("user-commands"))
	//})
	return err
}

func (h commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	switch cmd.CommandName() {
	case userspb.AuthorizeUserCommand:
		return h.doAuthorizeUser(ctx, cmd)
	}
	return nil, nil
}

func (h commandHandlers) doAuthorizeUser(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*userspb.AuthorizeUser)
	return nil, h.app.Queries.AuthorizeUser.Handle(ctx, queries.AuthorizeUser{ID: payload.GetId()})

}
