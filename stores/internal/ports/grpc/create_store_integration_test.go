package grpc

import (
	"context"
	"fmt"
	"github.com/rezaAmiri123/microservice/stores/internal/domain"
	"github.com/rezaAmiri123/microservice/stores/storespb"
	"github.com/stretchr/testify/mock"
)

func (s *serverSuite) TestStoresService_CreateStore_OK() {
	s.mocks.app.On("CreateStore", mock.Anything, mock.AnythingOfType("commands.CreateStore")).Return(nil)

	_, err := s.client.CreateStore(context.Background(), &storespb.CreateStoreRequest{
		Name:     "store-name",
		Location: "store-location",
	})
	s.Assert().NoError(err)
}

func (s *serverSuite) TestStoresService_CreateStore_Error() {
	s.mocks.app.On("CreateStore", mock.Anything, mock.AnythingOfType("commands.CreateStore")).Return(domain.ErrStoreNameIsBlank)
	s.mocks.logger.On("Errorf", mock.Anything, mock.Anything)
	fmt.Println(s.mocks.logger)
	_, err := s.client.CreateStore(context.Background(), &storespb.CreateStoreRequest{
		Name:     "store-name",
		Location: "store-location",
	})
	s.Assert().Error(err)
}
