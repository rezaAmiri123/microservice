// Code generated by mockery v2.27.1. DO NOT EDIT.

package am

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockMessageHandler is an autogenerated mock type for the MessageHandler type
type MockMessageHandler struct {
	mock.Mock
}

// HandleMessage provides a mock function with given fields: ctx, msg
func (_m *MockMessageHandler) HandleMessage(ctx context.Context, msg IncomingMessage) error {
	ret := _m.Called(ctx, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, IncomingMessage) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockMessageHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockMessageHandler creates a new instance of MockMessageHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockMessageHandler(t mockConstructorTestingTNewMockMessageHandler) *MockMessageHandler {
	mock := &MockMessageHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
