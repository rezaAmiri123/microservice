package constants

// ServiceName The name of this module/service
const ServiceName = "depot"

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
	HttpServerKey               = "httpServer"
	GrpcServerKey               = "grpcServer"
	DatabaseKey                 = "database"
	TracerKey                   = "tracer"

	ShoppingListsRepoKey = "shoppingListRepo"
	StoresCacheRepoKey   = "storesCacheRepo"
	ProductsCacheRepoKey = "productsCacheRepo"
)

// Repository Table Names
const (
	OutboxTableName    = "outbox"
	InboxTableName     = "inbox"
	EventsTableName    = "events"
	SnapshotsTableName = "snapshots"
	SagasTableName     = "sagas"

	ShoppingListsTableName = "shopping_lists"
	StoresCacheTableName   = "stores_cache"
	ProductsCacheTableName = "products_cache"
)

// Metric Names
const (
	ShoppingListCreatedCount   = "shopping_list_created_count"
	ShoppingListAssignedCount  = "shopping_list_assigned_count"
	ShoppingListCompletedCount = "shopping_list_completed_count"
)
