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

func (s serverTx) CreateInvoice(ctx context.Context, request *paymentspb.CreateInvoiceRequest) (resp *paymentspb.CreateInvoiceResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.CreateInvoice(ctx, request)
}

func (s server) CreateInvoice(ctx context.Context, request *paymentspb.CreateInvoiceRequest) (*paymentspb.CreateInvoiceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.CreateInvoice")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()
	id := uuid.New().String()
	err := s.cfg.App.Commands.CreateInvoice.Handle(ctx, commands.CreateInvoice{
		ID:        id,
		OrderID:   request.GetOrderId(),
		PaymentID: request.GetPaymentId(),
		Amount:    request.GetAmount(),
	})
	if err != nil {
		s.cfg.Logger.Errorf("failed to create invoice: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to create invoice: %s", err)
	}
	resp := &paymentspb.CreateInvoiceResponse{Id: id}

	return resp, nil
}
