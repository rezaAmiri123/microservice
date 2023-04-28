package agent

import (
	"database/sql"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/amotel"
	"github.com/rezaAmiri123/microservice/pkg/amprom"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/microservice/pkg/db/postgresotel"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/jetstream"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/tm"
	"github.com/rezaAmiri123/microservice/search/internal/adapters/grpc"
	"github.com/rezaAmiri123/microservice/search/internal/adapters/pg"
	"github.com/rezaAmiri123/microservice/search/internal/app"
	"github.com/rezaAmiri123/microservice/search/internal/constants"
	"github.com/rezaAmiri123/microservice/search/internal/domain"
	"github.com/rezaAmiri123/microservice/search/internal/ports/handlers"
)

func (a *Agent) setupApplication() error {
	dbConn, err := postgres.NewDB(postgres.Config{
		PGDriver:     a.PGDriver,
		PGHost:       a.PGHost,
		PGPort:       a.PGPort,
		PGUser:       a.PGUser,
		PGDBName:     a.PGDBName,
		PGPassword:   a.PGPassword,
		PGSearchPath: a.PGSearchPath,
	})
	if err != nil {
		return fmt.Errorf("cannot load db: %w", err)
	}
	//if err = dbConn.Ping(); err != nil {
	//	fmt.Println("cannot ping db")
	//	return fmt.Errorf("cannot ping db: %w", err)
	//}

	a.container.AddSingleton(constants.DatabaseKey, func(c di.Container) (any, error) {
		return dbConn, nil
	})

	js, err := a.nats()
	if err != nil {
		return err
	}

	stream := jetstream.NewStream(a.NatsStream, js, a.container.Get(constants.LoggerKey).(logger.Logger))

	a.container.AddScoped(constants.DatabaseTransactionKey, func(c di.Container) (any, error) {
		//return c.Get(constants.DatabaseKey).(*sql.DB).Begin()
		return dbConn.Begin()
	})

	a.container.AddScoped(constants.MessageSubscriberKey, func(c di.Container) (any, error) {
		return am.NewMessageSubscriber(
			stream,
			amotel.OtelMessageContextExtractor(),
			amprom.ReceivedMessagesCounter(constants.ServiceName),
		), nil
	})
	a.container.AddScoped(constants.InboxStoreKey, func(c di.Container) (any, error) {
		tx := postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx))
		return postgres.NewInboxStore(constants.InboxTableName, tx), nil
	})

	//a.container.AddScoped(constants.InboxStoreKey, func(c di.Container) (any, error) {
	//	db := postgresotel.Trace(c.Get(constants.DatabaseKey).(*sql.DB))
	//	return postgres.NewInboxStore(constants.InboxTableName, db), nil
	//})

	a.container.AddScoped(constants.UsersRepoKey, func(c di.Container) (any, error) {
		grpcAddress := fmt.Sprintf("%s:%d", a.Config.GRPCUserClientAddr, a.Config.GRPCUserClientPort)
		return pg.NewUserCacheRepository(
			constants.UsersCacheTableName,
			postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
			grpc.NewUserRepository(grpcAddress, c.Get(constants.LoggerKey).(logger.Logger)),
		), nil
	})

	a.container.AddScoped(constants.OrdersRepoKey, func(c di.Container) (any, error) {
		return pg.NewOrderRepository(
			constants.OrdersTableName,
			postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
		), nil
	})
	a.container.AddScoped(constants.ProductsRepoKey, func(c di.Container) (any, error) {
		grpcAddress := fmt.Sprintf("%s:%d", a.Config.GRPCStoreClientAddr, a.Config.GRPCStoreClientPort)
		return pg.NewProductCacheRepository(
			constants.ProductsCacheTableName,
			postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
			grpc.NewProductRepository(grpcAddress, c.Get(constants.LoggerKey).(logger.Logger)),
		), nil
	})

	a.container.AddScoped(constants.StoresRepoKey, func(c di.Container) (any, error) {
		grpcAddress := fmt.Sprintf("%s:%d", a.Config.GRPCStoreClientAddr, a.Config.GRPCStoreClientPort)
		return pg.NewStoreCacheRepository(
			constants.StoresCacheTableName,
			postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
			grpc.NewStoreRepository(grpcAddress, c.Get(constants.LoggerKey).(logger.Logger)),
		), nil
	})

	a.container.AddScoped(constants.ApplicationKey, func(c di.Container) (any, error) {
		return app.New(
			c.Get(constants.OrdersRepoKey).(domain.OrderRepository),
		), nil
	})
	a.container.AddScoped(constants.IntegrationEventHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewIntegrationEventHandlers(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.OrdersRepoKey).(domain.OrderRepository),
			c.Get(constants.ProductsRepoKey).(domain.ProductCacheRepository),
			c.Get(constants.StoresRepoKey).(domain.StoreCacheRepository),
			c.Get(constants.UsersRepoKey).(domain.UserCacheRepository),
			tm.InboxHandler(c.Get(constants.InboxStoreKey).(tm.InboxStore)),
		), nil
	})

	if err = handlers.RegisterIntegrationEventHandlersTx(a.container); err != nil {
		return err
	}
	return nil
}

func (a *Agent) nats() (nats.JetStreamContext, error) {
	nc, err := nats.Connect(a.NatsURL)
	if err != nil {
		return nil, err
	}
	// defer nc.Close()
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     a.NatsStream,
		Subjects: []string{fmt.Sprintf("%s.>", a.NatsStream)},
	})

	return js, err
}
