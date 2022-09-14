package command

import (
	"context"
	"time"

	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/service_user/internal/domain/user"
)

type CreateUserHandler struct {
	logger   logger.Logger
	userRepo user.Repository
}

func NewCreateUserHandler(userRepo user.Repository, logger logger.Logger) CreateUserHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}
	return CreateUserHandler{userRepo: userRepo, logger: logger}
}

func (h CreateUserHandler) Handle(ctx context.Context, arg *user.CreateUserParams) (*user.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateUserHandler.Handle")
	defer span.Finish()

	if err := user.SetUUID(); err != nil {
		return err
	}
	if err := user.Validate(ctx); err != nil {
		return err
	}
	if err := user.HashPassword(); err != nil {
		return err
	}

	e := &kafkaMessages.Email{
		To:      []string{user.Email},
		From:    "admin@example.com",
		Subject: "register user subject",
		Body:    "register user body",
	}
	msg := &kafkaMessages.CreateEmail{Email: e}

	message, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	err = h.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}
	err = h.kafka.PublishMessage(ctx, kafka.Message{
		Topic: kafkaClient.CreateEmailTopic,
		Value: message,
		Time:  time.Now().UTC(),
	})
	if err != nil {
		h.logger.Errorf("can not send kafka message %v", err)
	}
	return err
}
