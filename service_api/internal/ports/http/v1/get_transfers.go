package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_api/internal/domain/api"
)

func (h *HttpServer) GetTransfers() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.metrics.GetTransfersHttpRequests.Inc()

		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "HttpServer.GetTransfers")
		defer span.Finish()

		req, err := h.validateGetTransfersRequest(c)
		if err != nil {
			h.log.WarnMsg("validation", err)
			h.traceErr(span, err)
			return echo.NewHTTPError(http.StatusBadGateway, err.Error())
		}

		// payload := auth.PayloadFromCtx(ctx)
		// req.OwnerId = payload.UserID

		res, err := h.app.Queries.GetTransfers.Handle(ctx, req)
		if err != nil {
			h.log.WarnMsg("CreateTransfer", err)
			h.traceErr(span, err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		h.metrics.SuccessHttpRequests.Inc()
		return c.JSON(http.StatusOK, res)
	}
}

func (h *HttpServer) validateGetTransfersRequest(c echo.Context) (*api.GetTransfersRequest, error) {
	req := &api.GetTransfersRequest{}

	if c.QueryParam("page") != "" {
		n, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			return nil, fmt.Errorf("page: %v", err)
		}
		req.Page = int64(n)
	}

	if c.QueryParam("page") != "" {
		n, err := strconv.Atoi(c.QueryParam("size"))
		if err != nil {
			return nil, fmt.Errorf("size: %v", err)
		}
		req.Size = int64(n)
	}

	req.Order = c.QueryParam("order")

	return req, nil
}
