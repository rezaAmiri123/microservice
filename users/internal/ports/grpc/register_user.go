package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"github.com/rezaAmiri123/microservice/users/internal/app"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"github.com/stackus/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

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
		err = errors.Wrap(err, "grpc RegisterUser")
		s.cfg.Logger.Error(err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	s.cfg.Logger.Debugf("user %s is registered", req.GetUsername())

	return &userspb.RegisterUserResponse{Id: id}, nil
}
