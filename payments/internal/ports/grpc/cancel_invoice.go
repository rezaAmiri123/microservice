package grpc

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/payments/internal/app/commands"
	"github.com/rezaAmiri123/microservice/payments/paymentspb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) CancelInvoice(ctx context.Context, request *paymentspb.CancelInvoiceRequest) (resp *paymentspb.CancelInvoiceResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.CancelInvoice(ctx, request)
}

func (s server) CancelInvoice(ctx context.Context, request *paymentspb.CancelInvoiceRequest) (*paymentspb.CancelInvoiceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.CancelInvoice")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()

	err := s.cfg.App.Commands.CancelInvoice.Handle(ctx, commands.CancelInvoice{
		ID: request.GetId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to cancel an invoice: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to cancel an invoice: %s", err)
	}
	resp := &paymentspb.CancelInvoiceResponse{}

	return resp, nil
}
