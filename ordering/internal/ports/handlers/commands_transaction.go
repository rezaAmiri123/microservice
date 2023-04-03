package handlers

import (
	"context"
	"github.com/rezaAmiri123/microservice/ordering/internal/constants"
	"github.com/rezaAmiri123/microservice/pkg/am"
	"github.com/rezaAmiri123/microservice/pkg/di"
)

func RegisterCommandHandlersTx(container di.Container) error {
	rawMsgHandler := am.MessageHandlerFunc(func(ctx context.Context, msg am.IncomingMessage) (err error) {
		ctx = container.Scoped(ctx)
		//defer func(tx *sql.Tx) {
		//	if p := recover(); p != nil {
		//		_ = tx.Rollback()
		//		panic(p)
		//	} else if err != nil {
		//		_ = tx.Rollback()
		//	} else {
		//		err = tx.Commit()
		//	}
		//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

		return di.Get(ctx, constants.CommandHandlersKey).(am.MessageHandler).HandleMessage(ctx, msg)
	})

	subscriber := container.Get(constants.MessageSubscriberKey).(am.MessageSubscriber)

	return RegisterCommandHandlers(subscriber, rawMsgHandler)
}
