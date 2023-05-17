package agent

import (
	"database/sql"
	"github.com/rezaAmiri123/microservice/notifications/internal/app"
	"github.com/rezaAmiri123/microservice/notifications/internal/constants"
	"github.com/rezaAmiri123/microservice/notifications/internal/ports/handlers"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/amotel"
	"github.com/rezaAmiri123/microservice/pkg/amprom"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/pkg/db/postgresotel"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/tm"
)

func (a *Agent) setupEventHandler() error {
	a.container.AddScoped(constants.InboxStoreKey, func(c di.Container) (any, error) {
		db := postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB))
		return postgres.NewInboxStore(constants.InboxTableName, db), nil
	})

	a.container.AddScoped(constants.MessageSubscriberKey, func(c di.Container) (any, error) {
		return am.NewMessageSubscriber(
			c.Get(constants.StreamKey).(am.MessageStream),
			amotel.OtelMessageContextExtractor(),
			amprom.ReceivedMessagesCounter(constants.ServiceName),
		), nil
	})

	a.container.AddScoped(constants.IntegrationEventHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewIntegrationEventHandlers(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.ApplicationKey).(app.App),
			c.Get(constants.UsersRepoKey).(app.UserCacheRepository),
			tm.InboxHandler(c.Get(constants.InboxStoreKey).(tm.InboxStore)),
		), nil
	})

	if err := handlers.RegisterIntegrationEventHandlers(
		a.container.Get(constants.MessageSubscriberKey).(am.MessageSubscriber),
		a.container.Get(constants.IntegrationEventHandlersKey).(am.MessageHandler),
	); err != nil {
		return err
	}

	return nil
}
