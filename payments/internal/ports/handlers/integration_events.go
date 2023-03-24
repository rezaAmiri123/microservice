package handlers

//
//type integrationHandlers[T ddd.Event] struct {
//	app app.Application
//}
//
//var _ ddd.EventHandler[ddd.Event] = (*integrationHandlers[ddd.Event])(nil)
//
//func NewIntegrationEventHandlers(reg registry.Registry, app app.Application, mws ...am.MessageHandlerMiddleware) am.MessageHandler {
//	return am.NewEventHandler(reg, integrationHandlers[ddd.Event]{
//		app: app,
//	}, mws...)
//}
//
//func RegisterIntegrationEventHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) error {
//	_, err := subscriber.Subscribe()
//}
