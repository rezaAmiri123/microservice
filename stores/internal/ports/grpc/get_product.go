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

func (s serverTx) GetProduct(ctx context.Context, request *storespb.GetProductRequest) (resp *storespb.GetProductResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.GetProduct(ctx, request)
}

func (s server) GetProduct(ctx context.Context, request *storespb.GetProductRequest) (*storespb.GetProductResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.GetProduct")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()

	product, err := s.cfg.App.Queries.GetProduct.Handle(ctx, queries.GetProduct{
		ID: request.GetId(),
	})
	if err != nil {
		s.cfg.Logger.Errorf("failed to get product: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.NotFound, "failed to get product: %s", err)
	}

	return &storespb.GetProductResponse{Product: s.productFromDomain(product)}, nil
}
