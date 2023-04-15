// Code generated by mockery v2.24.0. DO NOT EDIT.

package logger

import (
	tag "github.com/rezaAmiri123/microservice/pkg/logger/tag"
	mock "github.com/stretchr/testify/mock"
)

// MockTagLogger is an autogenerated mock type for the TagLogger type
type MockTagLogger struct {
	mock.Mock
}

// Debug provides a mock function with given fields: msg, tags
func (_m *MockTagLogger) Debug(msg string, tags ...tag.Tag) {
	_va := make([]interface{}, len(tags))
	for _i := range tags {
		_va[_i] = tags[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Error provides a mock function with given fields: msg, tags
func (_m *MockTagLogger) Error(msg string, tags ...tag.Tag) {
	_va := make([]interface{}, len(tags))
	for _i := range tags {
		_va[_i] = tags[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Fatal provides a mock function with given fields: msg, tags
func (_m *MockTagLogger) Fatal(msg string, tags ...tag.Tag) {
	_va := make([]interface{}, len(tags))
	for _i := range tags {
		_va[_i] = tags[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Info provides a mock function with given fields: msg, tags
func (_m *MockTagLogger) Info(msg string, tags ...tag.Tag) {
	_va := make([]interface{}, len(tags))
	for _i := range tags {
		_va[_i] = tags[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Warn provides a mock function with given fields: msg, tags
func (_m *MockTagLogger) Warn(msg string, tags ...tag.Tag) {
	_va := make([]interface{}, len(tags))
	for _i := range tags {
		_va[_i] = tags[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// WithTags provides a mock function with given fields: tags
func (_m *MockTagLogger) WithTags(tags ...tag.Tag) TagLogger {
	_va := make([]interface{}, len(tags))
	for _i := range tags {
		_va[_i] = tags[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 TagLogger
	if rf, ok := ret.Get(0).(func(...tag.Tag) TagLogger); ok {
		r0 = rf(tags...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(TagLogger)
		}
	}

	return r0
}

type mockConstructorTestingTNewMockTagLogger interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockTagLogger creates a new instance of MockTagLogger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTagLogger(t mockConstructorTestingTNewMockTagLogger) *MockTagLogger {
	mock := &MockTagLogger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}