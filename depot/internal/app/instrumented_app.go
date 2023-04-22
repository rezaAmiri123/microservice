package app

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rezaAmiri123/microservice/depot/internal/app/commands"
	"github.com/rezaAmiri123/microservice/depot/internal/constants"
)

type instrumentedApp struct {
	App
	ShoppingListCreated   prometheus.Counter
	ShoppingListAssigned  prometheus.Counter
	ShoppingListCompleted prometheus.Counter
}

var _ App = (*instrumentedApp)(nil)
var (
	shoppingListCreated = promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.ShoppingListCreatedCount,
	})
	shoppingListAssigned = promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.ShoppingListAssignedCount,
	})
	shoppingListCompleted = promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.ShoppingListCompletedCount,
	})
)

func NewInstrumentedApp(app App) App {
	return instrumentedApp{
		App:                   app,
		ShoppingListCreated:   shoppingListCreated,
		ShoppingListAssigned:  shoppingListAssigned,
		ShoppingListCompleted: shoppingListCompleted,
	}
}

func (a instrumentedApp) CreateShoppingList(ctx context.Context, cmd commands.CreateShoppingList) error {
	err := a.App.CreateShoppingList(ctx, cmd)
	if err != nil {
		return err
	}
	a.ShoppingListCreated.Inc()
	return nil
}

func (a instrumentedApp) AssignShoppingList(ctx context.Context, cmd commands.AssignShoppingList) error {
	err := a.App.AssignShoppingList(ctx, cmd)
	if err != nil {
		return err
	}
	a.ShoppingListAssigned.Inc()
	return nil
}

func (a instrumentedApp) CompleteShoppingList(ctx context.Context, cmd commands.CompleteShoppingList) error {
	err := a.App.CompleteShoppingList(ctx, cmd)
	if err != nil {
		return err
	}
	a.ShoppingListCompleted.Inc()
	return nil
}
