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

func (s *server) AuthorizeUser(ctx context.Context, req *userspb.AuthorizeUserRequest) (*userspb.AuthorizeUserResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("UserID", req.GetId()),
	)

	err := s.cfg.App.AuthorizeUser(ctx, app.AuthorizeUser{
		ID: req.GetId(),
	})
	if err != nil {
		err = errors.Wrap(err, "grpc AuthorizeUser")
		s.cfg.Logger.Error(err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	s.cfg.Logger.Debugf("authorize user %s ", req.GetId())

	return &userspb.AuthorizeUserResponse{}, nil
}
