package agent

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/rezaAmiri123/microservice/depot/internal/constants"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/jetstream"
	"github.com/rezaAmiri123/microservice/pkg/kafka/kafkastream"
	"github.com/rezaAmiri123/microservice/pkg/logger"
)

func (a *Agent) setupEventServer() error {
	var stream am.MessageStream
	switch a.EventServerType {
	case "kafka":
		stream = kafkastream.NewStream(a.NatsStream, a.container.Get(constants.LoggerKey).(logger.Logger), a.KafkaBrokers)
	case "nats":
		js, err := a.nats()
		if err != nil {
			return err
		}
		stream = jetstream.NewStream(a.NatsStream, js, a.container.Get(constants.LoggerKey).(logger.Logger))
	default:
		return fmt.Errorf("event server typeis unknown")
	}

	a.container.AddSingleton(constants.StreamKey, func(c di.Container) (any, error) {
		return stream, nil
	})
	return nil
}
func (a *Agent) nats() (nats.JetStreamContext, error) {
	nc, err := nats.Connect(a.NatsURL)
	if err != nil {
		return nil, err
	}
	// defer nc.Close()
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     a.NatsStream,
		Subjects: []string{fmt.Sprintf("%s.>", a.NatsStream)},
	})

	return js, err
}
