package app

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rezaAmiri123/microservice/ordering/internal/app/commands"
	"github.com/rezaAmiri123/microservice/ordering/internal/constants"
)

type instrumentedApp struct {
	App
	ordersCreated  prometheus.Counter
	ordersComplete prometheus.Counter
	ordersApproved prometheus.Counter
}

var _ App = (*instrumentedApp)(nil)
var (
	ordersCreated = promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.OrdersCreatedCount,
	})
	ordersComplete = promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.OrdersCompletedCount,
	})
	ordersApproved = promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.OrdersApprovedCount,
	})
)

func NewInstrumentedApp(app App) App {
	// Prometheus counters
	//ordersCreated := promauto.NewCounter(prometheus.CounterOpts{
	//	Name: constants.OrdersCreatedCount,
	//})

	return instrumentedApp{
		App:            app,
		ordersCreated:  ordersCreated,
		ordersComplete: ordersComplete,
		ordersApproved: ordersApproved,
	}
}

func (a instrumentedApp) CreateOrder(ctx context.Context, cmd commands.CreateOrder) error {
	err := a.App.CreateOrder(ctx, cmd)
	if err != nil {
		return err
	}
	a.ordersCreated.Inc()
	return nil
}

func (a instrumentedApp) CompleteOrder(ctx context.Context, cmd commands.CompleteOrder) error {
	err := a.App.CompleteOrder(ctx, cmd)
	if err != nil {
		return err
	}
	a.ordersComplete.Inc()
	return nil
}

func (a instrumentedApp) ApproveOrder(ctx context.Context, cmd commands.ApproveOrder) error {
	err := a.App.ApproveOrder(ctx, cmd)
	if err != nil {
		return err
	}
	a.ordersApproved.Inc()
	return nil
}
