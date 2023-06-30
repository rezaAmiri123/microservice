//go:generate mockgen -source interface.go -destination mock/interface.go -package mock
package kafka

import (
	"context"
	"sync"

	"github.com/segmentio/kafka-go"
)

// MessageProcessor processor methods must implement kafka.Worker func method interface
type MessageProcessor interface {
	ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int)
}

type Producer interface {
	PublishMessage(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}

type ConsumerGroup interface {
	ConsumeTopic(ctx context.Context, cancel context.CancelFunc, groupID, topic string, poolSize int, worker Worker)
	GetNewKafkaReader(kafkaURL []string, topic, groupID string) *kafka.Reader
	GetNewKafkaWriter(topic string) *kafka.Writer
}
