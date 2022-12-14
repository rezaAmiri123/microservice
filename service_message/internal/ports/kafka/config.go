package kafka

import (
	kafkaClient "github.com/rezaAmiri123/microservice/pkg/kafka"
)

type Config struct {
	KafkaTopics KafkaTopics
	Kafka       kafkaClient.Config
}
type KafkaTopics struct {
	UserCreate kafkaClient.TopicConfig
}
