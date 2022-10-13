package agent

import (
	"context"

	kafkaClient "github.com/rezaAmiri123/microservice/pkg/kafka"
	"github.com/rezaAmiri123/microservice/service_message/internal/ports/kafka"
)

func (a *Agent) setupKafka() error {
	ctx, cancel := context.WithCancel(context.Background())
	a.closers = append(a.closers, closer{cancel: cancel})
	kafkaCfg := kafka.Config{
		Kafka: kafkaClient.Config{
			Brokers:    a.KafkaBrokers,
			GroupID:    a.KafkaGroupID,
			InitTopics: a.KafkaInitTopics,
		},
		KafkaTopics: kafka.KafkaTopics{
			EmailCreate: kafkaClient.TopicConfig{
				TopicName: kafkaClient.CreateEmailTopic,
			},
		},
	}
	messageMessageProcessor := kafka.NewMessageProcessor(a.logger, kafkaCfg, a.metric, a.Application)
	cg := kafkaClient.NewConsumerGroup(kafkaCfg.Kafka.Brokers, kafkaCfg.Kafka.GroupID, a.logger)
	//kafkaConn, err := kafkaClient.NewKafkaConn(ctx, &a.KafkaConfig.Kafka)
	//if err != nil {
	//	return errors.Wrap(err, "kafka.NewKafkaCon")
	//}
	//
	topics := []string{
		kafkaCfg.KafkaTopics.EmailCreate.TopicName,
	}
	// TODO we need a context to can ended goroutine
	go cg.ConsumeTopic(ctx, topics, kafka.PoolSize, messageMessageProcessor.ProcessMessage)

	return nil
}

type closer struct {
	cancel func()
}

func (c closer) Close() error {
	c.cancel()
	return nil
}
