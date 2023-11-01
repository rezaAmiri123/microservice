// Code generated by mockery v2.33.0. DO NOT EDIT.

package v1

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockLogServer is an autogenerated mock type for the LogServer type
type MockLogServer struct {
	mock.Mock
}

// Consume provides a mock function with given fields: _a0, _a1
func (_m *MockLogServer) Consume(_a0 context.Context, _a1 *ConsumeRequest) (*ConsumeResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *ConsumeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ConsumeRequest) (*ConsumeResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ConsumeRequest) *ConsumeResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ConsumeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ConsumeRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServers provides a mock function with given fields: _a0, _a1
func (_m *MockLogServer) GetServers(_a0 context.Context, _a1 *GetServersRequest) (*GetServersResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *GetServersResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetServersRequest) (*GetServersResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetServersRequest) *GetServersResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetServersResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetServersRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Produce provides a mock function with given fields: _a0, _a1
func (_m *MockLogServer) Produce(_a0 context.Context, _a1 *ProduceRequest) (*ProduceResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *ProduceResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ProduceRequest) (*ProduceResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ProduceRequest) *ProduceResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ProduceResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ProduceRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedLogServer provides a mock function with given fields:
func (_m *MockLogServer) mustEmbedUnimplementedLogServer() {
	_m.Called()
}

// NewMockLogServer creates a new instance of MockLogServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLogServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLogServer {
	mock := &MockLogServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
