package kafka

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/service_message/internal/app"
	"github.com/rezaAmiri123/microservice/service_message/internal/metrics"
	"github.com/segmentio/kafka-go"
	"sync"
)

const PoolSize = 30

type messageProcessor struct {
	log    logger.Logger
	cfg    Config
	metric *metrics.MessageServiceMetric
	app    *app.Application
}

func NewMessageProcessor(log logger.Logger, cfg Config, metric *metrics.MessageServiceMetric, app *app.Application) *messageProcessor {
	return &messageProcessor{log: log, cfg: cfg, metric: metric, app: app}
}

func (s *messageProcessor) ProcessMessage(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		m, err := r.FetchMessage(ctx)
		if err != nil {
			s.log.Warnf("workerID: %v, err: %v", workerID, err)
			continue
		}
		s.logProcessMessage(m, workerID)
		switch m.Topic {
		case s.cfg.KafkaTopics.UserCreate.TopicName:
			s.processCreateUser(ctx, r, m)
		}
	}
}
