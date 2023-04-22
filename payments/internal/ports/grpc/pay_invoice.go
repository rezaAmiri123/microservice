package grpc

import (
	"context"
	"github.com/rezaAmiri123/microservice/payments/internal/app/commands"
	"github.com/rezaAmiri123/microservice/payments/paymentspb"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	grpcCodes "google.golang.org/grpc/codes"
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
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("InvoiceID", request.GetId()),
	)

	err := s.cfg.App.PayInvoice(ctx, commands.PayInvoice{
		ID: request.GetId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to pay an invoice: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, status.Errorf(grpcCodes.Internal, "failed to pay an invoice: %s", err)
	}
	resp := &paymentspb.PayInvoiceResponse{}

	return resp, nil
}
