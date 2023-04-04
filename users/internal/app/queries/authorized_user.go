package queries

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
)

type (
	AuthorizeUser struct {
		ID string `json:"id"`
	}
	AuthorizeUserHandler struct {
		logger    logger.Logger
		repo      domain.UserRepository
		publisher ddd.EventPublisher[ddd.AggregateEvent]
	}
)

func NewAuthorizeUserHandler(repo domain.UserRepository, logger logger.Logger, publisher ddd.EventPublisher[ddd.AggregateEvent]) *AuthorizeUserHandler {
	if repo == nil {
		panic("userRepo is nil")
	}
	return &AuthorizeUserHandler{repo: repo, logger: logger, publisher: publisher}
}

func (h AuthorizeUserHandler) Handle(ctx context.Context, cmd AuthorizeUser) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AuthorizeUserHandler.Handle")
	defer span.Finish()

	user, err := h.repo.Find(ctx, cmd.ID)
	if err != nil {
		h.logger.Errorf("cannot find user: %v", err)
		return err
	}

	if err = user.Authorize(); err != nil {
		return err
	}

	// publish domain events
	if err = h.publisher.Publish(ctx, user.Events()...); err != nil {
		h.logger.Errorf("cannot publish user event: %v", err)
		return err
	}

	return nil
}
