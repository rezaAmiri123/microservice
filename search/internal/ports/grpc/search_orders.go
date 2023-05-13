package grpc

import (
	"context"
	"database/sql"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"github.com/rezaAmiri123/microservice/search/internal/constants"
	"github.com/rezaAmiri123/microservice/search/internal/domain"
	"github.com/rezaAmiri123/microservice/search/searchpb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func (s serverTx) SearchOrders(ctx context.Context, request *searchpb.SearchOrdersRequest) (resp *searchpb.SearchOrdersResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.SearchOrders(ctx, request)
}

func (s server) SearchOrders(ctx context.Context, request *searchpb.SearchOrdersRequest) (*searchpb.SearchOrdersResponse, error) {
	span := trace.SpanFromContext(ctx)

	filters := request.GetFilters()

	span.SetAttributes(
		attribute.String("UserID", filters.GetUserId()),
	)

	orders, err := s.cfg.App.SearchOrders(ctx, domain.SearchOrders{
		Filters: domain.Filters{
			// TODO there are some other fields
			UserID:     filters.GetUserId(),
			Status:     filters.GetStatus(),
			ProductIDs: filters.GetProductIds(),
			StoreIDs:   filters.GetStoreIds(),
			After:      filters.GetAfter().AsTime(),
			Before:     filters.GetBefore().AsTime(),
			MinTotal:   filters.GetMinTotal(),
			MaxTotal:   filters.GetMaxTotal(),
		},
		Next:  request.GetNext(),
		Limit: int(request.GetLimit()),
	})

	if err != nil {
		s.cfg.Logger.Errorf("SearchOrders: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	protoOrders := []*searchpb.Order{}
	for _, order := range orders {
		protoOrders = append(protoOrders, s.orderFromDomain(order))
	}
	resp := &searchpb.SearchOrdersResponse{
		Orders: protoOrders,
	}

	return resp, nil
}
