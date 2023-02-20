package commands

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
)

type (
	EnableUser struct {
		ID string `json:"id"`
	}
	EnableUserHandler struct {
		logger    logger.Logger
		repo      domain.UserRepository
		publisher ddd.EventPublisher
	}
)

func NewEnableUserHandler(repo domain.UserRepository, logger logger.Logger, publisher ddd.EventPublisher) *EnableUserHandler {
	if repo == nil {
		panic("userRepo is nil")
	}
	return &EnableUserHandler{repo: repo, logger: logger, publisher: publisher}
}

func (h EnableUserHandler) Handle(ctx context.Context, reg EnableUser) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "EnableUserHandler.Handle")
	defer span.Finish()

	u, err := h.repo.Find(ctx, reg.ID)
	if err != nil {
		return err
	}

	if err = u.Enable(); err != nil {
		return err
	}
	if err = h.repo.Update(ctx, u); err != nil {
		return err
	}

	// publish domain events
	if err = h.publisher.Publish(ctx, u.GetEvents()...); err != nil {
		return err
	}

	return nil
}
