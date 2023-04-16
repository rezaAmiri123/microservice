package app

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rezaAmiri123/microservice/baskets/internal/app/commands"
	"github.com/rezaAmiri123/microservice/baskets/internal/constants"
)

type instrumentedApp struct {
	App
	basketsStarted  prometheus.Counter
	basketsCheckout prometheus.Counter
	basketsCanceled prometheus.Counter
}

var _ App = (*instrumentedApp)(nil)

func NewInstrumentedApp(app App) App {
	// Prometheus counters
	basketsStarted := promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.BasketsStartedCount,
	})
	basketsCheckedOut := promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.BasketsCheckedOutCount,
	})
	basketsCanceled := promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.BaksetsCanceledCount,
	})

	return instrumentedApp{
		App:             app,
		basketsStarted:  basketsStarted,
		basketsCheckout: basketsCheckedOut,
		basketsCanceled: basketsCanceled,
	}
}

func (a instrumentedApp) StartBasket(ctx context.Context, start commands.StartBasket) error {
	err := a.App.StartBasket(ctx, start)
	if err != nil {
		return err
	}
	a.basketsStarted.Inc()
	return nil
}

func (a instrumentedApp) CheckoutBasket(ctx context.Context, checkout commands.CheckoutBasket) error {
	err := a.App.CheckoutBasket(ctx, checkout)
	if err != nil {
		return err
	}
	a.basketsCheckout.Inc()
	return nil
}

func (a instrumentedApp) CancelBasket(ctx context.Context, cancel commands.CancelBasket) error {
	err := a.App.CancelBasket(ctx, cancel)
	if err != nil {
		return err
	}
	a.basketsCanceled.Inc()
	return nil
}
