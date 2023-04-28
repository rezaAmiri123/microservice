package grpc

import (
	"context"
	"database/sql"
	"github.com/rezaAmiri123/microservice/pkg/di"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"github.com/rezaAmiri123/microservice/users/internal/app"
	"github.com/rezaAmiri123/microservice/users/internal/constants"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	grpcCodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) AuthorizeUser(ctx context.Context, request *userspb.AuthorizeUserRequest) (resp *userspb.AuthorizeUserResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.AuthorizeUser(ctx, request)
}

func (s *server) AuthorizeUser(ctx context.Context, req *userspb.AuthorizeUserRequest) (*userspb.AuthorizeUserResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("UserID", req.GetId()),
	)

	err := s.cfg.App.AuthorizeUser(ctx, app.AuthorizeUser{
		ID: req.GetId(),
	})
	if err != nil {
		//s.cfg.Logger.Errorf("failed to authorize user: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, status.Errorf(grpcCodes.Internal, "failed to authorize user: %s", err)
	}
	return &userspb.AuthorizeUserResponse{}, nil
}
