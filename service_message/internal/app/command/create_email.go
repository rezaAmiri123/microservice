package command

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/service_message/internal/domain/message"
)

type CreateEmailHandler struct {
	logger logger.Logger
	repo   message.Repository
}

func NewCreateEmailHandler(userRepo message.Repository, logger logger.Logger) *CreateEmailHandler {
	if userRepo == nil {
		panic("repo is nil")
	}
	return &CreateEmailHandler{repo: userRepo, logger: logger}
}

func (h CreateEmailHandler) Handle(ctx context.Context, arg *message.CreateEmailParams) (*message.Email, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateEmailHandler.Handle")
	defer span.Finish()

	e, err := h.repo.CreateEmail(ctx, arg)
	if err != nil {
		h.logger.Errorf("connot create email: %v", err)
		return &message.Email{}, fmt.Errorf("connot create email: %w", err)
	}

	return e, nil
}
