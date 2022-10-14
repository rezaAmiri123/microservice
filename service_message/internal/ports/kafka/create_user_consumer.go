package kafka

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rezaAmiri123/microservice/pkg/tracing"
	"github.com/rezaAmiri123/microservice/service_message/internal/domain/message"
	kafkaMessages "github.com/rezaAmiri123/microservice/service_message/proto/kafka"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
)

const (
	retryAttempts = 3
	retryDelay    = 300 * time.Millisecond
)

//var (
//	retryOptions = []retry.Option{retry.Attempts(retryAttempts), retry.Delay(retryDelay), retry.DelayType(retry.BackOffDelay)}
//)

func (s *messageProcessor) processCreateUser(ctx context.Context, r *kafka.Reader, m kafka.Message) {
	s.metric.CreateEmailKafkaRequests.Inc()
	ctx, span := tracing.StartKafkaConsumerTracerSpan(ctx, m.Headers, "productMessageProcessor.processCreateUser")
	defer span.Finish()

	var msg kafkaMessages.CreateUser
	if err := proto.Unmarshal(m.Value, &msg); err != nil {
		s.log.WarnMsg("proto.Unmarshal", err)
		s.commitErrMessage(ctx, r, m)
		return
	}

	userId, err := uuid.Parse(msg.GetUserID())
	if err != nil {
		s.log.Errorf("no user id", err)
		s.commitErrMessage(ctx, r, m)
		return
	}
	req := &message.CreateEmailParams{
		UserID:  userId,
		From:    "example@example.com",
		To:      []string{msg.GetEmail()},
		Subject: "create user confirmation",
		Body:    "email body",
	}
	// send email

	_, err = s.app.Commands.CreateEmail.Handle(ctx, req)
	if err != nil {
		s.log.Errorf("error create email consumer", err)
	}
	s.commitMessage(ctx, r, m)
}
