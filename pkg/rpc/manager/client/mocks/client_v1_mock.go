// Code generated by MockGen. DO NOT EDIT.
// Source: client_v1.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	manager "d7y.io/api/v2/pkg/apis/manager/v1"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockV1 is a mock of V1 interface.
type MockV1 struct {
	ctrl     *gomock.Controller
	recorder *MockV1MockRecorder
}

// MockV1MockRecorder is the mock recorder for MockV1.
type MockV1MockRecorder struct {
	mock *MockV1
}

// NewMockV1 creates a new mock instance.
func NewMockV1(ctrl *gomock.Controller) *MockV1 {
	mock := &MockV1{ctrl: ctrl}
	mock.recorder = &MockV1MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockV1) EXPECT() *MockV1MockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockV1) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockV1MockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockV1)(nil).Close))
}

// CreateModel mocks base method.
func (m *MockV1) CreateModel(arg0 context.Context, arg1 *manager.CreateModelRequest, arg2 ...grpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateModel", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateModel indicates an expected call of CreateModel.
func (mr *MockV1MockRecorder) CreateModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateModel", reflect.TypeOf((*MockV1)(nil).CreateModel), varargs...)
}

// GetObjectStorage mocks base method.
func (m *MockV1) GetObjectStorage(arg0 context.Context, arg1 *manager.GetObjectStorageRequest, arg2 ...grpc.CallOption) (*manager.ObjectStorage, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectStorage", varargs...)
	ret0, _ := ret[0].(*manager.ObjectStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectStorage indicates an expected call of GetObjectStorage.
func (mr *MockV1MockRecorder) GetObjectStorage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectStorage", reflect.TypeOf((*MockV1)(nil).GetObjectStorage), varargs...)
}

// GetScheduler mocks base method.
func (m *MockV1) GetScheduler(arg0 context.Context, arg1 *manager.GetSchedulerRequest, arg2 ...grpc.CallOption) (*manager.Scheduler, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetScheduler", varargs...)
	ret0, _ := ret[0].(*manager.Scheduler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScheduler indicates an expected call of GetScheduler.
func (mr *MockV1MockRecorder) GetScheduler(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScheduler", reflect.TypeOf((*MockV1)(nil).GetScheduler), varargs...)
}

// KeepAlive mocks base method.
func (m *MockV1) KeepAlive(arg0 time.Duration, arg1 *manager.KeepAliveRequest, arg2 <-chan struct{}, arg3 ...grpc.CallOption) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "KeepAlive", varargs...)
}

// KeepAlive indicates an expected call of KeepAlive.
func (mr *MockV1MockRecorder) KeepAlive(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeepAlive", reflect.TypeOf((*MockV1)(nil).KeepAlive), varargs...)
}

// ListApplications mocks base method.
func (m *MockV1) ListApplications(arg0 context.Context, arg1 *manager.ListApplicationsRequest, arg2 ...grpc.CallOption) (*manager.ListApplicationsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListApplications", varargs...)
	ret0, _ := ret[0].(*manager.ListApplicationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApplications indicates an expected call of ListApplications.
func (mr *MockV1MockRecorder) ListApplications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplications", reflect.TypeOf((*MockV1)(nil).ListApplications), varargs...)
}

// ListBuckets mocks base method.
func (m *MockV1) ListBuckets(arg0 context.Context, arg1 *manager.ListBucketsRequest, arg2 ...grpc.CallOption) (*manager.ListBucketsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBuckets", varargs...)
	ret0, _ := ret[0].(*manager.ListBucketsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBuckets indicates an expected call of ListBuckets.
func (mr *MockV1MockRecorder) ListBuckets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuckets", reflect.TypeOf((*MockV1)(nil).ListBuckets), varargs...)
}

// ListSchedulers mocks base method.
func (m *MockV1) ListSchedulers(arg0 context.Context, arg1 *manager.ListSchedulersRequest, arg2 ...grpc.CallOption) (*manager.ListSchedulersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSchedulers", varargs...)
	ret0, _ := ret[0].(*manager.ListSchedulersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchedulers indicates an expected call of ListSchedulers.
func (mr *MockV1MockRecorder) ListSchedulers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedulers", reflect.TypeOf((*MockV1)(nil).ListSchedulers), varargs...)
}

// UpdateScheduler mocks base method.
func (m *MockV1) UpdateScheduler(arg0 context.Context, arg1 *manager.UpdateSchedulerRequest, arg2 ...grpc.CallOption) (*manager.Scheduler, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateScheduler", varargs...)
	ret0, _ := ret[0].(*manager.Scheduler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScheduler indicates an expected call of UpdateScheduler.
func (mr *MockV1MockRecorder) UpdateScheduler(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScheduler", reflect.TypeOf((*MockV1)(nil).UpdateScheduler), varargs...)
}

// UpdateSeedPeer mocks base method.
func (m *MockV1) UpdateSeedPeer(arg0 context.Context, arg1 *manager.UpdateSeedPeerRequest, arg2 ...grpc.CallOption) (*manager.SeedPeer, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSeedPeer", varargs...)
	ret0, _ := ret[0].(*manager.SeedPeer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSeedPeer indicates an expected call of UpdateSeedPeer.
func (mr *MockV1MockRecorder) UpdateSeedPeer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSeedPeer", reflect.TypeOf((*MockV1)(nil).UpdateSeedPeer), varargs...)
}
