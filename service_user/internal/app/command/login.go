package command

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	pkgGrpc "github.com/rezaAmiri123/microservice/pkg/grpc"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/pkg/token"
	"github.com/rezaAmiri123/microservice/service_user/internal/domain/user"
	"github.com/rezaAmiri123/microservice/service_user/internal/utils"
)

type LoginHandler struct {
	logger      logger.Logger
	repo        user.Repository
	maker       token.Maker
	makerConfig token.Config
}

func NewLoginHandler(userRepo user.Repository, logger logger.Logger, maker token.Maker, makerConfig token.Config) *LoginHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}
	return &LoginHandler{repo: userRepo, logger: logger, maker: maker, makerConfig: makerConfig}
}

func (h LoginHandler) Handle(ctx context.Context, arg *user.LoginRequestParams) (*user.LoginResponseParams, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateUserHandler.Handle")
	defer span.Finish()

	u, err := h.repo.GetUserByUsername(ctx, arg.Username)
	if err != nil {
		h.logger.Errorf("connot get user: %v", err)
		return &user.LoginResponseParams{}, fmt.Errorf("connot get user: %w", err)
	}

	err = utils.CheckPassword(arg.Password, u.Password)
	if err != nil {
		h.logger.Errorf("password is incorrect: %v", err)
		return &user.LoginResponseParams{}, fmt.Errorf("password is incorrect: %w", err)
	}

	accessToken, accessPayload, err := h.maker.CreateToken(
		u.Username,
		u.UserID.String(),
		h.makerConfig.AccessTokenDuration,
	)
	if err != nil {
		h.logger.Errorf("cannot create access token: %v", err)
		return &user.LoginResponseParams{}, fmt.Errorf("cannot create access token: %w", err)
	}
	refreshToken, refreshPayload, err := h.maker.CreateToken(
		u.Username,
		u.UserID.String(),
		h.makerConfig.RefreshTokenDuration,
	)
	if err != nil {
		h.logger.Errorf("cannot create refresh token: %v", err)
		return &user.LoginResponseParams{}, fmt.Errorf("cannot create refresh token: %w", err)
	}

	mtdt := pkgGrpc.ExtractMetadata(ctx)
	session, err := h.repo.CreateSession(ctx, &user.CreateSessionParams{
		SessionID:    refreshPayload.ID,
		Username:     u.Username,
		RefreshToken: refreshToken,
		UserAgent:    mtdt.UserAgent,
		ClientIp:     mtdt.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		h.logger.Errorf("cannot create session: %v", err)
		return &user.LoginResponseParams{}, fmt.Errorf("cannot create session: %w", err)
	}

	res := &user.LoginResponseParams{
		AccessToken:    accessToken,
		AccessPayload:  accessPayload,
		RefreshToken:   refreshToken,
		RefreshPayload: refreshPayload,
		Session:        session,
		User:           u,
	}

	return res, nil
}
