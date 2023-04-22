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

func (s serverTx) ConfirmPayment(ctx context.Context, request *paymentspb.ConfirmPaymentRequest) (resp *paymentspb.ConfirmPaymentResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.ConfirmPayment(ctx, request)
}

func (s server) ConfirmPayment(ctx context.Context, request *paymentspb.ConfirmPaymentRequest) (*paymentspb.ConfirmPaymentResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("PaymentID", request.GetId()),
	)

	err := s.cfg.App.ConfirmPayment(ctx, commands.ConfirmPayment{
		ID: request.GetId(),
	})
	if err != nil {
		s.cfg.Logger.Errorf("failed to confirm a payment: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, status.Errorf(grpcCodes.Internal, "failed to confirm a payment: %s", err)
	}
	resp := &paymentspb.ConfirmPaymentResponse{}

	return resp, nil
}
