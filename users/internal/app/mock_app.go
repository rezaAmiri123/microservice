// Code generated by mockery v2.27.1. DO NOT EDIT.

package app

import (
	context "context"

	domain "github.com/rezaAmiri123/microservice/users/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockApp is an autogenerated mock type for the App type
type MockApp struct {
	mock.Mock
}

// AuthorizeUser provides a mock function with given fields: ctx, cmd
func (_m *MockApp) AuthorizeUser(ctx context.Context, cmd AuthorizeUser) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, AuthorizeUser) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DisableUser provides a mock function with given fields: ctx, cmd
func (_m *MockApp) DisableUser(ctx context.Context, cmd DisableUser) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, DisableUser) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnableUser provides a mock function with given fields: ctx, cmd
func (_m *MockApp) EnableUser(ctx context.Context, cmd EnableUser) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, EnableUser) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUser provides a mock function with given fields: ctx, cmd
func (_m *MockApp) GetUser(ctx context.Context, cmd GetUser) (*domain.User, error) {
	ret := _m.Called(ctx, cmd)

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, GetUser) (*domain.User, error)); ok {
		return rf(ctx, cmd)
	}
	if rf, ok := ret.Get(0).(func(context.Context, GetUser) *domain.User); ok {
		r0 = rf(ctx, cmd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, GetUser) error); ok {
		r1 = rf(ctx, cmd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: ctx, cmd
func (_m *MockApp) RegisterUser(ctx context.Context, cmd RegisterUser) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, RegisterUser) error); ok {
		r0 = rf(ctx, cmd)
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