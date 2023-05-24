package grpc

import (
	"context"
	"github.com/rezaAmiri123/microservice/baskets/basketspb"
	"github.com/rezaAmiri123/microservice/baskets/internal/domain"
	"github.com/rezaAmiri123/microservice/pkg/es"
	"github.com/stretchr/testify/mock"
)

func (s *serverSuite) TestBasketService_AddItem() {
	product := &domain.Product{
		ID:      "product-id",
		StoreID: "store-id",
		Name:    "product-name",
		Price:   10.00,
	}
	store := &domain.Store{
		ID:   "store-id",
		Name: "store-name",
	}
	s.mocks.baskets.On("Load", mock.Anything, "basket-id").Return(&domain.Basket{
		Aggregate: es.NewAggregate("basket-id", domain.BasketAggregate),
		UserID:    "user-id",
		Items: map[string]domain.Item{
			"product-id": {
				StoreID:      "store-id",
				ProductID:    "product-id",
				StoreName:    "store-name",
				ProductName:  "product-name",
				ProductPrice: 1.00,
				Quantity:     1,
			},
		},
		Status: domain.BasketIsOpen,
	}, nil)
	s.mocks.baskets.On("Save", mock.Anything, mock.AnythingOfType("*domain.Basket")).Return(nil)
	s.mocks.products.On("Find", mock.Anything, "product-id").Return(product, nil)
	s.mocks.stores.On("Find", mock.Anything, "store-id").Return(store, nil)

	_, err := s.client.AddItem(context.Background(), &basketspb.AddItemRequest{
		Id:        "basket-id",
		ProductId: "product-id",
		Quantity:  1,
	})
	s.Assert().NoError(err)
}
