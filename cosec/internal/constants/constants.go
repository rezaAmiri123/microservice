package constants

// ServiceName The name of this module/service
const ServiceName = "cosec"

// GRPC Service Names
const (
	StoresServiceName = "STORES"
	UsersServiceName  = "USERS"
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
	StreamKey                   = "stream"

	SagaKey         = "saga"
	OrchestratorKey = "orchestrator"
)

// Repository Table Names
//const (
//	OutboxTableName    = "outbox"
//	InboxTableName     = "inbox"
//	EventsTableName    = "events"
//	SnapshotsTableName = "snapshots"
//	SagasTableName     = "sagas"
//)

// Repository Table Names
const (
	OutboxTableName    = ServiceName + ".outbox"
	InboxTableName     = ServiceName + ".inbox"
	EventsTableName    = ServiceName + ".events"
	SnapshotsTableName = ServiceName + ".snapshots"
	SagasTableName     = ServiceName + ".sagas"
)
