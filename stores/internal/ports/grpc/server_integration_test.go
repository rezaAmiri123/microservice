package grpc

import (
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/stores/internal/app"
	"github.com/rezaAmiri123/microservice/stores/storespb"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"
)

type serverSuite struct {
	mocks struct {
		//products  *domain.MockProductRepository
		//stores    *domain.MockStoreRepository
		//malls     *domain.MockMallRepository
		//catalogs  *domain.MockCatalogRepository
		//publisher *ddd.MockEventPublisher[ddd.Event]
		logger *logger.MockLogger
		app    *app.MockApp
	}
	server *grpc.Server
	client storespb.StoresServiceClient
	suite.Suite
}

func TestServer(t *testing.T) {
	suite.Run(t, &serverSuite{})
}

func (s *serverSuite) SetupSuite() {}

func (s *serverSuite) TearDownSuite() {}

func (s *serverSuite) SetupTest() {
	const grpcTestPort = ":10912"

	var err error
	//create server
	s.server = grpc.NewServer()
	var listener net.Listener
	listener, err = net.Listen("tcp", grpcTestPort)
	if err != nil {
		s.T().Fatal(err)
	}

	// create mocks
	s.mocks = struct {
		//products  *domain.MockProductRepository
		//stores    *domain.MockStoreRepository
		//malls     *domain.MockMallRepository
		//catalogs  *domain.MockCatalogRepository
		//publisher *ddd.MockEventPublisher[ddd.Event]
		logger *logger.MockLogger
		app    *app.MockApp
	}{
		//products:  domain.NewMockProductRepository(s.T()),
		//stores:    domain.NewMockStoreRepository(s.T()),
		//malls:     domain.NewMockMallRepository(s.T()),
		//catalogs:  domain.NewMockCatalogRepository(s.T()),
		//publisher: ddd.NewMockEventPublisher[ddd.Event](s.T()),
		logger: logger.NewMockLogger(s.T()),
		app:    app.NewMockApp(s.T()),
	}

	// create app
	//application := app.New(
	//	s.mocks.stores,
	//	s.mocks.products,
	//	s.mocks.catalogs,
	//	s.mocks.malls,
	//	s.mocks.publisher,
	//	s.mocks.logger,
	//)

	// register app with server
	if err = RegisterServer(s.mocks.app, s.server, s.mocks.logger); err != nil {
		s.T().Fatal(err)
	}

	go func(listener net.Listener) {
		err := s.server.Serve(listener)
		if err != nil {
			s.T().Fatal(err)
		}
	}(listener)

	// create client
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(grpcTestPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.T().Fatal(err)
	}
	s.client = storespb.NewStoresServiceClient(conn)

}
func (s *serverSuite) TearDownTest() {
	s.server.GracefulStop()
}
