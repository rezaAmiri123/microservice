package handlers

import (
	"context"
	"github.com/rezaAmiri123/microservice/cosec/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/sec"
)

func RegisterReplyHandlers(subscriber am.ReplySubscriber, orchestrator sec.Orchestrator[*domain.CreateOrderData]) error {
	replyMsgHandler := am.MessageHandlerFunc[am.IncomingReplyMessage](func(ctx context.Context, replyMsg am.IncomingReplyMessage) error {
		return orchestrator.HandleReply(ctx, replyMsg)
	})
	return subscriber.Subscribe(orchestrator.ReplyTopic(), replyMsgHandler, am.GroupName("cosec-replies"))
}
