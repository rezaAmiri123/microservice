package handlers

import (
	"github.com/rezaAmiri123/microservice/cosec/internal"
	"github.com/rezaAmiri123/microservice/cosec/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/registry"
	"github.com/rezaAmiri123/microservice/pkg/sec"
)

func NewReplyHandlers(reg registry.Registry, orchestrator sec.Orchestrator[*domain.CreateOrderData], mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewReplyHandler(reg, orchestrator, mws...)
}

func RegisterReplyHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) error {
	_, err := subscriber.Subscribe(internal.CreateOrderReplyChannel, handlers, am.GroupName("cosec-replies"))
	return err
}
