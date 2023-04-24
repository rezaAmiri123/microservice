// Code generated by mockery v2.24.0. DO NOT EDIT.

package app

import (
	context "context"

	commands "github.com/rezaAmiri123/microservice/baskets/internal/app/commands"

	mock "github.com/stretchr/testify/mock"
)

// MockApp is an autogenerated mock type for the App type
type MockApp struct {
	mock.Mock
}

// AddItem provides a mock function with given fields: ctx, add
func (_m *MockApp) AddItem(ctx context.Context, add commands.AddItem) error {
	ret := _m.Called(ctx, add)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.AddItem) error); ok {
		r0 = rf(ctx, add)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CancelBasket provides a mock function with given fields: ctx, cancel
func (_m *MockApp) CancelBasket(ctx context.Context, cancel commands.CancelBasket) error {
	ret := _m.Called(ctx, cancel)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.CancelBasket) error); ok {
		r0 = rf(ctx, cancel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckoutBasket provides a mock function with given fields: ctx, checkout
func (_m *MockApp) CheckoutBasket(ctx context.Context, checkout commands.CheckoutBasket) error {
	ret := _m.Called(ctx, checkout)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.CheckoutBasket) error); ok {
		r0 = rf(ctx, checkout)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartBasket provides a mock function with given fields: ctx, start
func (_m *MockApp) StartBasket(ctx context.Context, start commands.StartBasket) error {
	ret := _m.Called(ctx, start)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.StartBasket) error); ok {
		r0 = rf(ctx, start)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockApp interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockApp creates a new instance of MockApp. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockApp(t mockConstructorTestingTNewMockApp) *MockApp {
	mock := &MockApp{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
