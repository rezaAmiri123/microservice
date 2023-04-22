package app

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rezaAmiri123/microservice/payments/internal/app/commands"
	"github.com/rezaAmiri123/microservice/payments/internal/constants"
)

type instrumentedApp struct {
	App
	invoicesPaid prometheus.Counter
}

var _ App = (*instrumentedApp)(nil)

var (
	invoicesPaidCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.InvoicesPaidCount,
	})
)

func NewInstrumentedApp(app App) App {
	return instrumentedApp{
		App:          app,
		invoicesPaid: invoicesPaidCount,
	}
}

func (a instrumentedApp) PayInvoice(ctx context.Context, cmd commands.PayInvoice) error {
	err := a.App.PayInvoice(ctx, cmd)
	if err != nil {
		return err
	}
	a.invoicesPaid.Inc()
	return nil
}
