package handlers

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/users/internal/app"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"github.com/stackus/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
)

type commandHandlers struct {
	app    app.App
	logger logger.Logger
}

func NewCommandHandlers(reg registry.Registry, logger logger.Logger, app app.App, replyPublisher am.ReplyPublisher, mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewCommandHandler(reg, replyPublisher, commandHandlers{
		app:    app,
		logger: logger,
	}, mws...)
}

func RegisterCommandHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) error {
	_, err := subscriber.Subscribe(userspb.CommandChannel, handlers, am.MessageFilter{
		userspb.AuthorizeUserCommand,
	}, am.GroupName("user-commands"))
	//})
	return err
}

func (h commandHandlers) HandleCommand(ctx context.Context, cmd ddd.Command) (reply ddd.Reply, err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		if err != nil {
			span.AddEvent(
				"Encountered an error handling command",
				trace.WithAttributes(errorsotel.ErrAttrs(err)...),
			)
			h.logger.Error(errors.Wrap(err, "command commandHandlers.HandleCommand"))
		}
		span.AddEvent("Handled command", trace.WithAttributes(
			attribute.Int64("TookMS", time.Since(started).Milliseconds()),
		))
	}(time.Now())

	span.AddEvent("Handling command", trace.WithAttributes(
		attribute.String("Command", cmd.CommandName()),
	))

	switch cmd.CommandName() {
	case userspb.AuthorizeUserCommand:
		return h.doAuthorizeUser(ctx, cmd)
	}
	return nil, nil
}

func (h commandHandlers) doAuthorizeUser(ctx context.Context, cmd ddd.Command) (ddd.Reply, error) {
	payload := cmd.Payload().(*userspb.AuthorizeUser)
	err := h.app.AuthorizeUser(ctx, app.AuthorizeUser{ID: payload.GetId()})
	if err != nil {
		return nil, errors.Wrap(err, "commandHandlers.doAuthorizeUser ")
	}
	h.logger.Debugf("command handler: authorize user %s", payload.GetId())
	return nil, nil

}
