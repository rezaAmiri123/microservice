package agent

import (
	"github.com/rezaAmiri123/microservice/pkg/logger/applogger"
)

func (a *Agent) setupLogger() error {
	appLogger := applogger.NewAppLogger(applogger.Config{
		LogLevel:   a.LogLevel,
		LogDevMode: a.LogDevMode,
		LogEncoder: a.LogEncoder,
	})
	appLogger.InitLogger()
	appLogger.WithName("ApiService")
	a.logger = appLogger
	return nil
}
