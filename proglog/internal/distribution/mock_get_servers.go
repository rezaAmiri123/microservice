// Code generated by mockery v2.33.0. DO NOT EDIT.

package distribution

import mock "github.com/stretchr/testify/mock"

// MockGetServers is an autogenerated mock type for the GetServers type
type MockGetServers struct {
	mock.Mock
}

// GetServers provides a mock function with given fields:
func (_m *MockGetServers) GetServers() ([]*Server, error) {
	ret := _m.Called()

	var r0 []*Server
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*Server, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*Server); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Server)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockGetServers creates a new instance of MockGetServers. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGetServers(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGetServers {
	mock := &MockGetServers{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
