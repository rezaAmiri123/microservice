package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/rezaAmiri123/microservice/baskets/basketspb"
	"github.com/rezaAmiri123/microservice/baskets/internal/app/commands"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) AddItem(ctx context.Context, request *basketspb.AddItemRequest) (resp *basketspb.AddItemResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.AddItem(ctx, request)
}

func (s server) AddItem(ctx context.Context, request *basketspb.AddItemRequest) (*basketspb.AddItemResponse, error) {
	//span, ctx := opentracing.StartSpanFromContext(ctx, "server.AddItem")
	//defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()
	id := uuid.New().String()
	err := s.cfg.App.Commands.AddItem.Handle(ctx, commands.AddItem{
		ID:        id,
		BasketID:  request.GetId(),
		ProductID: request.GetProductId(),
		Quantity:  int(request.GetQuantity()),
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to add item: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to add item: %s", err)
	}

	resp := &basketspb.AddItemResponse{}
	return resp, nil
}
