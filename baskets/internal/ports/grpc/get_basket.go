package grpc

import (
	"context"
	"fmt"
	"github.com/rezaAmiri123/microservice/baskets/basketspb"
	"github.com/rezaAmiri123/microservice/baskets/internal/app/queries"
	"github.com/rezaAmiri123/microservice/pkg/errorsotel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func (s server) GetBasket(ctx context.Context, request *basketspb.GetBasketRequest) (*basketspb.GetBasketResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("BasketID", request.GetId()),
	)

	basket, err := s.cfg.App.GetBasket(ctx, queries.GetBasket{
		ID: request.GetId(),
	})

	if err != nil {
		s.cfg.Logger.Errorf("GetBasket: %s", err)
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	fmt.Println("userid: ", basket.UserID)
	fmt.Println("status: ", basket.Status)
	resp := &basketspb.GetBasketResponse{
		Basket: s.basketFromDomain(basket),
	}

	return resp, nil
}
