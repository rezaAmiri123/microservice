package constants

// ServiceName The name of this module/service
const ServiceName = "baskets"

// GRPC Service Names
const (
	StoresServiceName    = "STORES"
	CustomersServiceName = "CUSTOMERS"
)

// Dependency Injection Keys
const (
	RegistryKey                 = "registry"
	DomainDispatcherKey         = "domainDispatcher"
	DatabaseTransactionKey      = "tx"
	MessagePublisherKey         = "messagePublisher"
	MessageSubscriberKey        = "messageSubscriber"
	EventPublisherKey           = "eventPublisher"
	CommandPublisherKey         = "commandPublisher"
	ReplyPublisherKey           = "replyPublisher"
	SagaStoreKey                = "sagaStore"
	InboxStoreKey               = "inboxStore"
	ApplicationKey              = "app"
	DomainEventHandlersKey      = "domainEventHandlers"
	IntegrationEventHandlersKey = "integrationEventHandlers"
	CommandHandlersKey          = "commandHandlers"
	ReplyHandlersKey            = "replyHandlers"
	LoggerKey                   = "loggerKey"
	HttpServerKey               = "httpServer"
	GrpcServerKey               = "grpcServer"
	DatabaseKey                 = "database"
	TracerKey                   = "tracer"

	BasketsRepoKey  = "basketsRepo"
	StoresRepoKey   = "storesRepo"
	ProductsRepoKey = "productsRepo"
)

// Repository Table Names
const (
	OutboxTableName    = "outbox"
	InboxTableName     = "inbox"
	EventsTableName    = "events"
	SnapshotsTableName = "snapshots"
	SagasTableName     = "sagas"

	BasketTableName        = "baskets"
	StoresCacheTableName   = "stores_cache"
	ProductsCacheTableName = "products_cache"
)

// Repository Table Names
//const (
//	OutboxTableName    = ServiceName + ".outbox"
//	InboxTableName     = ServiceName + ".inbox"
//	EventsTableName    = ServiceName + ".events"
//	SnapshotsTableName = ServiceName + ".snapshots"
//	SagasTableName     = ServiceName + ".sagas"
//
//	StoresCacheTableName   = ServiceName + ".stores_cache"
//	ProductsCacheTableName = ServiceName + ".products_cache"
//	BasketTableName        = ServiceName + ".baskets"
//)

// Metric Names
const (
	BasketsStartedCount    = "baskets_started_count"
	BasketsCheckedOutCount = "baskets_checked_out_count"
	BaksetsCanceledCount   = "baskets_canceled_count"
)
