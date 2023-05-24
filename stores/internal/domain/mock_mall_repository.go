// Code generated by mockery v2.27.1. DO NOT EDIT.

package domain

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockMallRepository is an autogenerated mock type for the MallRepository type
type MockMallRepository struct {
	mock.Mock
}

// AddStore provides a mock function with given fields: ctx, storeID, name, location
func (_m *MockMallRepository) AddStore(ctx context.Context, storeID string, name string, location string) error {
	ret := _m.Called(ctx, storeID, name, location)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, storeID, name, location)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// All provides a mock function with given fields: ctx
func (_m *MockMallRepository) All(ctx context.Context) ([]*MallStore, error) {
	ret := _m.Called(ctx)

	var r0 []*MallStore
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*MallStore, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*MallStore); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*MallStore)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllParticipating provides a mock function with given fields: ctx
func (_m *MockMallRepository) AllParticipating(ctx context.Context) ([]*MallStore, error) {
	ret := _m.Called(ctx)

	var r0 []*MallStore
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*MallStore, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*MallStore); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*MallStore)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx, storeID
func (_m *MockMallRepository) Find(ctx context.Context, storeID string) (*MallStore, error) {
	ret := _m.Called(ctx, storeID)

	var r0 *MallStore
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*MallStore, error)); ok {
		return rf(ctx, storeID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *MallStore); ok {
		r0 = rf(ctx, storeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*MallStore)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, storeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RenameStore provides a mock function with given fields: ctx, storeID, name
func (_m *MockMallRepository) RenameStore(ctx context.Context, storeID string, name string) error {
	ret := _m.Called(ctx, storeID, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, storeID, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetStoreParticipation provides a mock function with given fields: ctx, storeID, participating
func (_m *MockMallRepository) SetStoreParticipation(ctx context.Context, storeID string, participating bool) error {
	ret := _m.Called(ctx, storeID, participating)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) error); ok {
		r0 = rf(ctx, storeID, participating)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockMallRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockMallRepository creates a new instance of MockMallRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockMallRepository(t mockConstructorTestingTNewMockMallRepository) *MockMallRepository {
	mock := &MockMallRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
