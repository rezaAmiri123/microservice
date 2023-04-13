package grpc

import (
	"context"
	"github.com/rezaAmiri123/microservice/baskets/basketspb"
	"github.com/rezaAmiri123/microservice/baskets/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/es"
	"github.com/stretchr/testify/mock"
)

func (s *serverSuite) TestBasketService_StartBasket() {
	s.mocks.baskets.On("Load", mock.Anything, mock.AnythingOfType("string")).Return(&domain.Basket{
		Aggregate: es.NewAggregate("basket-id", domain.BasketAggregate),
	}, nil)
	s.mocks.baskets.On("Save", mock.Anything, mock.AnythingOfType("*domain.Basket")).Return(nil)
	s.mocks.publisher.On("Publish", mock.Anything, mock.AnythingOfType("ddd.event")).Return(nil)

	_, err := s.client.StartBasket(context.Background(), &basketspb.StartBasketRequest{UserId: "user-id"})
	s.Assert().NoError(err)
}
