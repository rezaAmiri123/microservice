package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/microservice/pkg/auth"
)

func (mw *MiddlewareManager) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationHeader := c.Request().Header.Get(auth.AuthorizationHeaderKey)
		// authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			// ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return c.JSON(http.StatusUnauthorized, err)
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			return c.JSON(http.StatusUnauthorized, err)
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != auth.AuthorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			return c.JSON(http.StatusUnauthorized, err)
		}

		accessToken := fields[1]
		payload, err := mw.app.Commands.LoginVerify.Handle(c.Request().Context(), accessToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}

		c.Set("payload", payload)

		ctx := context.WithValue(c.Request().Context(), auth.AuthorizationPayloadKey, payload)
		c.SetRequest(c.Request().WithContext(ctx))

		// mw.logger.Info(
		// 	"AuthMiddleware, RequestID: %s,  IP: %s, UserID: %s",
		// 	utils.GetRequestID(c),
		// 	utils.GetIPAddress(c),
		// 	user.UUID,
		// )

		return next(c)
	}
}
