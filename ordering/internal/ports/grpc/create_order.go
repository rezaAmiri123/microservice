package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/ordering/internal/app/commands"
	"github.com/rezaAmiri123/microservice/ordering/internal/domain"
	"github.com/rezaAmiri123/microservice/ordering/orderingpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) CreateOrder(ctx context.Context, request *orderingpb.CreateOrderRequest) (resp *orderingpb.CreateOrderResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.CreateOrder(ctx, request)
}

func (s server) CreateOrder(ctx context.Context, request *orderingpb.CreateOrderRequest) (*orderingpb.CreateOrderResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.CreateOrder")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()
	id := uuid.New().String()

	items := make([]domain.Item, len(request.GetItems()))
	for i, item := range request.GetItems() {
		items[i] = s.itemToDomain(item)
	}

	err := s.cfg.App.Commands.CreateOrder.Handle(ctx, commands.CreateOrder{
		ID:        id,
		UserID:    request.GetUserId(),
		PaymentID: request.GetPaymentId(),
		Items:     items,
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to create order: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to create create order: %s", err)
	}

	resp := &orderingpb.CreateOrderResponse{Id: id}

	return resp, nil
}
