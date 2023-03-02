package agent

import (
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/logger/applogger"
)

func (a *Agent) setupLogger() error {
	a.container.AddSingleton("logger", func(c di.Container) (any, error) {
		appLogger := applogger.NewAppLogger(applogger.Config{
			LogLevel:   a.LogLevel,
			LogDevMode: a.LogDevMode,
			LogEncoder: a.LogEncoder,
		})
		appLogger.InitLogger()
		appLogger.WithName("Users")
		//a.logger = appLogger
		return appLogger, nil
	})

	return nil
}
