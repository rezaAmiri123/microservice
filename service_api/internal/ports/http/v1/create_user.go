package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_api/internal/domain/api"
)

func (h *HttpServer) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.metrics.CreateUserHttpRequests.Inc()

		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "HttpServer.CreateUser")
		defer span.Finish()

		req := &api.CreateUserRequest{}
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

		res, err := h.app.Commands.CreateUser.Handle(ctx, req)
		if err != nil {
			h.log.WarnMsg("CreateUser", err)
			h.traceErr(span, err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		h.metrics.SuccessHttpRequests.Inc()
		return c.JSON(http.StatusOK, res)
	}
}
