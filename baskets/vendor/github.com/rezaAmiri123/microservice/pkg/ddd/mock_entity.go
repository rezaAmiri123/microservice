// Code generated by mockery v2.27.1. DO NOT EDIT.

package ddd

import mock "github.com/stretchr/testify/mock"

// MockEntity is an autogenerated mock type for the Entity type
type MockEntity struct {
	mock.Mock
}

// EntityName provides a mock function with given fields:
func (_m *MockEntity) EntityName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *MockEntity) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SetID provides a mock function with given fields: _a0
func (_m *MockEntity) SetID(_a0 string) {
	_m.Called(_a0)
}

// SetName provides a mock function with given fields: _a0
func (_m *MockEntity) SetName(_a0 string) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewMockEntity interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEntity creates a new instance of MockEntity. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEntity(t mockConstructorTestingTNewMockEntity) *MockEntity {
	mock := &MockEntity{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
