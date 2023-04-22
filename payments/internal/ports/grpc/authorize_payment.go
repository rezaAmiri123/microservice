package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/rezaAmiri123/microservice/payments/internal/app/commands"
	"github.com/rezaAmiri123/microservice/payments/paymentspb"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	grpcCodes "google.golang.org/grpc/codes"
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
	span := trace.SpanFromContext(ctx)

	id := uuid.New().String()

	span.SetAttributes(
		attribute.String("PaymentID", id),
		attribute.String("UserID", request.GetUserId()),
	)

	err := s.cfg.App.AuthorizePayment(ctx, commands.AuthorizePayment{
		ID:     id,
		UserID: request.GetUserId(),
		Amount: request.GetAmount(),
	})
	if err != nil {
		s.cfg.Logger.Errorf("failed to authorize a payment: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, status.Errorf(grpcCodes.Internal, "failed to authorize a payment: %s", err)
	}
	resp := &paymentspb.AuthorizePaymentResponse{Id: id}

	return resp, nil
}
