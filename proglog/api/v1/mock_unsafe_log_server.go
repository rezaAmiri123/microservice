// Code generated by mockery v2.33.0. DO NOT EDIT.

package v1

import mock "github.com/stretchr/testify/mock"

// MockUnsafeLogServer is an autogenerated mock type for the UnsafeLogServer type
type MockUnsafeLogServer struct {
	mock.Mock
}

// mustEmbedUnimplementedLogServer provides a mock function with given fields:
func (_m *MockUnsafeLogServer) mustEmbedUnimplementedLogServer() {
	_m.Called()
}

// NewMockUnsafeLogServer creates a new instance of MockUnsafeLogServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUnsafeLogServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUnsafeLogServer {
	mock := &MockUnsafeLogServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
