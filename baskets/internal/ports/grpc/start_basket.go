package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/baskets/basketspb"
	"github.com/rezaAmiri123/microservice/baskets/internal/app/commands"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) StartBasket(ctx context.Context, request *basketspb.StartBasketRequest) (resp *basketspb.StartBasketResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.StartBasket(ctx, request)
}

func (s server) StartBasket(ctx context.Context, request *basketspb.StartBasketRequest) (*basketspb.StartBasketResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.StartBasket")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()
	id := uuid.New().String()

	err := s.cfg.App.Commands.StartBasket.Handle(ctx, commands.StartBasket{
		ID:     id,
		UserID: request.GetUserId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to start basket: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to start basket: %s", err)
	}

	resp := &basketspb.StartBasketResponse{Id: id}

	return resp, nil
}
