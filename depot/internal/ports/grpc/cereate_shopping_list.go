package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/depot/depotpb"
	"github.com/rezaAmiri123/microservice/depot/internal/app/commands"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) CreateShoppingList(ctx context.Context, request *depotpb.CreateShoppingListRequest) (resp *depotpb.CreateShoppingListResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.CreateShoppingList(ctx, request)
}

func (s server) CreateShoppingList(ctx context.Context, request *depotpb.CreateShoppingListRequest) (*depotpb.CreateShoppingListResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.CreateShoppingList")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()
	id := uuid.New().String()

	items := make([]commands.OrderItem, 0, len(request.GetItems()))
	for _, item := range request.GetItems() {
		items = append(items, s.itemToDomain(item))
	}

	err := s.cfg.App.Commands.CreateShoppingList.Handle(ctx, commands.CreateShoppingList{
		ID:      id,
		OrderID: request.GetOrderId(),
		Items:   items,
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to create shopping list: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to create shopping list: %s", err)
	}

	resp := &depotpb.CreateShoppingListResponse{Id: id}

	return resp, nil
}
