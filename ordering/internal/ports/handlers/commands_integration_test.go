///go:build integration

package handlers

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/rezaAmiri123/microservice/ordering/internal/app"
	"github.com/rezaAmiri123/microservice/ordering/internal/app/commands"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
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

const commandStreamName = "mallbots"

type commandIntegrationEventsTestsuite struct {
	container        testcontainers.Container
	natsConn         *nats.Conn
	reg              registry.Registry
	js               nats.JetStreamContext
	commandPublisher am.CommandPublisher
	replyPublisher   am.ReplyPublisher
	subscriber       am.MessageSubscriber
	mocks            struct {
		app *app.MockApp
	}
	suite.Suite
}

func TestCommandIntegrationEvents(t *testing.T) {
	suite.Run(t, &commandIntegrationEventsTestsuite{})
}

func (s *commandIntegrationEventsTestsuite) SetupSuite() {
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
	if err = orderingpb.Registrations(s.reg); err != nil {
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
		Name:     commandStreamName,
		Subjects: []string{fmt.Sprintf("%s.>", commandStreamName)},
	})
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *commandIntegrationEventsTestsuite) TearDownSuite() {
	s.natsConn.Close()
	if err := s.container.Terminate(context.Background()); err != nil {
		s.T().Fatal(err)
	}
}

func (s *commandIntegrationEventsTestsuite) SetupTest() {
	s.mocks = struct {
		app *app.MockApp
	}{
		app: app.NewMockApp(s.T()),
	}
	log := applogger.NewAppLogger(applogger.Config{})

	stream := jetstream.NewStream(commandStreamName, s.js, log)
	s.subscriber = stream

	messagePublisher := am.NewMessagePublisher(stream)
	s.commandPublisher = am.NewCommandPublisher(s.reg, messagePublisher)
	s.replyPublisher = am.NewReplyPublisher(s.reg, messagePublisher)

	commandHandler := am.NewCommandHandler(s.reg, s.replyPublisher, commandHandlers{
		app: s.mocks.app,
	})
	if err := RegisterCommandHandlers(s.subscriber, commandHandler); err != nil {
		s.T().Fatal(err)
	}

}

func (s *commandIntegrationEventsTestsuite) TearDownTest() {
	if err := s.subscriber.Unsubscribe(); err != nil {
		s.T().Fatal(err)
	}
}

func (s *commandIntegrationEventsTestsuite) wait(aFn func(done chan struct{})) {
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

func (s *commandIntegrationEventsTestsuite) TestCommandChannel_RejectOrder() {
	s.wait(func(done chan struct{}) {
		s.mocks.app.On("CancelOrder", mock.Anything, commands.CancelOrder{
			ID: "order-id"}).
			Return(nil).Run(func(_ mock.Arguments) {
			close(done)
		})
		cmd := ddd.NewCommand(orderingpb.RejectOrderCommand, &orderingpb.RejectOrder{Id: "order-id"})
		cmd.Metadata().Set(am.CommandReplyChannelHdr, "reply")

		err := s.commandPublisher.Publish(context.Background(), orderingpb.CommandChannel, cmd)
		if err != nil {
			s.T().Fatal(err)
		}
	})
}

func (s *commandIntegrationEventsTestsuite) TestCommandChannel_ApproveOrderCommand() {
	s.wait(func(done chan struct{}) {
		s.mocks.app.On("ApproveOrder", mock.Anything, commands.ApproveOrder{
			ID:         "order-id",
			ShoppingID: "shopping-id"}).
			Return(nil).Run(func(_ mock.Arguments) {
			close(done)
		})
		cmd := ddd.NewCommand(orderingpb.ApproveOrderCommand, &orderingpb.ApproveOrder{
			Id:         "order-id",
			ShoppingId: "shopping-id",
		})
		cmd.Metadata().Set(am.CommandReplyChannelHdr, "reply")

		err := s.commandPublisher.Publish(context.Background(), orderingpb.CommandChannel, cmd)
		if err != nil {
			s.T().Fatal(err)
		}
	})
}
