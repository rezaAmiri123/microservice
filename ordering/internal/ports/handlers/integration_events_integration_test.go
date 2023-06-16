//go:build integration

package handlers

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/rezaAmiri123/microservice/baskets/basketspb"
	"github.com/rezaAmiri123/microservice/depot/depotpb"
	"github.com/rezaAmiri123/microservice/ordering/internal/app"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/ddd"
	"github.com/rezaAmiri123/microservice/pkg/jetstream"
	"github.com/rezaAmiri123/microservice/pkg/logger/applogger"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
	"time"
)

const streamName = "mallbots"

type integrationEventsTestsuite struct {
	container  testcontainers.Container
	natsConn   *nats.Conn
	reg        registry.Registry
	js         nats.JetStreamContext
	publisher  am.EventPublisher
	subscriber am.MessageSubscriber
	mocks      struct {
		app *app.MockApp
	}
	suite.Suite
}

func TestIntegrationEvents(t *testing.T) {
	suite.Run(t, &integrationEventsTestsuite{})
}

func (s *integrationEventsTestsuite) SetupSuite() {
	var err error
	ctx := context.Background()
	natsReq := testcontainers.ContainerRequest{
		Image:        "nats:2-alpine",
		Hostname:     "nats",
		ExposedPorts: []string{"4222/tcp"},
		WaitingFor: wait.ForAll(
			wait.ForListeningPort("4222/tcp"),
			wait.ForLog("Server is ready"),
		),
		Cmd: []string{"-js"},
	}
	s.container, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: natsReq,
		Started:          true,
	})
	if err != nil {
		s.T().Fatal(err)
	}
	s.reg = registry.New()
	if err = basketspb.Registration(s.reg); err != nil {
		s.T().Fatal(err)
	}
	if err = depotpb.Registrations(s.reg); err != nil {
		s.T().Fatal(err)
	}
	endpoint, err := s.container.Endpoint(ctx, "")
	if err != nil {
		s.T().Fatal(err)
	}
	s.natsConn, err = nats.Connect(
		endpoint,
		nats.Timeout(5*time.Second),
		nats.RetryOnFailedConnect(true),
	)
	if err != nil {
		s.T().Fatal(err)
	}
	s.js, err = s.natsConn.JetStream()
	if err != nil {
		s.T().Fatal(err)
	}

	_, err = s.js.AddStream(&nats.StreamConfig{
		Name:     streamName,
		Subjects: []string{fmt.Sprintf("%s.>", streamName)},
	})
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *integrationEventsTestsuite) TearDownSuite() {
	s.natsConn.Close()
	if err := s.container.Terminate(context.Background()); err != nil {
		s.T().Fatal(err)
	}
}

func (s *integrationEventsTestsuite) SetupTest() {
	s.mocks = struct {
		app *app.MockApp
	}{
		app: app.NewMockApp(s.T()),
	}
	log := applogger.NewAppLogger(applogger.Config{})

	stream := jetstream.NewStream(streamName, s.js, log)
	s.publisher = am.NewEventPublisher(s.reg, stream)
	s.subscriber = stream

	//application := app.New(s.mocks.orders, s.publisher, log)
	handler := am.NewEventHandler(s.reg, integrationHandlers[ddd.Event]{
		app: s.mocks.app,
	})
	if err := RegisterIntegrationEventHandlers(s.subscriber, handler); err != nil {
		s.T().Fatal(err)
	}

}

func (s *integrationEventsTestsuite) TearDownTest() {
	if err := s.subscriber.Unsubscribe(); err != nil {
		s.T().Fatal(err)
	}
}

func (s *integrationEventsTestsuite) wait(aFn func(done chan struct{})) {
	done := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	aFn(done)
	select {
	case <-done:
		// time.Sleep(1 * time.Second)
	case <-ctx.Done():
		s.T().Error("test timed out")
	}
}

func (s *integrationEventsTestsuite) TestBasketAggregateChannel_BasketCheckedOut() {
	s.wait(func(done chan struct{}) {
		s.mocks.app.On("CreateOrder", mock.Anything, mock.AnythingOfType("commands.CreateOrder")).Return(nil).Run(func(_ mock.Arguments) {
			close(done)
		})
		_ = s.publisher.Publish(context.Background(), basketspb.BasketAggregateChannel,
			ddd.NewEvent(basketspb.BasketCheckedOutEvent, &basketspb.BasketCheckedOut{
				Id:        "basket-id",
				UserId:    "user-id",
				PaymentId: "payment-id",
				Items: []*basketspb.BasketCheckedOut_Item{{
					StoreId:     "store-id",
					StoreName:   "store-name",
					ProductId:   "product-id",
					ProductName: "product-name",
					Price:       10,
					Quantity:    2,
				}},
			}),
		)

	})

}

func (s *integrationEventsTestsuite) TestShoppingListAggregateChannel_ShoppingListCompleted() {
	s.wait(func(done chan struct{}) {
		s.mocks.app.On("ReadyOrder", mock.Anything, mock.AnythingOfType("commands.ReadyOrder")).Return(nil).Run(func(_ mock.Arguments) {
			close(done)
		})
		if err := s.publisher.Publish(context.Background(), depotpb.ShoppingListAggregateChannel,
			ddd.NewEvent(depotpb.ShoppingListCompletedEvent, &depotpb.ShoppingListCompleted{
				OrderId: "order-id",
				Id:      "basket-id",
			}),
		); err != nil {
			s.T().Fatal(err)
		}
	})
}

//
//func (s *integrationEventsTestsuite) {
//
//}
//
//func (s *integrationEventsTestsuite) {
//
//}
//
//func (s *integrationEventsTestsuite) {
//
//}
