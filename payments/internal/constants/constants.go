package constants

// ServiceName The name of this module/service
const ServiceName = "payments"

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
	LoggerKey                   = "logger"
	DomainEventHandlersKey      = "domainEventHandlers"
	IntegrationEventHandlersKey = "integrationEventHandlers"
	CommandHandlersKey          = "commandHandlers"
	ReplyHandlersKey            = "replyHandlers"
	HttpServerKey               = "httpServer"
	GrpcServerKey               = "grpcServer"
	DatabaseKey                 = "database"

	InvoicesRepoKey = "invoicesRepo"
	PaymentsRepoKey = "paymentsRepo"
)

// Repository Table Names
const (
	OutboxTableName    = "outbox"
	InboxTableName     = "inbox"
	EventsTableName    = "events"
	SnapshotsTableName = "snapshots"
	SagasTableName     = "sagas"

	InvoicesTableName = "invoices"
	PaymentsTableName = "payments"
)
