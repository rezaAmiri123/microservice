package grpc

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/payments/internal/app/commands"
	"github.com/rezaAmiri123/microservice/payments/paymentspb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) PayInvoice(ctx context.Context, request *paymentspb.PayInvoiceRequest) (resp *paymentspb.PayInvoiceResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.PayInvoice(ctx, request)
}

func (s server) PayInvoice(ctx context.Context, request *paymentspb.PayInvoiceRequest) (*paymentspb.PayInvoiceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.PayInvoice")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()

	err := s.cfg.App.Commands.PayInvoice.Handle(ctx, commands.PayInvoice{
		ID: request.GetId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to pay an invoice: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to pay an invoice: %s", err)
	}
	resp := &paymentspb.PayInvoiceResponse{}

	return resp, nil
}
