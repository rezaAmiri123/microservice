package grpc

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/depot/depotpb"
	"github.com/rezaAmiri123/microservice/depot/internal/app/commands"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) CompleteShoppingList(ctx context.Context, request *depotpb.CompleteShoppingListRequest) (resp *depotpb.CompleteShoppingListResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.CompleteShoppingList(ctx, request)
}

func (s server) CompleteShoppingList(ctx context.Context, request *depotpb.CompleteShoppingListRequest) (*depotpb.CompleteShoppingListResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.CompleteShoppingList")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()

	err := s.cfg.App.Commands.CompleteShoppingList.Handle(ctx, commands.CompleteShoppingList{
		ID: request.GetId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to complete shopping list: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to complete shopping list: %s", err)
	}

	resp := &depotpb.CompleteShoppingListResponse{}

	return resp, nil
}
