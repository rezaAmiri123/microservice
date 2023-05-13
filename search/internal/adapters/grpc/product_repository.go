package grpc

import (
	"context"
	pkgGrpc "github.com/rezaAmiri123/microservice/pkg/grpc"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/search/internal/domain"
	"github.com/rezaAmiri123/microservice/stores/storespb"
	"google.golang.org/grpc"
)

type ProductRepository struct {
	endpoint string
	logger   logger.Logger
}

var _ domain.ProductRepository = (*ProductRepository)(nil)

func NewProductRepository(endpoint string, logger logger.Logger) ProductRepository {
	return ProductRepository{
		endpoint: endpoint,
		logger:   logger,
	}
}

func (r ProductRepository) Find(ctx context.Context, productID string) (product *domain.Product, err error) {
	var conn *grpc.ClientConn
	conn, err = r.dial(ctx)
	if err != nil {
		return nil, err
	}

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	resp, err := storespb.NewStoresServiceClient(conn).GetProduct(ctx, &storespb.GetProductRequest{Id: productID})
	if err != nil {
		return nil, err
	}

	return r.productToDomain(resp.Product), nil
}

func (r ProductRepository) productToDomain(product *storespb.Product) *domain.Product {
	return &domain.Product{
		ID:      product.GetId(),
		Name:    product.GetName(),
		StoreID: product.GetStoreId(),
	}
}

func (r ProductRepository) dial(ctx context.Context) (*grpc.ClientConn, error) {
	return pkgGrpc.NewContextGrpcClient(ctx, r.endpoint, nil, r.logger)
}
