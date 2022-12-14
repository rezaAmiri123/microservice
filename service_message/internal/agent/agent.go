package agent

import (
	"io"
	"sync"

	"github.com/rezaAmiri123/microservice/service_message/internal/app"
	"github.com/rezaAmiri123/microservice/service_message/internal/domain/message"
	"github.com/rezaAmiri123/microservice/service_message/internal/metrics"

	// "github.com/rezaAmiri123/test-microservice/pkg/auth"

	"github.com/rezaAmiri123/microservice/pkg/logger"
	// "github.com/rezaAmiri123/microservice/wallet_service/internal/app"
	// "github.com/rezaAmiri123/microservice/wallet_service/internal/domain/wallet"
	// "github.com/rezaAmiri123/microservice/wallet_service/internal/metrics"
)

type Config struct {
	// ServerTLSConfig *tls.Config
	// token
	SecretKey string `mapstructure:"SECRET_KEY"`

	// GRPCServerAddr string `mapstructure:"GRPC_SERVER_ADDR"`
	// GRPCServerPort int    `mapstructure:"GRPC_SERVER_PORT"`

	// postgres.Config
	PGDriver   string `mapstructure:"POSTGRES_DRIVER"`
	PGHost     string `mapstructure:"POSTGRES_HOST"`
	PGPort     string `mapstructure:"POSTGRES_PORT"`
	PGUser     string `mapstructure:"POSTGRES_USER"`
	PGDBName   string `mapstructure:"POSTGRES_DB_NAME"`
	PGPassword string `mapstructure:"POSTGRES_PASSWORD"`

	// kafka config
	KafkaBrokers    []string `mapstructure:"KAFKA_BROKERS"`
	KafkaGroupID    string   `mapstructure:"KAFKA_GROUP_ID"`
	KafkaInitTopics bool     `mapstructure:"KAFKA_INIT_TOPICS"`

	// applogger.Config
	LogLevel   string `mapstructure:"LOG_LEVEL"`
	LogDevMode bool   `mapstructure:"LOG_DEV_MOD"`
	LogEncoder string `mapstructure:"LOG_ENCODER"`

	// tracing.Config
	TracerServiceName string `mapstructure:"TRACER_SERVICE_NAME"`
	TracerHostPort    string `mapstructure:"TRACER_HOST_PORT"`
	TracerEnable      bool   `mapstructure:"TRACER_ENABLE"`
	TracerLogSpans    bool   `mapstructure:"TRACER_LOG_SPANS"`

	// metrics.Config
	MetricServiceName     string `mapstructure:"METRIC_SERVICE_NAME"`
	MetricServiceHostPort string `mapstructure:"METRIC_SERVICE_HOST_PORT"`
}

type Agent struct {
	Config

	logger logger.Logger
	metric *metrics.MessageServiceMetric
	// httpServer  *http.Server
	// grpcServer  *grpc.Server
	repository  message.Repository
	Application *app.Application
	// AuthClient  auth.AuthClient

	shutdown     bool
	shutdowns    chan struct{}
	shutdownLock sync.Mutex
	closers      []io.Closer
}

func NewAgent(config Config) (*Agent, error) {
	a := &Agent{
		Config:    config,
		shutdowns: make(chan struct{}),
	}
	setupsFn := []func() error{
		a.setupLogger,
		a.setupMetric,

		//a.setupRepository,
		a.setupTracing,
		a.setupApplication,
		a.setupKafka,
		//a.setupAuthClient,
		//a.setupHttpServer,
		// a.setupGrpcServer,
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
		//func() error {
		//	return a.httpServer.Shutdown(context.Background())
		//},
		// func() error {
		// 	a.grpcServer.GracefulStop()
		// 	return nil
		// },
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
