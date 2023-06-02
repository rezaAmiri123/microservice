package grpc

import (
	"context"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
	"github.com/rezaAmiri123/microservice/stores/storespb"
	"github.com/stretchr/testify/mock"
)

func (s *serverSuite) TestStoresService_AddProduct_OK() {
	s.mocks.app.On("AddProduct", mock.Anything, mock.AnythingOfType("commands.AddProduct")).Return(nil)

	_, err := s.client.AddProduct(context.Background(), &storespb.AddProductRequest{
		Name:        "product-name",
		StoreId:     "store-id",
		Price:       854.345,
		Description: "description",
		Sku:         "sku",
	})
	s.Assert().NoError(err)
}

func (s *serverSuite) TestStoresService_AddProduct_Error() {
	s.mocks.app.On("AddProduct", mock.Anything, mock.AnythingOfType("commands.AddProduct")).Return(domain.ErrStoreNameIsBlank)
	s.mocks.logger.On("Errorf", mock.Anything, mock.Anything)

	_, err := s.client.AddProduct(context.Background(), &storespb.AddProductRequest{
		Name:        "product-name",
		StoreId:     "store-id",
		Price:       854.345,
		Description: "description",
		Sku:         "sku",
	})
	s.Assert().Error(err)
}
