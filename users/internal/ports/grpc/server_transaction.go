package grpc

import (
	"context"
	"database/sql"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/logger"
	"github.com/rezaAmiri123/microservice/users/internal/app"
	"github.com/rezaAmiri123/microservice/users/internal/constants"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"google.golang.org/grpc"
)

type serverTx struct {
	c di.Container
	userspb.UnimplementedUserServiceServer
}

var _ userspb.UserServiceServer = (*serverTx)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	userspb.RegisterUserServiceServer(registrar, serverTx{
		c: container,
	})
	return nil
}

func (s serverTx) RegisterUser(ctx context.Context, request *userspb.RegisterUserRequest) (resp *userspb.RegisterUserResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := server{}
	next.cfg = &Config{}
	next.cfg.App = di.Get(ctx, constants.ApplicationKey).(*app.Application)
	next.cfg.Logger = di.Get(ctx, constants.LoggerKey).(logger.Logger)
	return next.RegisterUser(ctx, request)
}

//func (s serverTx) AuthorizeCustomer(ctx context.Context, request *customerspb.AuthorizeCustomerRequest) (resp *customerspb.AuthorizeCustomerResponse, err error) {
//	ctx = s.c.Scoped(ctx)
//	defer func(tx *sql.Tx) {
//		err = s.closeTx(tx, err)
//	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))
//
//	next := server{app: di.Get(ctx, constants.ApplicationKey).(application.App)}
//
//	return next.AuthorizeCustomer(ctx, request)
//}

//func (s serverTx) GetCustomer(ctx context.Context, request *customerspb.GetCustomerRequest) (resp *customerspb.GetCustomerResponse, err error) {
//	ctx = s.c.Scoped(ctx)
//	defer func(tx *sql.Tx) {
//		err = s.closeTx(tx, err)
//	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))
//
//	next := server{app: di.Get(ctx, constants.ApplicationKey).(application.App)}
//
//	return next.GetCustomer(ctx, request)
//}

//func (s serverTx) EnableCustomer(ctx context.Context, request *customerspb.EnableCustomerRequest) (resp *customerspb.EnableCustomerResponse, err error) {
//	ctx = s.c.Scoped(ctx)
//	defer func(tx *sql.Tx) {
//		err = s.closeTx(tx, err)
//	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))
//
//	next := server{app: di.Get(ctx, constants.ApplicationKey).(application.App)}
//
//	return next.EnableCustomer(ctx, request)
//}
//
//func (s serverTx) DisableCustomer(ctx context.Context, request *customerspb.DisableCustomerRequest) (resp *customerspb.DisableCustomerResponse, err error) {
//	ctx = s.c.Scoped(ctx)
//	defer func(tx *sql.Tx) {
//		err = s.closeTx(tx, err)
//	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))
//
//	next := server{app: di.Get(ctx, constants.ApplicationKey).(application.App)}
//
//	return next.DisableCustomer(ctx, request)
//}

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
