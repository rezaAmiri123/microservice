package v1

import (
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/service_api/internal/app"
	"github.com/rezaAmiri123/microservice/service_api/internal/metrics"
	apimiddleware "github.com/rezaAmiri123/microservice/service_api/internal/ports/http/middleware"
	"github.com/rezaAmiri123/test-microservice/docs"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const (
	maxHeaderBytes = 1 << 20
	stackSize      = 1 << 10 // 1 KB
	bodyLimit      = "2M"
	readTimeout    = 15 * time.Second
	writeTimeout   = 15 * time.Second
	gzipLevel      = 5
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
// @securityDefinitions.apikey ApiKeyAuth

type HttpServer struct {
	app     *app.Application
	metrics *metrics.ApiServiceMetric
	// authClient auth.AuthClient
	validate *validator.Validate
	log      logger.Logger
}

func NewHttpServer(
	debug bool,
	application *app.Application,
	metrics *metrics.ApiServiceMetric,
	log logger.Logger,
	// authClient auth.AuthClient,
) (*echo.Echo, error) {
	httpServer := &HttpServer{
		app:      application,
		metrics:  metrics,
		validate: validator.New(),
		log:      log,
		// authClient: authClient,
	}
	mw := apimiddleware.NewMiddlewareManager(log, application)
	//router := newEchoRouter(httpServer)
	e := echo.New()

	e.Server.ReadTimeout = readTimeout
	e.Server.WriteTimeout = writeTimeout
	e.Server.MaxHeaderBytes = maxHeaderBytes

	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "API Gateway"
	docs.SwaggerInfo.Description = "API Gateway microservices."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Use(middleware.Logger())
	//e.Use(mw.RequestLoggerMiddleware)
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         stackSize,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.BodyLimit(bodyLimit))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID},
	}))
	// if debug {
	// 	e.Use(mw.DebugMiddleware)
	// }
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().URL.Path, "swagger") {
				return true
			}
			return false
		},
	}))

	//e.Use(mw.RateLimitMiddleware())

	v1 := e.Group("/api/v1")

	userGroup := v1.Group("/users")
	userGroup.POST("/register", httpServer.CreateUser())
	userGroup.POST("/login", httpServer.Login())

	financeGroup := v1.Group("/finance")
	financeGroup.POST("/account/create", httpServer.CreateAccount(), mw.AuthMiddleware)
	financeGroup.POST("/transfer/create", httpServer.CreateTransfer(), mw.AuthMiddleware)
	financeGroup.GET("/transfer/list", httpServer.GetTransfers(), mw.AuthMiddleware)

	return e, nil
}

func (h *HttpServer) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
	h.metrics.ErrorHttpRequests.Inc()
}
