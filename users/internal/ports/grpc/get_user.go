package grpc

import (
	"context"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"github.com/rezaAmiri123/microservice/users/internal/app"
	"github.com/rezaAmiri123/microservice/users/userspb"
	"github.com/stackus/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func (s *server) GetUser(ctx context.Context, req *userspb.GetUserRequest) (*userspb.GetUserResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("UserID", req.GetId()),
	)
	user, err := s.cfg.App.GetUser(ctx, app.GetUser{
		ID: req.GetId(),
	})
	if err != nil {
		err = errors.Wrap(err, "grpc GetUser")
		s.cfg.Logger.Error(err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	s.cfg.Logger.Debugf("retrieve user %s ", req.GetId())

	return &userspb.GetUserResponse{User: s.userFromDomain(user)}, nil
}
