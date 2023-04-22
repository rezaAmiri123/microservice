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

func (s serverTx) CreateInvoice(ctx context.Context, request *paymentspb.CreateInvoiceRequest) (resp *paymentspb.CreateInvoiceResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.CreateInvoice(ctx, request)
}

func (s server) CreateInvoice(ctx context.Context, request *paymentspb.CreateInvoiceRequest) (*paymentspb.CreateInvoiceResponse, error) {
	span := trace.SpanFromContext(ctx)

	id := uuid.New().String()

	span.SetAttributes(
		attribute.String("InvoiceID", id),
		attribute.String("OrderID", request.GetOrderId()),
	)

	err := s.cfg.App.CreateInvoice(ctx, commands.CreateInvoice{
		ID:        id,
		OrderID:   request.GetOrderId(),
		PaymentID: request.GetPaymentId(),
		Amount:    request.GetAmount(),
	})
	if err != nil {
		s.cfg.Logger.Errorf("failed to create invoice: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, status.Errorf(grpcCodes.Internal, "failed to create invoice: %s", err)
	}
	resp := &paymentspb.CreateInvoiceResponse{Id: id}

	return resp, nil
}
