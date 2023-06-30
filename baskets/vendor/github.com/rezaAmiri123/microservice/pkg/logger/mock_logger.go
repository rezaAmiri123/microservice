// Code generated by mockery v2.27.1. DO NOT EDIT.

package logger

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockLogger is an autogenerated mock type for the Logger type
type MockLogger struct {
	mock.Mock
}

// DPanic provides a mock function with given fields: args
func (_m *MockLogger) DPanic(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// DPanicf provides a mock function with given fields: template, args
func (_m *MockLogger) DPanicf(template string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, template)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Debug provides a mock function with given fields: args
func (_m *MockLogger) Debug(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Debugf provides a mock function with given fields: template, args
func (_m *MockLogger) Debugf(template string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, template)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Err provides a mock function with given fields: msg, err
func (_m *MockLogger) Err(msg string, err error) {
	_m.Called(msg, err)
}

// Error provides a mock function with given fields: args
func (_m *MockLogger) Error(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Errorf provides a mock function with given fields: template, args
func (_m *MockLogger) Errorf(template string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, template)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Fatal provides a mock function with given fields: args
func (_m *MockLogger) Fatal(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Fatalf provides a mock function with given fields: template, args
func (_m *MockLogger) Fatalf(template string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, template)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// GrpcClientInterceptorLogger provides a mock function with given fields: method, req, reply, _a3, metaData, err
func (_m *MockLogger) GrpcClientInterceptorLogger(method string, req interface{}, reply interface{}, _a3 time.Duration, metaData map[string][]string, err error) {
	_m.Called(method, req, reply, _a3, metaData, err)
}

// GrpcMiddlewareAccessLogger provides a mock function with given fields: method, _a1, metaData, err
func (_m *MockLogger) GrpcMiddlewareAccessLogger(method string, _a1 time.Duration, metaData map[string][]string, err error) {
	_m.Called(method, _a1, metaData, err)
}

// Info provides a mock function with given fields: args
func (_m *MockLogger) Info(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Infof provides a mock function with given fields: template, args
func (_m *MockLogger) Infof(template string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, template)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// InitLogger provides a mock function with given fields:
func (_m *MockLogger) InitLogger() {
	_m.Called()
}

// KafkaLogCommittedMessage provides a mock function with given fields: topic, partition, offset
func (_m *MockLogger) KafkaLogCommittedMessage(topic string, partition int, offset int64) {
	_m.Called(topic, partition, offset)
}

// KafkaProcessMessage provides a mock function with given fields: topic, partition, message, workerID, offset, _a5
func (_m *MockLogger) KafkaProcessMessage(topic string, partition int, message string, workerID int, offset int64, _a5 time.Time) {
	_m.Called(topic, partition, message, workerID, offset, _a5)
}

// Printf provides a mock function with given fields: template, args
func (_m *MockLogger) Printf(template string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, template)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Sync provides a mock function with given fields:
func (_m *MockLogger) Sync() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Warn provides a mock function with given fields: args
func (_m *MockLogger) Warn(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// WarnMsg provides a mock function with given fields: msg, err
func (_m *MockLogger) WarnMsg(msg string, err error) {
	_m.Called(msg, err)
}

// Warnf provides a mock function with given fields: template, args
func (_m *MockLogger) Warnf(template string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, template)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// WithName provides a mock function with given fields: name
func (_m *MockLogger) WithName(name string) {
	_m.Called(name)
}

type mockConstructorTestingTNewMockLogger interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockLogger creates a new instance of MockLogger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockLogger(t mockConstructorTestingTNewMockLogger) *MockLogger {
	mock := &MockLogger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
