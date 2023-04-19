package grpc

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/users/internal/app/commands"
	"github.com/rezaAmiri123/microservice/users/internal/constants"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) RegisterUser(ctx context.Context, request *userspb.RegisterUserRequest) (resp *userspb.RegisterUserResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	//next := server{}
	//next.cfg = &Config{}
	//next.cfg.App = di.Get(ctx, constants.ApplicationKey).(*app.Application)
	//next.cfg.Logger = di.Get(ctx, constants.LoggerKey).(logger.Logger)
	next := s.getNextServer()
	return next.RegisterUser(ctx, request)
}

func (s *server) RegisterUser(ctx context.Context, req *userspb.RegisterUserRequest) (*userspb.RegisterUserResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.RegisterUser")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()
	id := uuid.New().String()
	err := s.cfg.App.Commands.RegisterUser.Handle(ctx, commands.RegisterUser{
		ID:       id,
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
	})
	if err != nil {
		//s.cfg.Logger.Errorf("failed to register user: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}
	return &userspb.RegisterUserResponse{Id: id}, nil
}
