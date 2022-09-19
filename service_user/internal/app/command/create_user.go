package command

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/service_user/internal/domain/user"
	"github.com/rezaAmiri123/microservice/service_user/internal/utils"
)

type CreateUserHandler struct {
	logger logger.Logger
	repo   user.Repository
}

func NewCreateUserHandler(userRepo user.Repository, logger logger.Logger) *CreateUserHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}
	return &CreateUserHandler{repo: userRepo, logger: logger}
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
	return u, nil
}
