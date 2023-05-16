package agent

import (
	"context"
	"github.com/rezaAmiri123/microservice/baskets/internal/constants"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"net/http"

	//"github.com/rezaAmiri123/microservice/users/internal/app"
	"io"
	"sync"
	"time"
)

type Config struct {
	// token
	SecretKey            string        `mapstructure:"SECRET_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`

	GRPCServerAddr string `mapstructure:"GRPC_SERVER_ADDR"`
	GRPCServerPort int    `mapstructure:"GRPC_SERVER_PORT"`

	GRPCServerTLSEnabled        bool   `mapstructure:"GRPC_SERVER_TLS_ENABLED"`
	GRPCServerTLSClientAuthType string `mapstructure:"GRPC_SERVER_TLS_CLIENT_AUTH_TYPE"`
	GRPCServerTLSCertFile       string `mapstructure:"GRPC_SERVER_TLS_CERT_FILE"`
	GRPCServerTLSKeyFile        string `mapstructure:"GRPC_SERVER_TLS_KEY_FILE"`
	GRPCServerTLSCAFile         string `mapstructure:"GRPC_SERVER_TLS_CA_FILE"`
	GRPCServerTLSServerAddress  string `mapstructure:"GRPC_SERVER_TLS_SERVER_ADDRESS"`

	//grpc clients
	GRPCStoreClientAddr string `mapstructure:"GRPC_STORE_CLIENT_ADDR"`
	GRPCStoreClientPort int    `mapstructure:"GRPC_STORE_CLIENT_PORT"`

	GRPCStoreClientTLSEnabled       bool   `mapstructure:"GRPC_STORE_CLIENT_TLS_ENABLED"`
	GRPCStoreClientTLSCertFile      string `mapstructure:"GRPC_STORE_CLIENT_TLS_CERT_FILE"`
	GRPCStoreClientTLSKeyFile       string `mapstructure:"GRPC_STORE_CLIENT_TLS_KEY_FILE"`
	GRPCStoreClientTLSCAFile        string `mapstructure:"GRPC_STORE_CLIENT_TLS_CA_FILE"`
	GRPCStoreClientTLSServerAddress string `mapstructure:"GRPC_STORE_CLIENT_TLS_SERVER_ADDRESS"`

	// Http Server
	HttpServerAddr string `mapstructure:"HTTP_SERVER_ADDR"`
	HttpServerPort int    `mapstructure:"HTTP_SERVER_PORT"`

	//DBConfig     adapters.GORMConfig
	// postgres.Config
	PGDriver     string `mapstructure:"POSTGRES_DRIVER"`
	PGHost       string `mapstructure:"POSTGRES_HOST"`
	PGPort       string `mapstructure:"POSTGRES_PORT"`
	PGUser       string `mapstructure:"POSTGRES_USER"`
	PGDBName     string `mapstructure:"POSTGRES_DB_NAME"`
	PGPassword   string `mapstructure:"POSTGRES_PASSWORD"`
	PGSearchPath string `mapstructure:"POSTGRES_SEARCH_PATH"`
	//DBConn       string `mapstructure:"DB_CONN"`

	// Event Server
	EventServerType string `mapstructure:"EVENT_SERVER_TYPE"`
	// kafka config
	KafkaBrokers []string `mapstructure:"KAFKA_BROKERS"`

	NatsURL    string `mapstructure:"NATS_URL"`
	NatsStream string `mapstructure:"NATS_STREAM"`
	// applogger.Config
	LogLevel   string `mapstructure:"LOG_LEVEL"`
	LogDevMode bool   `mapstructure:"LOG_DEV_MOD"`
	LogEncoder string `mapstructure:"LOG_ENCODER"`

	// tracing.Config
	TracerServiceName string `mapstructure:"TRACER_SERVICE_NAME"`
	TracerHostPort    string `mapstructure:"TRACER_HOST_PORT"`
	TracerEnable      bool   `mapstructure:"TRACER_ENABLE"`
	TracerLogSpans    bool   `mapstructure:"TRACER_LOG_SPANS"`

	// trace config
	ExporterEndpoint string `mapstructure:"OTEL_EXPORTER_OTLP_ENDPOINT"  default:"http://collector:4317"`
	ServiceName      string `mapstructure:"OTEL_SERVICE_NAME"  default:"mallbots"`
	// metrics.Config
	MetricServiceName     string `mapstructure:"METRIC_SERVICE_NAME"`
	MetricServiceHostPort string `mapstructure:"METRIC_SERVICE_HOST_PORT"`
}

type Agent struct {
	Config

	//GrpcServerTLSConfig *tls.Config

	container di.Container
	//logger    logger.Logger
	//metric *metrics.UserServiceMetric
	//httpServer *http.Server
	//grpcServer *grpc.Server
	//repository  user.Repository
	//Application *app.Application
	//Maker       token.Maker
	// AuthClient  auth.AuthClient

	shutdown     bool
	shutdowns    chan struct{}
	shutdownLock sync.Mutex
	closers      []io.Closer
}

//type CloserFunc func() error
//
//func (f CloserFunc) Close() error {
//	return f()
//}

func NewAgent(config Config) (*Agent, error) {
	a := &Agent{
		Config:    config,
		container: di.New(),
		shutdowns: make(chan struct{}),
	}
	setupsFn := []func() error{
		a.setupLogger,
		a.setupRegistry,

		//a.setupRepository,
		a.setupTracer,
		a.setupEventServer,
		a.setupApplication,
		//a.setupAuthClient,
		a.setupGrpcServer,
		a.setupHttpServer,
		//a.setupGRPCServer,
		//a.setupTracer,
	}
	for _, fn := range setupsFn {
		if err := fn(); err != nil {
			return nil, err
		}
	}
	return a, nil
}

func (a *Agent) Shutdown() error {
	a.shutdownLock.Lock()
	defer a.shutdownLock.Unlock()

	if a.shutdown {
		return nil
	}
	a.shutdown = true
	close(a.shutdowns)
	shutdown := []func() error{

		func() error {
			grpcServer := a.container.Get(constants.GrpcServerKey).(*grpc.Server)
			grpcServer.GracefulStop()
			return nil
		},
		func() error {
			httpServer := a.container.Get(constants.HttpServerKey).(*http.Server)
			return httpServer.Shutdown(context.Background())
		},
		func() error {
			tp := a.container.Get(constants.TracerKey).(*trace.TracerProvider)
			return tp.Shutdown(context.Background())
		},
		func() error {
			stream := a.container.Get(constants.StreamKey).(am.MessageStream)
			return stream.Unsubscribe()
		},

		//func() error {
		//	return a.jaegerCloser.Close()
		//},
	}
	for _, fn := range shutdown {
		if err := fn(); err != nil {
			return err
		}
	}
	for _, closer := range a.closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}
