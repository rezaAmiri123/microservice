package constants

// ServiceName The name of this module/service
const ServiceName = "search"

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
	LoggerKey                   = "logger"
	TracerKey                   = "tracer"
	DatabaseKey                 = "database"
	HttpServerKey               = "httpServer"
	GrpcServerKey               = "rpcServer"
	StreamKey                   = "stream"

	OrdersRepoKey   = "ordersRepo"
	UsersRepoKey    = "usersRepo"
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

	OrdersTableName        = "orders"
	UsersCacheTableName    = "users_cache"
	StoresCacheTableName   = "stores_cache"
	ProductsCacheTableName = "products_cache"
)
