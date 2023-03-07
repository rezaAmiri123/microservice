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

func (s serverTx) AddProduct(ctx context.Context, request *storespb.AddProductRequest) (resp *storespb.AddProductResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.AddProduct(ctx, request)
}

func (s server) AddProduct(ctx context.Context, request *storespb.AddProductRequest) (*storespb.AddProductResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.AddProduct")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()
	id := uuid.New().String()
	err := s.cfg.App.Commands.AddProduct.Handle(ctx, commands.AddProduct{
		ID:          id,
		StoreID:     request.GetStoreId(),
		Name:        request.GetName(),
		Description: request.GetDescription(),
		SKU:         request.GetSku(),
		Price:       request.GetPrice(),
	})
	if err != nil {
		s.cfg.Logger.Errorf("failed to add product: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to add product: %s", err)
	}
	resp := &storespb.AddProductResponse{Id: id}

	return resp, nil
}
