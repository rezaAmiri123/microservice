package constants

// ServiceName The name of this module/service
const ServiceName = "ordering"

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

	OrdersRepoKey = "ordersRepo"
)

// Repository Table Names
const (
	OrdersTableName = "orders"

	OutboxTableName    = "outbox"
	InboxTableName     = "inbox"
	EventsTableName    = "events"
	SnapshotsTableName = "snapshots"
	SagasTableName     = "sagas"
)

// Metric Names
const (
	OrdersCreatedCount   = "orders_created_count"
	OrdersCompletedCount = "orders_completed_count"
	OrdersApprovedCount  = "orders_approved_count"
)
