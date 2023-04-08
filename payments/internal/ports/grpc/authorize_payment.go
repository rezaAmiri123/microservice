package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/payments/internal/app/commands"
	"github.com/rezaAmiri123/microservice/payments/paymentspb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) AuthorizePayment(ctx context.Context, request *paymentspb.AuthorizePaymentRequest) (resp *paymentspb.AuthorizePaymentResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.AuthorizePayment(ctx, request)
}

func (s server) AuthorizePayment(ctx context.Context, request *paymentspb.AuthorizePaymentRequest) (*paymentspb.AuthorizePaymentResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.AuthorizePayment")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()
	id := uuid.New().String()
	err := s.cfg.App.Commands.AuthorizePayment.Handle(ctx, commands.AuthorizePayment{
		ID:     id,
		UserID: request.GetUserId(),
		Amount: request.GetAmount(),
	})
	if err != nil {
		s.cfg.Logger.Errorf("failed to authorize a payment: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to authorize a payment: %s", err)
	}
	resp := &paymentspb.AuthorizePaymentResponse{Id: id}

	return resp, nil
}
