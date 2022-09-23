package command

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/service_user/internal/domain/user"
	"github.com/rezaAmiri123/microservice/service_user/internal/utils"
)

type UpdateUserHandler struct {
	logger logger.Logger
	repo   user.Repository
}

func NewUpdateUserHandler(userRepo user.Repository, logger logger.Logger) *UpdateUserHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}
	return &UpdateUserHandler{repo: userRepo, logger: logger}
}

func (h UpdateUserHandler) Handle(ctx context.Context, arg *user.UpdateUserParams, username string) (*user.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UpdateUserHandler.Handle")
	defer span.Finish()

	if arg.Password != "" {
		hashedPassword, err := utils.HashPassword(arg.Password)
		if err != nil {
			h.logger.Errorf("connot hash the password: %v", err)
			return &user.User{}, fmt.Errorf("connot hash the password: %w", err)
		}
		arg.Password = hashedPassword
	}

	u, err := h.repo.UpdateUser(ctx, arg, username)
	if err != nil {
		h.logger.Errorf("connot create user: %v", err)
		return &user.User{}, fmt.Errorf("connot create user: %w", err)
	}

	// remove the password from response
	u.Password = ""
	return u, nil
}
