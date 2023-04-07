package grpc

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/depot/depotpb"
	"github.com/rezaAmiri123/microservice/depot/internal/app/commands"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) AssignShoppingList(ctx context.Context, request *depotpb.AssignShoppingListRequest) (resp *depotpb.AssignShoppingListResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.AssignShoppingList(ctx, request)
}

func (s server) AssignShoppingList(ctx context.Context, request *depotpb.AssignShoppingListRequest) (*depotpb.AssignShoppingListResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.AssignShoppingList")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()

	err := s.cfg.App.Commands.AssignShoppingList.Handle(ctx, commands.AssignShoppingList{
		ID:    request.GetId(),
		BotID: request.GetBotId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to assign shopping list: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to assign shopping list: %s", err)
	}

	resp := &depotpb.AssignShoppingListResponse{}

	return resp, nil
}
