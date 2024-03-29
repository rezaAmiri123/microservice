package constants

// ServiceName The name of this module/service
const ServiceName = "stores"

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
	DatabaseKey                 = "dbConn"
	MessagePublisherKey         = "messagePublisher"
	MessageSubscriberKey        = "messageSubscriber"
	EventPublisherKey           = "eventPublisher"
	CommandPublisherKey         = "commandPublisher"
	ReplyPublisherKey           = "replyPublisher"
	AggregateStoreKey           = "aggregateStore"
	SagaStoreKey                = "sagaStore"
	InboxStoreKey               = "inboxStore"
	LoggerKey                   = "logger"
	ApplicationKey              = "app"
	DomainEventHandlersKey      = "domainEventHandlers"
	IntegrationEventHandlersKey = "integrationEventHandlers"
	CommandHandlersKey          = "commandHandlers"
	ReplyHandlersKey            = "replyHandlers"
	HttpServerKey               = "webServer"
	GrpcServerKey               = "grpcServer"
	TracerKey                   = "tracer"
	StreamKey                   = "sream"

	CatalogHandlersKey = "catalogHandlers"
	MallHandlersKey    = "mallHandlers"

	StoresRepoKey   = "storesRepo"
	ProductsRepoKey = "productsRepo"
	CatalogRepoKey  = "catalogRepo"
	MallRepoKey     = "mallRepo"
)

// Repository Table Names
//const (
//	OutboxTableName    = "outbox"
//	InboxTableName     = "inbox"
//	EventsTableName    = "events"
//	SnapshotsTableName = "snapshots"
//	SagasTableName     = "sagas"
//
//	CatalogTableName = "products"
//	MallTableName    = "stores"
//)

// Repository Table Names
const (
	OutboxTableName    = ServiceName + ".outbox"
	InboxTableName     = ServiceName + ".inbox"
	EventsTableName    = ServiceName + ".events"
	SnapshotsTableName = ServiceName + ".snapshots"
	SagasTableName     = ServiceName + ".sagas"

	CatalogTableName = ServiceName + ".products"
	MallTableName    = ServiceName + ".stores"
)
