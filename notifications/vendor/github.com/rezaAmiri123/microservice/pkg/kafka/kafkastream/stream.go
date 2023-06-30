package kafkastream

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"sync"
	"time"
)

const maxRetires = 5

type Stream struct {
	streamName string
	mu         sync.Mutex
	logger     logger.Logger
	// kafka tools
	brokers []string
	writer  *kafka.Writer
}

var _ am.MessageStream = (*Stream)(nil)

func NewStream(
	streamName string,
	logger logger.Logger,
	brokers []string,
) *Stream {
	return &Stream{
		streamName: streamName,
		logger:     logger,
		brokers:    brokers,
		writer:     NewWriter(brokers, kafka.LoggerFunc(logger.Errorf)),
	}
}
func (s *Stream) Publish(ctx context.Context, topicName string, rawMsg am.Message) (err error) {
	var data []byte

	metadata, err := structpb.NewStruct(rawMsg.Metadata())
	if err != nil {
		return err
	}

	data, err = proto.Marshal(&StreamMessage{
		Id:       rawMsg.ID(),
		Name:     rawMsg.MessageName(),
		Data:     rawMsg.Data(),
		Metadata: metadata,
		SentAt:   timestamppb.New(rawMsg.SentAt()),
	})

	if err != nil {
		return err
	}
	KafkaMsg := kafka.Message{
		Topic: topicName,
		Value: data,
		Time:  time.Now().UTC(),
		//Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	}
	// TODO implement retry to publish
	return s.writer.WriteMessages(ctx, KafkaMsg)
}

func (s *Stream) Subscribe(topicName string, handler am.MessageHandler, options ...am.SubscriberOption) (am.Subscription, error) {
	//var err error

	s.mu.Lock()
	defer s.mu.Unlock()

	subCfg := am.NewSubscriberConfig(options)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Topic:                  topicName,
		Brokers:                s.brokers,
		GroupID:                subCfg.GroupName(),
		MinBytes:               minBytes,
		MaxBytes:               maxBytes,
		QueueCapacity:          queueCapacity,
		HeartbeatInterval:      heartbeatInterval,
		CommitInterval:         commitInterval,
		PartitionWatchInterval: partitionWatchInterval,
		MaxAttempts:            subCfg.MaxRedeliver(),
		MaxWait:                maxWait,
		Dialer: &kafka.Dialer{
			Timeout: dialTimeout,
		},
		//GroupTopics: topicName,
	})
	//defer func() {
	//	if err := reader.Close(); err != nil {
	//		//c.log.Warn("consumerGroup.r.Close: %v", err)
	//	}
	//}()
	worker := s.handleMsg(subCfg, handler, reader)

	go func() {
		for {
			msg, err := reader.FetchMessage(context.Background())
			if err != nil {
				if err == io.EOF {
					return
				}
			}
			worker(msg)
		}
	}()
	return am.SubscriptionFunc(func() error {
		return reader.Close()
	}), nil

}
func (s *Stream) Unsubscribe() error {
	return nil
}
func (s *Stream) handleMsg(cfg am.SubscriberConfig, handler am.MessageHandler, reader *kafka.Reader) func(kafka.Message) {
	var filters map[string]struct{}
	if len(cfg.MessageFilters()) > 0 {
		filters = make(map[string]struct{})
		for _, key := range cfg.MessageFilters() {
			filters[key] = struct{}{}
		}
	}
	return func(kafkaMsg kafka.Message) {
		var err error

		m := &StreamMessage{}
		err = proto.Unmarshal(kafkaMsg.Value, m)
		if err != nil {
			//s.logger.Warn().Err(err).Msg("failed to unmarshal the *nats.Msg")
			return
		}

		if filters != nil {
			if _, exists := filters[m.GetName()]; !exists {
				//err = kafkaMsg.Ack()
				if err != nil {
					//s.logger.Warn().Err(err).Msg("failed to Ack a filtered message")
				}
				return
			}
		}
		wCtx, cancel := context.WithTimeout(context.Background(), cfg.AckWait())
		defer cancel()

		msg := &rawMessage{
			id:   m.GetId(),
			name: m.GetName(),
			//subject:    kafkaMsg.Subject,
			data:       m.GetData(),
			metadata:   m.GetMetadata().AsMap(),
			sentAt:     m.SentAt.AsTime(),
			receivedAt: time.Now(),
			acked:      false,
			ackFn:      func() error { return reader.CommitMessages(wCtx, kafkaMsg) },
			nackFn:     func() error { return reader.CommitMessages(wCtx, kafkaMsg) },
			//extendFn:   func() error { return natsMsg.InProgress() },
			//killFn:      func() error { return reader.Stats() },
		}
		errc := make(chan error)
		go func() {
			errc <- handler.HandleMessage(wCtx, msg)
		}()

		if cfg.AckType() == am.AckTypeAuto {
			err = msg.Ack()
			if err != nil {
				//s.logger.Warn().Err(err).Msg("failed to auto-Ack a message")
				s.logger.Warn("failed to auto-Ack a message")
			}
		}

		select {
		case err = <-errc:
			if err == nil {
				if ackErr := msg.Ack(); ackErr != nil {
					//s.logger.Warn().Err(err).Msg("failed to Ack a message")
					s.logger.Warn("failed to Ack a message")
				}
				return
			}
			//s.logger.Error().Err(err).Msg("error while handling message")
			if nakErr := msg.NAck(); nakErr != nil {
				//s.logger.Warn().Err(err).Msg("failed to Nack a message")
			}
		case <-wCtx.Done():
			// TODO logging?
			return
		}

	}
}

func (s *Stream) Close() error {
	return s.writer.Close()
}
