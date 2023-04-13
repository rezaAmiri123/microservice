// Code generated by mockery v2.24.0. DO NOT EDIT.

package domain

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockProductRepository is an autogenerated mock type for the ProductRepository type
type MockProductRepository struct {
	mock.Mock
}

// Find provides a mock function with given fields: ctx, productID
func (_m *MockProductRepository) Find(ctx context.Context, productID string) (*Product, error) {
	ret := _m.Called(ctx, productID)

	var r0 *Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*Product, error)); ok {
		return rf(ctx, productID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *Product); ok {
		r0 = rf(ctx, productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockProductRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockProductRepository creates a new instance of MockProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockProductRepository(t mockConstructorTestingTNewMockProductRepository) *MockProductRepository {
	mock := &MockProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
