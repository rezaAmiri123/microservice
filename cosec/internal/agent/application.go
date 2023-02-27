package agent

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/rezaAmiri123/microservice/cosec/internal"
	"github.com/rezaAmiri123/microservice/cosec/internal/domain"
	"github.com/rezaAmiri123/microservice/cosec/internal/handlers"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/jetstream"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/registry/serdes"
	"github.com/rezaAmiri123/microservice/pkg/sec"
	"github.com/rezaAmiri123/microservice/users/userspb"
)

func (a *Agent) setupApplication() (err error) {
	// setup Driven adapters
	reg := registry.New()
	if err = registrations(reg); err != nil {
		return err
	}
	//if err = orderingpb.Registrations(reg); err != nil {
	//	return err
	//}
	if err = userspb.Registrations(reg); err != nil {
		return err
	}
	//if err = depotpb.Registrations(reg); err != nil {
	//	return err
	//}
	//if err = paymentspb.Registrations(reg); err != nil {
	//	return err
	//}
	js, _ := a.nats()
	stream := jetstream.NewStream("stream", js, a.logger)
	eventStream := am.NewEventStream(reg, stream)
	commandStream := am.NewCommandStream(reg, stream)
	replyStream := am.NewReplyStream(reg, stream)
	//sagaStore := pg.NewSagaStore("cosec.sagas", mono.DB(), reg)
	var sagaStore sec.SagaStore
	sagaRepo := sec.NewSagaRepository[*domain.CreateOrderData](reg, sagaStore)

	// setup application
	orchestrator := sec.NewOrchestrator[*domain.CreateOrderData](internal.NewCreateOrderSaga(), sagaRepo, commandStream)
	integrationEventHandlers := handlers.NewIntegrationEventHandlers(orchestrator)
	// setup Driver adapters
	if err = handlers.RegisterIntegrationEventHandlers(eventStream, integrationEventHandlers); err != nil {
		return err
	}
	if err = handlers.RegisterReplyHandlers(replyStream, orchestrator); err != nil {
		return err
	}
	return err
}
func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Saga data
	if err = serde.RegisterKey(internal.CreateOrderSagaName, domain.CreateOrderData{}); err != nil {
		return err
	}

	return nil
}

func (a *Agent) nats() (nats.JetStreamContext, error) {
	nc, err := nats.Connect("localhost")
	if err != nil {
		return nil, err
	}
	// defer nc.Close()
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "stream",
		Subjects: []string{fmt.Sprintf("%s.>", "stream")},
	})

	return js, err
}
