// Code generated by mockery v2.27.1. DO NOT EDIT.

package orderingpb

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockOrderingServiceServer is an autogenerated mock type for the OrderingServiceServer type
type MockOrderingServiceServer struct {
	mock.Mock
}

// CancelOrder provides a mock function with given fields: _a0, _a1
func (_m *MockOrderingServiceServer) CancelOrder(_a0 context.Context, _a1 *CancelOrderRequest) (*CancelOrderResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *CancelOrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *CancelOrderRequest) *CancelOrderResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*CancelOrderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *CancelOrderRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CompleteOrder provides a mock function with given fields: _a0, _a1
func (_m *MockOrderingServiceServer) CompleteOrder(_a0 context.Context, _a1 *CompleteOrderRequest) (*CompleteOrderResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *CompleteOrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *CompleteOrderRequest) (*CompleteOrderResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *CompleteOrderRequest) *CompleteOrderResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*CompleteOrderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *CompleteOrderRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrder provides a mock function with given fields: _a0, _a1
func (_m *MockOrderingServiceServer) CreateOrder(_a0 context.Context, _a1 *CreateOrderRequest) (*CreateOrderResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *CreateOrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *CreateOrderRequest) *CreateOrderResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*CreateOrderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *CreateOrderRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrder provides a mock function with given fields: _a0, _a1
func (_m *MockOrderingServiceServer) GetOrder(_a0 context.Context, _a1 *GetOrderRequest) (*GetOrderResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *GetOrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetOrderRequest) (*GetOrderResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetOrderRequest) *GetOrderResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetOrderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetOrderRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadyOrder provides a mock function with given fields: _a0, _a1
func (_m *MockOrderingServiceServer) ReadyOrder(_a0 context.Context, _a1 *ReadyOrderRequest) (*ReadyOrderResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *ReadyOrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ReadyOrderRequest) (*ReadyOrderResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ReadyOrderRequest) *ReadyOrderResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ReadyOrderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ReadyOrderRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedOrderingServiceServer provides a mock function with given fields:
func (_m *MockOrderingServiceServer) mustEmbedUnimplementedOrderingServiceServer() {
	_m.Called()
}

type mockConstructorTestingTNewMockOrderingServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockOrderingServiceServer creates a new instance of MockOrderingServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockOrderingServiceServer(t mockConstructorTestingTNewMockOrderingServiceServer) *MockOrderingServiceServer {
	mock := &MockOrderingServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
