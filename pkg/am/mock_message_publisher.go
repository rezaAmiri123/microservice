// Code generated by mockery v2.27.1. DO NOT EDIT.

package am

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockMessagePublisher is an autogenerated mock type for the MessagePublisher type
type MockMessagePublisher struct {
	mock.Mock
}

// Publish provides a mock function with given fields: ctx, topicName, msg
func (_m *MockMessagePublisher) Publish(ctx context.Context, topicName string, msg Message) error {
	ret := _m.Called(ctx, topicName, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, Message) error); ok {
		r0 = rf(ctx, topicName, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockMessagePublisher interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockMessagePublisher creates a new instance of MockMessagePublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockMessagePublisher(t mockConstructorTestingTNewMockMessagePublisher) *MockMessagePublisher {
	mock := &MockMessagePublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}