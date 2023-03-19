package grpc

import (
	"context"
	"database/sql"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/stores/internal/app/queries"
	"github.com/rezaAmiri123/microservice/stores/internal/constants"
	"github.com/rezaAmiri123/microservice/stores/storespb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) GetStore(ctx context.Context, request *storespb.GetStoreRequest) (resp *storespb.GetStoreResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.GetStore(ctx, request)
}

func (s server) GetStore(ctx context.Context, request *storespb.GetStoreRequest) (*storespb.GetStoreResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.GetStore")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()

	store, err := s.cfg.App.Queries.GetStore.Handle(ctx, queries.GetStore{
		ID: request.GetId(),
	})
	if err != nil {
		s.cfg.Logger.Errorf("failed to get store: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.NotFound, "failed to get store: %s", err)
	}

	return &storespb.GetStoreResponse{Store: s.storeFromDomain(store)}, nil
}
