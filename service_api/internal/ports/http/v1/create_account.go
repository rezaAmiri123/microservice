package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/auth"
	"github.com/rezaAmiri123/microservice/service_api/internal/domain/api"
)

func (h *HttpServer) CreateAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.metrics.CreateAccountHttpRequests.Inc()

		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "HttpServer.CreateAccount")
		defer span.Finish()

		req := &api.CreateAccountRequest{}
		if err := c.Bind(req); err != nil {
			h.log.WarnMsg("Bind", err)
			h.traceErr(span, err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := h.validate.StructCtx(ctx, req); err != nil {
			h.log.WarnMsg("validate", err)
			h.traceErr(span, err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		payload := auth.PayloadFromCtx(ctx)
		req.OwnerId = payload.UserID

		res, err := h.app.Commands.CreateAccount.Handle(ctx, req)
		if err != nil {
			h.log.WarnMsg("CreateAccount", err)
			h.traceErr(span, err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		h.metrics.SuccessHttpRequests.Inc()
		return c.JSON(http.StatusOK, res)
	}
}
