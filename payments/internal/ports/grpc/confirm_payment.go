package grpc

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/payments/internal/app/commands"
	"github.com/rezaAmiri123/microservice/payments/paymentspb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) ConfirmPayment(ctx context.Context, request *paymentspb.ConfirmPaymentRequest) (resp *paymentspb.ConfirmPaymentResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.ConfirmPayment(ctx, request)
}

func (s server) ConfirmPayment(ctx context.Context, request *paymentspb.ConfirmPaymentRequest) (*paymentspb.ConfirmPaymentResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.ConfirmPayment")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()
	err := s.cfg.App.Commands.ConfirmPayment.Handle(ctx, commands.ConfirmPayment{
		ID: request.GetId(),
	})
	if err != nil {
		s.cfg.Logger.Errorf("failed to confirm a payment: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to confirm a payment: %s", err)
	}
	resp := &paymentspb.ConfirmPaymentResponse{}

	return resp, nil
}
