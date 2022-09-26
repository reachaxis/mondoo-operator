// Code generated by MockGen. DO NOT EDIT.
// Source: ./scan_api_store.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	scan_api_store "go.mondoo.com/mondoo-operator/controllers/resource_monitor/scan_api_store"
)

// MockScanApiStore is a mock of ScanApiStore interface.
type MockScanApiStore struct {
	ctrl     *gomock.Controller
	recorder *MockScanApiStoreMockRecorder
}

// MockScanApiStoreMockRecorder is the mock recorder for MockScanApiStore.
type MockScanApiStoreMockRecorder struct {
	mock *MockScanApiStore
}

// NewMockScanApiStore creates a new mock instance.
func NewMockScanApiStore(ctrl *gomock.Controller) *MockScanApiStore {
	mock := &MockScanApiStore{ctrl: ctrl}
	mock.recorder = &MockScanApiStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScanApiStore) EXPECT() *MockScanApiStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockScanApiStore) Add(url, token, integrationMrn string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", url, token, integrationMrn)
}

// Add indicates an expected call of Add.
func (mr *MockScanApiStoreMockRecorder) Add(url, token, integrationMrn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockScanApiStore)(nil).Add), url, token, integrationMrn)
}

// Delete mocks base method.
func (m *MockScanApiStore) Delete(url string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", url)
}

// Delete indicates an expected call of Delete.
func (mr *MockScanApiStoreMockRecorder) Delete(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockScanApiStore)(nil).Delete), url)
}

// GetAll mocks base method.
func (m *MockScanApiStore) GetAll() []scan_api_store.ClientConfiguration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]scan_api_store.ClientConfiguration)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockScanApiStoreMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockScanApiStore)(nil).GetAll))
}

// Start mocks base method.
func (m *MockScanApiStore) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockScanApiStoreMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockScanApiStore)(nil).Start))
}