// Code generated by mockery v2.27.1. DO NOT EDIT.

package notificationspb

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockNotificationsServiceClient is an autogenerated mock type for the NotificationsServiceClient type
type MockNotificationsServiceClient struct {
	mock.Mock
}

// NotifyOrderCanceled provides a mock function with given fields: ctx, in, opts
func (_m *MockNotificationsServiceClient) NotifyOrderCanceled(ctx context.Context, in *NotifyOrderCanceledRequest, opts ...grpc.CallOption) (*NotifyOrderCanceledResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *NotifyOrderCanceledResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *NotifyOrderCanceledRequest, ...grpc.CallOption) (*NotifyOrderCanceledResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *NotifyOrderCanceledRequest, ...grpc.CallOption) *NotifyOrderCanceledResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*NotifyOrderCanceledResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *NotifyOrderCanceledRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NotifyOrderCreated provides a mock function with given fields: ctx, in, opts
func (_m *MockNotificationsServiceClient) NotifyOrderCreated(ctx context.Context, in *NotifyOrderCreatedRequest, opts ...grpc.CallOption) (*NotifyOrderCreatedResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *NotifyOrderCreatedResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *NotifyOrderCreatedRequest, ...grpc.CallOption) (*NotifyOrderCreatedResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *NotifyOrderCreatedRequest, ...grpc.CallOption) *NotifyOrderCreatedResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*NotifyOrderCreatedResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *NotifyOrderCreatedRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NotifyOrderReady provides a mock function with given fields: ctx, in, opts
func (_m *MockNotificationsServiceClient) NotifyOrderReady(ctx context.Context, in *NotifyOrderReadyRequest, opts ...grpc.CallOption) (*NotifyOrderReadyResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *NotifyOrderReadyResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *NotifyOrderReadyRequest, ...grpc.CallOption) (*NotifyOrderReadyResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *NotifyOrderReadyRequest, ...grpc.CallOption) *NotifyOrderReadyResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*NotifyOrderReadyResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *NotifyOrderReadyRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockNotificationsServiceClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockNotificationsServiceClient creates a new instance of MockNotificationsServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockNotificationsServiceClient(t mockConstructorTestingTNewMockNotificationsServiceClient) *MockNotificationsServiceClient {
	mock := &MockNotificationsServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
