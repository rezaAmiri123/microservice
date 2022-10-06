package middleware

import (
	"github.com/rezaAmiri123/microservice/service_api/internal/app"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
)

// Middleware manager
type MiddlewareManager struct {
	logger logger.Logger
	app    *app.Application
	//origins    []string
}

// Middleware manager constructor
func NewMiddlewareManager(logger logger.Logger, app *app.Application) *MiddlewareManager {
	return &MiddlewareManager{logger: logger, app: app}
}
