package grpc

import (
	"context"
	"database/sql"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/users/internal/app/commands"
	"github.com/rezaAmiri123/microservice/users/internal/constants"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) EnableUser(ctx context.Context, request *userspb.EnableUserRequest) (resp *userspb.EnableUserResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	//next := server{}
	//next.cfg = &Config{}
	//next.cfg.App = di.Get(ctx, constants.ApplicationKey).(*app.Application)
	//next.cfg.Logger = di.Get(ctx, constants.LoggerKey).(logger.Logger)
	next := s.getNextServer()
	return next.EnableUser(ctx, request)
}

func (s *server) EnableUser(ctx context.Context, req *userspb.EnableUserRequest) (*userspb.EnableUserResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "server.AuthorizeUser")
	defer span.Finish()

	//s.cfg.Metric.CreateUserGrpcRequests.Inc()
	err := s.cfg.App.Commands.EnableUser.Handle(ctx, commands.EnableUser{
		ID: req.GetId(),
	})
	if err != nil {
		s.cfg.Logger.Errorf("failed to authorize user: %s", err)
		//s.cfg.Metric.ErrorGrpcRequests.Inc()
		return nil, status.Errorf(codes.Internal, "failed to authorize user: %s", err)
	}
	return &userspb.EnableUserResponse{}, nil
}
