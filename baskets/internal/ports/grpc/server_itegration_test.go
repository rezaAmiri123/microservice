///go:build integration

package grpc

import (
	"github.com/rezaAmiri123/microservice/baskets/basketspb"
	"github.com/rezaAmiri123/microservice/baskets/internal/app"
	"github.com/rezaAmiri123/microservice/baskets/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"
)

type serverSuite struct {
	mocks struct {
		baskets   *domain.MockBasketRepository
		stores    *domain.MockStoreRepository
		products  *domain.MockProductRepository
		publisher *ddd.MockEventPublisher[ddd.Event]
		logger    *logger.MockLogger
	}
	server *grpc.Server
	client basketspb.BasketServiceClient
	suite.Suite
}

func TestServer(t *testing.T) {
	suite.Run(t, &serverSuite{})
}

func (s *serverSuite) SetupSuite()    {}
func (s *serverSuite) TearDownSuite() {}

func (s *serverSuite) SetupTest() {
	const grpcTestPort = ":10912"

	var err error
	// create server
	s.server = grpc.NewServer()
	var listener net.Listener
	listener, err = net.Listen("tcp", grpcTestPort)
	if err != nil {
		s.T().Fatal(err)
	}

	// create mocks
	s.mocks = struct {
		baskets   *domain.MockBasketRepository
		stores    *domain.MockStoreRepository
		products  *domain.MockProductRepository
		publisher *ddd.MockEventPublisher[ddd.Event]
		logger    *logger.MockLogger
	}{
		baskets:   domain.NewMockBasketRepository(s.T()),
		stores:    domain.NewMockStoreRepository(s.T()),
		products:  domain.NewMockProductRepository(s.T()),
		publisher: ddd.NewMockEventPublisher[ddd.Event](s.T()),
		logger:    logger.NewMockLogger(s.T()),
	}

	// create app
	application := app.New(s.mocks.baskets, s.mocks.stores, s.mocks.products, s.mocks.publisher, s.mocks.logger)
	// register app with server
	if err = RegisterServer(application, s.server, s.mocks.logger); err != nil {
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

	s.client = basketspb.NewBasketServiceClient(conn)
}

func (s *serverSuite) TearDownTest() {
	s.server.GracefulStop()
}
