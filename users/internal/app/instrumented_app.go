package app

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rezaAmiri123/microservice/users/internal/constants"
)

type instrumentedApp struct {
	App
	usersRegistered prometheus.Counter
}

var _ App = (*instrumentedApp)(nil)

var usersRegistered = promauto.NewCounter(prometheus.CounterOpts{
	Name: constants.UsersRegisteredCount,
})

func NewInstrumentedApp(app App) App {
	return instrumentedApp{
		App:             app,
		usersRegistered: usersRegistered,
	}
}

func (a instrumentedApp) RegisterUser(ctx context.Context, cmd RegisterUser) error {
	err := a.App.RegisterUser(ctx, cmd)
	if err != nil {
		return err
	}
	a.usersRegistered.Inc()
	return nil
}
