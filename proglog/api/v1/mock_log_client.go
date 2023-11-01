// Code generated by mockery v2.33.0. DO NOT EDIT.

package v1

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockLogClient is an autogenerated mock type for the LogClient type
type MockLogClient struct {
	mock.Mock
}

// Consume provides a mock function with given fields: ctx, in, opts
func (_m *MockLogClient) Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ConsumeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ConsumeRequest, ...grpc.CallOption) (*ConsumeResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ConsumeRequest, ...grpc.CallOption) *ConsumeResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ConsumeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ConsumeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServers provides a mock function with given fields: ctx, in, opts
func (_m *MockLogClient) GetServers(ctx context.Context, in *GetServersRequest, opts ...grpc.CallOption) (*GetServersResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *GetServersResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetServersRequest, ...grpc.CallOption) (*GetServersResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetServersRequest, ...grpc.CallOption) *GetServersResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetServersResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetServersRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Produce provides a mock function with given fields: ctx, in, opts
func (_m *MockLogClient) Produce(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*ProduceResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ProduceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ProduceRequest, ...grpc.CallOption) (*ProduceResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ProduceRequest, ...grpc.CallOption) *ProduceResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ProduceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ProduceRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockLogClient creates a new instance of MockLogClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLogClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLogClient {
	mock := &MockLogClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
