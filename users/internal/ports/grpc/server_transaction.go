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
	c      di.Container
	logger logger.Logger
	userspb.UnimplementedUserServiceServer
}

var _ userspb.UserServiceServer = (*serverTx)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	userspb.RegisterUserServiceServer(registrar, serverTx{
		c:      container,
		logger: container.Get(constants.LoggerKey).(logger.Logger),
	})
	return nil
}

func (s serverTx) getNextServer() server {
	cfg := &Config{
		App:    s.c.Get(constants.ApplicationKey).(app.App),
		Logger: s.logger,
	}
	return server{cfg: cfg}
}

func (s serverTx) closeTx(tx *sql.Tx, err error) error {
	if p := recover(); p != nil {
		_ = tx.Rollback()
		s.logger.Errorf("transaction rollback panic")
		panic(p)
	} else if err != nil {
		txErr := tx.Rollback()
		if txErr != nil {
			s.logger.Err("rollback error", txErr)
		} else {
			s.logger.Debug("transaction is rolled back")

		}
	} else {
		err = tx.Commit()
		s.logger.Debug("transaction is committed")
	}
	return err
}

func (s serverTx) RegisterUser(ctx context.Context, request *userspb.RegisterUserRequest) (resp *userspb.RegisterUserResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.RegisterUser(ctx, request)
}

func (s serverTx) GetUser(ctx context.Context, request *userspb.GetUserRequest) (resp *userspb.GetUserResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.GetUser(ctx, request)
}

func (s serverTx) EnableUser(ctx context.Context, request *userspb.EnableUserRequest) (resp *userspb.EnableUserResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.EnableUser(ctx, request)
}

func (s serverTx) DisableUser(ctx context.Context, request *userspb.DisableUserRequest) (resp *userspb.DisableUserResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.DisableUser(ctx, request)
}

func (s serverTx) AuthorizeUser(ctx context.Context, request *userspb.AuthorizeUserRequest) (resp *userspb.AuthorizeUserResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.AuthorizeUser(ctx, request)
}
