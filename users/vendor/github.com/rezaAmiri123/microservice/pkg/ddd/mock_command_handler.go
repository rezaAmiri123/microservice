// Code generated by mockery v2.27.1. DO NOT EDIT.

package ddd

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockCommandHandler is an autogenerated mock type for the CommandHandler type
type MockCommandHandler[T Command] struct {
	mock.Mock
}

// HandleCommand provides a mock function with given fields: ctx, cmd
func (_m *MockCommandHandler[T]) HandleCommand(ctx context.Context, cmd T) (Reply, error) {
	ret := _m.Called(ctx, cmd)

	var r0 Reply
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, T) (Reply, error)); ok {
		return rf(ctx, cmd)
	}
	if rf, ok := ret.Get(0).(func(context.Context, T) Reply); ok {
		r0 = rf(ctx, cmd)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Reply)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, T) error); ok {
		r1 = rf(ctx, cmd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockCommandHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockCommandHandler creates a new instance of MockCommandHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockCommandHandler[T Command](t mockConstructorTestingTNewMockCommandHandler) *MockCommandHandler[T] {
	mock := &MockCommandHandler[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
