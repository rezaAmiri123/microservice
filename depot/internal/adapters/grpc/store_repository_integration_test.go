package grpc

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/pkg/logger/applogger"
	"github.com/rezaAmiri123/microservice/stores/storespb"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"net"
	"testing"
)

const grpcTestPort = "localhost:10912"

type integrationStoreTestSuite struct {
	fakeServer *FakeStoreServer
	server     *grpc.Server
	logger     logger.Logger
	suite.Suite
}

func TestServer(t *testing.T) {
	suite.Run(t, &integrationStoreTestSuite{})
}

func (s *integrationStoreTestSuite) SetupSuite()    {}
func (s *integrationStoreTestSuite) TearDownSuite() {}

func (s *integrationStoreTestSuite) SetupTest() {
	var err error
	// create server
	s.server = grpc.NewServer()
	var listener net.Listener
	listener, err = net.Listen("tcp", grpcTestPort)
	if err != nil {
		s.T().Fatal(err)
	}
	s.fakeServer = NewFakeStoreServer()
	storespb.RegisterStoresServiceServer(s.server, s.fakeServer)

	log := applogger.NewAppLogger(applogger.Config{})
	log.InitLogger()
	s.logger = log
	go func(listener net.Listener) {
		err := s.server.Serve(listener)
		if err != nil {
			s.T().Fatal(err)
		}
	}(listener)
}

func (s *integrationStoreTestSuite) TearDownTest() {
	s.server.GracefulStop()
	s.fakeServer.Reset()
}

func (s *integrationStoreTestSuite) TestStoreRepository_Find() {
	store := &storespb.Store{
		Id:            "store-id",
		Name:          "store-name",
		Location:      "store-location",
		Participating: false,
	}

	s.fakeServer.AddStores(store)
	repo := NewStoreRepository(grpcTestPort, s.logger)
	gotStore, err := repo.Find(context.Background(), "store-id")
	if err != nil {
		s.T().Fatal(err)
	}
	s.Assert().Equal(gotStore.ID, store.GetId())
	s.Assert().Equal(gotStore.Name, store.GetName())
	s.Assert().Equal(gotStore.Location, store.GetLocation())
}
