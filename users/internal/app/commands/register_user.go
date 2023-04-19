package commands

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/users/internal/domain"
)

type (
	RegisterUser struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	RegisterUserHandler struct {
		logger    logger.Logger
		repo      domain.UserRepository
		publisher ddd.EventPublisher[ddd.AggregateEvent]
	}
)

func NewRegisterUserHandler(repo domain.UserRepository, logger logger.Logger, publisher ddd.EventPublisher[ddd.AggregateEvent]) *RegisterUserHandler {
	if repo == nil {
		panic("userRepo is nil")
	}
	return &RegisterUserHandler{repo: repo, logger: logger, publisher: publisher}
}

func (h RegisterUserHandler) Handle(ctx context.Context, reg RegisterUser) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RegisterUserHandler.Handle")
	defer span.Finish()

	u, err := domain.RegisterUser(reg.ID, reg.Username, reg.Password, reg.Email)
	if err != nil {
		//h.logger.Errorf("cannot register user: %v", err)
		return err
	}

	if err = h.repo.Save(ctx, u); err != nil {
		//h.logger.Errorf("cannot register user: %v", err)
		return err
	}

	// publish domain events
	if err = h.publisher.Publish(ctx, u.Events()...); err != nil {
		//h.logger.Errorf("cannot publish user event: %v", err)
		return err
	}

	return nil
}
