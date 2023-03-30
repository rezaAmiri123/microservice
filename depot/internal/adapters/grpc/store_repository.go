package grpc

import (
	"context"
	"github.com/rezaAmiri123/microservice/depot/internal/domain"
	pkgGrpc "github.com/rezaAmiri123/microservice/pkg/grpc"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/stores/storespb"
	"github.com/stackus/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type StoreRepository struct {
	endpoint string
	logger   logger.Logger
}

var _ domain.StoreRepository = (*StoreRepository)(nil)

func NewStoreRepository(endpoint string, logger logger.Logger) StoreRepository {
	return StoreRepository{
		endpoint: endpoint,
		logger:   logger,
	}
}

func (r StoreRepository) Find(ctx context.Context, storeID string) (*domain.Store, error) {
	var conn *grpc.ClientConn
	conn, err := r.dial(ctx)
	if err != nil {
		return nil, err
	}

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	resp, err := storespb.NewStoresServiceClient(conn).GetStore(ctx, &storespb.GetStoreRequest{
		Id: storeID,
	})

	if err != nil {
		if errors.GRPCCode(err) == codes.NotFound {
			return nil, errors.ErrNotFound.Msg("store was not located")
		}
		return nil, errors.Wrap(err, "requesting store")
	}
	return r.storeToDomain(resp.GetStore()), nil
}

func (r StoreRepository) storeToDomain(store *storespb.Store) *domain.Store {
	return &domain.Store{
		ID:       store.GetId(),
		Name:     store.GetName(),
		Location: store.GetLocation(),
	}
}
func (r StoreRepository) dial(ctx context.Context) (*grpc.ClientConn, error) {
	return pkgGrpc.NewContextGrpcClient(ctx, r.endpoint, nil, r.logger)
}
