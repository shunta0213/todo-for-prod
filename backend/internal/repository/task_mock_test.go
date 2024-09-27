// Code generated by MockGen. DO NOT EDIT.
// Source: task.go
//
// Generated by this command:
//
//	mockgen -source=task.go -destination=task_mock_test.go -package=repository
//

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockTask is a mock of Task interface.
type MockTask struct {
	ctrl     *gomock.Controller
	recorder *MockTaskMockRecorder
}

// MockTaskMockRecorder is the mock recorder for MockTask.
type MockTaskMockRecorder struct {
	mock *MockTask
}

// NewMockTask creates a new mock instance.
func NewMockTask(ctrl *gomock.Controller) *MockTask {
	mock := &MockTask{ctrl: ctrl}
	mock.recorder = &MockTaskMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTask) EXPECT() *MockTaskMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTask) Create() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create")
}

// Create indicates an expected call of Create.
func (mr *MockTaskMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTask)(nil).Create))
}

// Get mocks base method.
func (m *MockTask) Get() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Get")
}

// Get indicates an expected call of Get.
func (mr *MockTaskMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTask)(nil).Get))
}

// GetAll mocks base method.
func (m *MockTask) GetAll() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetAll")
}

// GetAll indicates an expected call of GetAll.
func (mr *MockTaskMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTask)(nil).GetAll))
}

// Update mocks base method.
func (m *MockTask) Update() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update")
}

// Update indicates an expected call of Update.
func (mr *MockTaskMockRecorder) Update() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTask)(nil).Update))
}
