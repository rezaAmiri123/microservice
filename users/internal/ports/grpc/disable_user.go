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

func (s serverTx) DisableUser(ctx context.Context, request *userspb.DisableUserRequest) (resp *userspb.DisableUserResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.DisableUser(ctx, request)
}

func (s *server) DisableUser(ctx context.Context, req *userspb.DisableUserRequest) (*userspb.DisableUserResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("UserID", req.GetId()),
	)

	err := s.cfg.App.DisableUser(ctx, app.DisableUser{
		ID: req.GetId(),
	})
	if err != nil {
		s.cfg.Logger.Errorf("failed to disable user: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, status.Errorf(grpcCodes.Internal, "failed to disable user: %s", err)
	}
	return &userspb.DisableUserResponse{}, nil
}
