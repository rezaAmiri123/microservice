package grpc

import (
	"context"
	"database/sql"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/stores/internal/app"
	"github.com/rezaAmiri123/microservice/stores/internal/constants"
	"github.com/rezaAmiri123/microservice/stores/storespb"
	"google.golang.org/grpc"
)

type serverTx struct {
	c di.Container
	storespb.UnimplementedStoresServiceServer
}

var _ storespb.StoresServiceServer = (*serverTx)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	storespb.RegisterStoresServiceServer(registrar, serverTx{
		c: container,
	})
	return nil
}

func (s serverTx) getNextServer() server {
	cfg := &Config{
		App:    s.c.Get(constants.ApplicationKey).(app.App),
		Logger: s.c.Get(constants.LoggerKey).(logger.Logger),
	}
	return server{cfg: cfg}
}

func (s serverTx) closeTx(tx *sql.Tx, err error) error {
	if p := recover(); p != nil {
		_ = tx.Rollback()
		panic(p)
	} else if err != nil {
		_ = tx.Rollback()
		return err
	} else {
		return tx.Commit()
	}
}

func (s serverTx) CreateStore(ctx context.Context, request *storespb.CreateStoreRequest) (resp *storespb.CreateStoreResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.CreateStore(ctx, request)
}

func (s serverTx) AddProduct(ctx context.Context, request *storespb.AddProductRequest) (resp *storespb.AddProductResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.AddProduct(ctx, request)
}

func (s serverTx) RemoveProduct(ctx context.Context, request *storespb.RemoveProductRequest) (resp *storespb.RemoveProductResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.RemoveProduct(ctx, request)
}

func (s serverTx) RebrandStore(ctx context.Context, request *storespb.RebrandStoreRequest) (resp *storespb.RebrandStoreResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.RebrandStore(ctx, request)
}

func (s serverTx) RebrandProduct(ctx context.Context, request *storespb.RebrandProductRequest) (resp *storespb.RebrandProductResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.RebrandProduct(ctx, request)
}

func (s serverTx) IncreaseProductPrice(ctx context.Context, request *storespb.IncreaseProductPriceRequest) (resp *storespb.IncreaseProductPriceResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.IncreaseProductPrice(ctx, request)
}

func (s serverTx) GetStores(ctx context.Context, request *storespb.GetStoresRequest) (resp *storespb.GetStoresResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.GetStores(ctx, request)
}

func (s serverTx) GetStore(ctx context.Context, request *storespb.GetStoreRequest) (resp *storespb.GetStoreResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.GetStore(ctx, request)
}

func (s serverTx) GetProduct(ctx context.Context, request *storespb.GetProductRequest) (resp *storespb.GetProductResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.GetProduct(ctx, request)
}

func (s serverTx) GetParticipatingStores(ctx context.Context, request *storespb.GetParticipatingStoresRequest) (resp *storespb.GetParticipatingStoresResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.GetParticipatingStores(ctx, request)
}

func (s serverTx) GetCatalog(ctx context.Context, request *storespb.GetCatalogRequest) (resp *storespb.GetCatalogResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.GetCatalog(ctx, request)
}

func (s serverTx) EnableParticipation(ctx context.Context, request *storespb.EnableParticipationRequest) (resp *storespb.EnableParticipationResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.EnableParticipation(ctx, request)
}

func (s serverTx) DisableParticipation(ctx context.Context, request *storespb.DisableParticipationRequest) (resp *storespb.DisableParticipationResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.DisableParticipation(ctx, request)
}

func (s serverTx) DecreaseProductPrice(ctx context.Context, request *storespb.DecreaseProductPriceRequest) (resp *storespb.DecreaseProductPriceResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.DecreaseProductPrice(ctx, request)
}
