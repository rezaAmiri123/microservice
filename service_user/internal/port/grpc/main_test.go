package grpc_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/pkg/logger/applogger"
	"github.com/rezaAmiri123/microservice/pkg/token/mock"
	"github.com/rezaAmiri123/microservice/service_user/internal/app"
	"github.com/rezaAmiri123/microservice/service_user/internal/app/command"
	repomock "github.com/rezaAmiri123/microservice/service_user/internal/domain/user/mock"
	"github.com/rezaAmiri123/microservice/service_user/internal/metrics"
	grpcervice "github.com/rezaAmiri123/microservice/service_user/internal/port/grpc"
	"github.com/stretchr/testify/require"
)

type TestGrpcServer struct {
	grpcServer *grpcervice.UserGRPCServer
	logger     logger.Logger
	repoMock   *repomock.MockRepository
	Metric     *metrics.UserServiceMetric
}

func NewTestGrpcServer(t *testing.T, ctrl *gomock.Controller) *TestGrpcServer {
	t.Helper()

	appLogger := getLogger()

	repoMock := repomock.NewMockRepository(ctrl)
	makerMock := maker_mock.NewMockMaker(ctrl)

	application := &app.Application{
		Commands: app.Commands{
			CreateUser: command.NewCreateUserHandler(repoMock, appLogger),
		},
	}

	metric := metrics.NewUserServiceMetric(&metrics.Config{
		MetricServiceName: fmt.Sprintf("rand%d", rand.Int()),
	})

	serverConfig := grpcervice.Config{App: application, Metric: metric, Logger: appLogger, Maker: makerMock}
	server, err := grpcervice.NewUserGRPCServer(&serverConfig)
	require.NoError(t, err)

	testServer := &TestGrpcServer{
		logger:     appLogger,
		Metric:     metric,
		repoMock:   repoMock,
		grpcServer: server,
	}
	return testServer
}

func getLogger() logger.Logger {
	appLogger := applogger.NewAppLogger(applogger.Config{})
	appLogger.InitLogger()
	appLogger.WithName("ServiceUserTest")
	return appLogger
}
