package kafka

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/tracing"
	"github.com/rezaAmiri123/microservice/service_message/internal/domain/message"
	kafkaMessages "github.com/rezaAmiri123/microservice/service_message/proto/kafka"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"time"
)

const (
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

//var (
//	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
//)

func (s *messageProcessor) processCreateEmail(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metric.CreateEmailKafkaRequests.Inc()
	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "productMessageProcessor.processCreateProduct")
	defer span.Finish()

	var msg kafkaMessages.CreateEmailRequest
	if err := proto.Unmarshal(m.Value, &msg); err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	protoEmail := msg.GetEmail()
	req := &message.CreateEmailParams{
		From:    protoEmail.GetFrom(),
		To:      protoEmail.GetTo(),
		Subject: protoEmail.GetSubject(),
		Body:    protoEmail.GetBody(),
	}
	// send email

	_, err := s.app.Commands.CreateEmail.Handle(ctx, req)
	if err != nil {
		s.log.Errorf("error create email consumer", err)
	}
	s.commitMessage(ctx, r, m)
}
