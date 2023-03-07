package grpc

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/stores/internal/app/commands"
	"github.com/rezaAmiri123/microservice/stores/internal/constants"
	"github.com/rezaAmiri123/microservice/stores/storespb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) CreateStore(ctx context.Context, request *storespb.CreateStoreRequest) (resp *storespb.CreateStoreResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.CreateStore(ctx, request)
}

func (s server) CreateStore(ctx context.Context, request *storespb.CreateStoreRequest) (*storespb.CreateStoreResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.CreateStore")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()
	id := uuid.New().String()
	err := s.cfg.App.Commands.CreateStore.Handle(ctx, commands.CreateStore{
		ID:       id,
		Name:     request.GetName(),
		Location: request.GetLocation(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("failed to create store: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to create store: %s", err)
	}

	resp := &storespb.CreateStoreResponse{Id: id}
	return resp, nil
}
