package grpc

import (
	"context"
	"database/sql"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"github.com/rezaAmiri123/microservice/stores/internal/app/commands"
	"github.com/rezaAmiri123/microservice/stores/internal/constants"
	"github.com/rezaAmiri123/microservice/stores/storespb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func (s serverTx) RebrandStore(ctx context.Context, request *storespb.RebrandStoreRequest) (resp *storespb.RebrandStoreResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.RebrandStore(ctx, request)
}

func (s server) RebrandStore(ctx context.Context, request *storespb.RebrandStoreRequest) (*storespb.RebrandStoreResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("StoreID", request.GetId()),
	)

	err := s.cfg.App.RebrandStore(ctx, commands.RebrandStore{
		ID:   request.GetId(),
		Name: request.GetName(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("RebrandStore: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	resp := &storespb.RebrandStoreResponse{}
	return resp, nil
}
