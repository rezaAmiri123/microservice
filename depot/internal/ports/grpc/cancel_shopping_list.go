package grpc

import (
	"context"
	"github.com/rezaAmiri123/microservice/depot/depotpb"
	"github.com/rezaAmiri123/microservice/depot/internal/app/commands"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func (s serverTx) CancelShoppingList(ctx context.Context, request *depotpb.CancelShoppingListRequest) (resp *depotpb.CancelShoppingListResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.CancelShoppingList(ctx, request)
}

func (s server) CancelShoppingList(ctx context.Context, request *depotpb.CancelShoppingListRequest) (*depotpb.CancelShoppingListResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("ShoppingListID", request.GetId()),
	)

	err := s.cfg.App.CancelShoppingList(ctx, commands.CancelShoppingList{
		ID: request.GetId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("CancelShoppingList: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	resp := &depotpb.CancelShoppingListResponse{}

	return resp, nil
}
