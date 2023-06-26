// Code generated by MockGen. DO NOT EDIT.
// Source: callback.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockExecutable is a mock of Executable interface.
type MockExecutable struct {
	ctrl     *gomock.Controller
	recorder *MockExecutableMockRecorder
}

// MockExecutableMockRecorder is the mock recorder for MockExecutable.
type MockExecutableMockRecorder struct {
	mock *MockExecutable
}

// NewMockExecutable creates a new mock instance.
func NewMockExecutable(ctrl *gomock.Controller) *MockExecutable {
	mock := &MockExecutable{ctrl: ctrl}
	mock.recorder = &MockExecutableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecutable) EXPECT() *MockExecutableMockRecorder {
	return m.recorder
}

// ExecuteAtReset mocks base method.
func (m *MockExecutable) ExecuteAtReset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExecuteAtReset")
}

// ExecuteAtReset indicates an expected call of ExecuteAtReset.
func (mr *MockExecutableMockRecorder) ExecuteAtReset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteAtReset", reflect.TypeOf((*MockExecutable)(nil).ExecuteAtReset))
}

// ExecuteAtStart mocks base method.
func (m *MockExecutable) ExecuteAtStart() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExecuteAtStart")
}

// ExecuteAtStart indicates an expected call of ExecuteAtStart.
func (mr *MockExecutableMockRecorder) ExecuteAtStart() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteAtStart", reflect.TypeOf((*MockExecutable)(nil).ExecuteAtStart))
}

// ExecuteAtStop mocks base method.
func (m *MockExecutable) ExecuteAtStop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExecuteAtStop")
}

// ExecuteAtStop indicates an expected call of ExecuteAtStop.
func (mr *MockExecutableMockRecorder) ExecuteAtStop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteAtStop", reflect.TypeOf((*MockExecutable)(nil).ExecuteAtStop))
}