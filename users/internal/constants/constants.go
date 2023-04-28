package constants

// ServiceName The name of this module/service
const ServiceName = "users"

// GRPC Service Names
const (
	//StoresServiceName    = "STORES"
	UsersServiceName = "USERS"
)

// Dependency Injection Keys
const (
	RegistryKey                 = "registry"
	DomainDispatcherKey         = "domainDispatcher"
	DatabaseTransactionKey      = "tx"
	DatabaseKey                 = "dbConn"
	MessagePublisherKey         = "messagePublisher"
	MessageSubscriberKey        = "messageSubscriber"
	EventPublisherKey           = "eventPublisher"
	CommandPublisherKey         = "commandPublisher"
	ReplyPublisherKey           = "replyPublisher"
	SagaStoreKey                = "sagaStore"
	InboxStoreKey               = "inboxStore"
	ApplicationKey              = "application"
	DomainEventHandlersKey      = "domainEventHandlers"
	IntegrationEventHandlersKey = "integrationEventHandlers"
	CommandHandlersKey          = "commandHandlers"
	ReplyHandlersKey            = "replyHandlers"
	LoggerKey                   = "logger"
	TracerKey                   = "tracer"
	UsersRepoKey                = "usersRepo"
)

// Repository Table Names
const (
	//OutboxTableName    = ServiceName + ".outbox"
	//InboxTableName     = ServiceName + ".inbox"
	//EventsTableName    = ServiceName + ".events"
	//SnapshotsTableName = ServiceName + ".snapshots"
	//SagasTableName     = ServiceName + ".sagas"
	OutboxTableName    = "outbox"
	InboxTableName     = "inbox"
	EventsTableName    = "events"
	SnapshotsTableName = "snapshots"
	SagasTableName     = "sagas"

	UsersTableName = "users"
)

// Metric Names
const (
	UsersRegisteredCount = "users_registered_count"
)
