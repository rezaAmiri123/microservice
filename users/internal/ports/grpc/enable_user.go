package grpc

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"github.com/rezaAmiri123/microservice/users/internal/app"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	grpcCodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) EnableUser(ctx context.Context, request *userspb.EnableUserRequest) (resp *userspb.EnableUserResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.EnableUser(ctx, request)
}

func (s *server) EnableUser(ctx context.Context, req *userspb.EnableUserRequest) (*userspb.EnableUserResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("UserID", req.GetId()),
	)

	err := s.cfg.App.EnableUser(ctx, app.EnableUser{
		ID: req.GetId(),
	})
	if err != nil {
		s.cfg.Logger.Errorf("failed to authorize user: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, status.Errorf(grpcCodes.Internal, "failed to authorize user: %s", err)
	}
	return &userspb.EnableUserResponse{}, nil
}
