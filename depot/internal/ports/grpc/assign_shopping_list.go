package grpc

import (
	"context"
	"github.com/rezaAmiri123/microservice/depot/depotpb"
	"github.com/rezaAmiri123/microservice/depot/internal/app/commands"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	grpcCodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) AssignShoppingList(ctx context.Context, request *depotpb.AssignShoppingListRequest) (resp *depotpb.AssignShoppingListResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.AssignShoppingList(ctx, request)
}

func (s server) AssignShoppingList(ctx context.Context, request *depotpb.AssignShoppingListRequest) (*depotpb.AssignShoppingListResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("ShoppingListID", request.GetId()),
	)

	err := s.cfg.App.AssignShoppingList(ctx, commands.AssignShoppingList{
		ID:    request.GetId(),
		BotID: request.GetBotId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to assign shopping list: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, status.Errorf(grpcCodes.Internal, "failed to assign shopping list: %s", err)
	}

	resp := &depotpb.AssignShoppingListResponse{}

	return resp, nil
}
