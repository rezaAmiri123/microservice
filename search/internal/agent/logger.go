package agent

import (
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/logger/applogger"
	"github.com/rezaAmiri123/microservice/search/internal/constants"
)

func (a *Agent) setupLogger() error {
	appLogger := applogger.NewAppLogger(applogger.Config{
		LogLevel:   a.LogLevel,
		LogDevMode: a.LogDevMode,
		LogEncoder: a.LogEncoder,
	})
	appLogger.InitLogger()
	appLogger.WithName(constants.ServiceName)
	a.container.AddSingleton(constants.LoggerKey, func(c di.Container) (any, error) {
		return appLogger, nil
	})

	return nil
}
