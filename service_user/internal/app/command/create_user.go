package command

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/opentracing/opentracing-go"
	kafkaClient "github.com/rezaAmiri123/microservice/pkg/kafka"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/pkg/tracing"
	kafkaMessage "github.com/rezaAmiri123/microservice/service_message/proto/kafka"
	"github.com/rezaAmiri123/microservice/service_user/internal/domain/user"
	"github.com/rezaAmiri123/microservice/service_user/internal/utils"
	"github.com/segmentio/kafka-go"
)

type CreateUserHandler struct {
	logger        logger.Logger
	repo          user.Repository
	kafkaProducer kafkaClient.Producer
}

func NewCreateUserHandler(userRepo user.Repository, logger logger.Logger, producer kafkaClient.Producer) *CreateUserHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}
	return &CreateUserHandler{repo: userRepo, logger: logger, kafkaProducer: producer}
}

func (h CreateUserHandler) Handle(ctx context.Context, arg *user.CreateUserParams) (*user.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateUserHandler.Handle")
	defer span.Finish()

	// TODO we need to hash the password
	hashedPassword, err := utils.HashPassword(arg.Password)
	if err != nil {
		h.logger.Errorf("connot hash the password: %v", err)
		return &user.User{}, fmt.Errorf("connot hash the password: %w", err)
	}
	arg.Password = hashedPassword

	u, err := h.repo.CreateUser(ctx, arg)
	if err != nil {
		h.logger.Errorf("connot create user: %v", err)
		return &user.User{}, fmt.Errorf("connot create user: %w", err)
	}

	// remove the password from response
	u.Password = ""

	err = h.sentEvent(ctx, u)
	if err != nil {
		h.logger.Errorf("connot send create user event: %v", err)
	}

	return u, nil
}

func (h CreateUserHandler) sentEvent(ctx context.Context, u *user.User) error {

	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateUserHandler.sentEvent")
	defer span.Finish()

	req := &kafkaMessage.CreateUser{
		UserID:   u.UserID.String(),
		Username: u.Username,
		Email:    u.Email,
		Bio:      u.Bio,
		Image:    u.Image,
	}

	message, err := proto.Marshal(req)
	if err != nil {
		return err
	}
	err = h.kafkaProducer.PublishMessage(ctx, kafka.Message{
		Topic:   kafkaClient.CreateUserTopic,
		Value:   message,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	})
	if err != nil {
		h.logger.Errorf("can not send kafka message %v", err)
	}
	return err
}
