// Code generated by mockery v2.20.0. DO NOT EDIT.

package userspb

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockUserServiceClient is an autogenerated mock type for the UserServiceClient type
type MockUserServiceClient struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, in, opts
func (_m *MockUserServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *CreateUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *CreateUserRequest, ...grpc.CallOption) (*CreateUserResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *CreateUserRequest, ...grpc.CallOption) *CreateUserResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*CreateUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *CreateUserRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, in, opts
func (_m *MockUserServiceClient) Login(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *LoginUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *LoginUserRequest, ...grpc.CallOption) (*LoginUserResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *LoginUserRequest, ...grpc.CallOption) *LoginUserResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LoginUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *LoginUserRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginVerify provides a mock function with given fields: ctx, in, opts
func (_m *MockUserServiceClient) LoginVerify(ctx context.Context, in *LoginVerifyUserRequest, opts ...grpc.CallOption) (*LoginVerifyUserResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *LoginVerifyUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *LoginVerifyUserRequest, ...grpc.CallOption) (*LoginVerifyUserResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *LoginVerifyUserRequest, ...grpc.CallOption) *LoginVerifyUserResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LoginVerifyUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *LoginVerifyUserRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, in, opts
func (_m *MockUserServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *UpdateUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *UpdateUserRequest, ...grpc.CallOption) (*UpdateUserResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *UpdateUserRequest, ...grpc.CallOption) *UpdateUserResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UpdateUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *UpdateUserRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockUserServiceClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockUserServiceClient creates a new instance of MockUserServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUserServiceClient(t mockConstructorTestingTNewMockUserServiceClient) *MockUserServiceClient {
	mock := &MockUserServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}