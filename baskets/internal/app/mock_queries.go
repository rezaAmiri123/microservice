// Code generated by mockery v2.27.1. DO NOT EDIT.

package app

import (
	context "context"

	domain "github.com/rezaAmiri123/microservice/baskets/internal/domain"
	mock "github.com/stretchr/testify/mock"

	queries "github.com/rezaAmiri123/microservice/baskets/internal/app/queries"
)

// MockQueries is an autogenerated mock type for the Queries type
type MockQueries struct {
	mock.Mock
}

// GetBasket provides a mock function with given fields: ctx, cmd
func (_m *MockQueries) GetBasket(ctx context.Context, cmd queries.GetBasket) (*domain.Basket, error) {
	ret := _m.Called(ctx, cmd)

	var r0 *domain.Basket
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, queries.GetBasket) (*domain.Basket, error)); ok {
		return rf(ctx, cmd)
	}
	if rf, ok := ret.Get(0).(func(context.Context, queries.GetBasket) *domain.Basket); ok {
		r0 = rf(ctx, cmd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Basket)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, queries.GetBasket) error); ok {
		r1 = rf(ctx, cmd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockQueries interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockQueries creates a new instance of MockQueries. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockQueries(t mockConstructorTestingTNewMockQueries) *MockQueries {
	mock := &MockQueries{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
