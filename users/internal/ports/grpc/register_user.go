package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"github.com/rezaAmiri123/microservice/users/internal/app"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	grpcCodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serverTx) RegisterUser(ctx context.Context, request *userspb.RegisterUserRequest) (resp *userspb.RegisterUserResponse, err error) {
	ctx = s.c.Scoped(ctx)
	//defer func(tx *sql.Tx) {
	//	err = s.closeTx(tx, err)
	//}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := s.getNextServer()
	return next.RegisterUser(ctx, request)
}

func (s *server) RegisterUser(ctx context.Context, req *userspb.RegisterUserRequest) (*userspb.RegisterUserResponse, error) {
	span := trace.SpanFromContext(ctx)

	id := uuid.New().String()
	span.SetAttributes(
		attribute.String("UserID", id),
	)

	err := s.cfg.App.RegisterUser(ctx, app.RegisterUser{
		ID:       id,
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
	})
	if err != nil {
		//s.cfg.Logger.Errorf("failed to register user: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, status.Errorf(grpcCodes.Internal, "failed to create user: %s", err)
	}
	return &userspb.RegisterUserResponse{Id: id}, nil
}
