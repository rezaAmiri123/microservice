// Code generated by mockery v2.24.0. DO NOT EDIT.

package basketspb

import mock "github.com/stretchr/testify/mock"

// MockUnsafeBasketServiceServer is an autogenerated mock type for the UnsafeBasketServiceServer type
type MockUnsafeBasketServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedBasketServiceServer provides a mock function with given fields:
func (_m *MockUnsafeBasketServiceServer) mustEmbedUnimplementedBasketServiceServer() {
	_m.Called()
}

type mockConstructorTestingTNewMockUnsafeBasketServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockUnsafeBasketServiceServer creates a new instance of MockUnsafeBasketServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUnsafeBasketServiceServer(t mockConstructorTestingTNewMockUnsafeBasketServiceServer) *MockUnsafeBasketServiceServer {
	mock := &MockUnsafeBasketServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
