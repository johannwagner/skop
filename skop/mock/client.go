// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	k8s "github.com/ericchiang/k8s"
	gomock "github.com/golang/mock/gomock"
	skop "github.com/thcyron/skop/skop"
	reflect "reflect"
)

// Client is a mock of Client interface
type Client struct {
	ctrl     *gomock.Controller
	recorder *ClientMockRecorder
}

// ClientMockRecorder is the mock recorder for Client
type ClientMockRecorder struct {
	mock *Client
}

// NewClient creates a new mock instance
func NewClient(ctrl *gomock.Controller) *Client {
	mock := &Client{ctrl: ctrl}
	mock.recorder = &ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Client) EXPECT() *ClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *Client) Create(ctx context.Context, res k8s.Resource, options ...k8s.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, res}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *ClientMockRecorder) Create(ctx, res interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, res}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*Client)(nil).Create), varargs...)
}

// Get mocks base method
func (m *Client) Get(ctx context.Context, namespace, name string, res k8s.Resource, options ...k8s.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, namespace, name, res}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (mr *ClientMockRecorder) Get(ctx, namespace, name, res interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, namespace, name, res}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*Client)(nil).Get), varargs...)
}

// Update mocks base method
func (m *Client) Update(ctx context.Context, res k8s.Resource, options ...k8s.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, res}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *ClientMockRecorder) Update(ctx, res interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, res}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*Client)(nil).Update), varargs...)
}

// Delete mocks base method
func (m *Client) Delete(ctx context.Context, res k8s.Resource, options ...k8s.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, res}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *ClientMockRecorder) Delete(ctx, res interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, res}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*Client)(nil).Delete), varargs...)
}

// Watch mocks base method
func (m *Client) Watch(ctx context.Context, res k8s.Resource) (skop.Watcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", ctx, res)
	ret0, _ := ret[0].(skop.Watcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *ClientMockRecorder) Watch(ctx, res interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*Client)(nil).Watch), ctx, res)
}

// Watcher is a mock of Watcher interface
type Watcher struct {
	ctrl     *gomock.Controller
	recorder *WatcherMockRecorder
}

// WatcherMockRecorder is the mock recorder for Watcher
type WatcherMockRecorder struct {
	mock *Watcher
}

// NewWatcher creates a new mock instance
func NewWatcher(ctrl *gomock.Controller) *Watcher {
	mock := &Watcher{ctrl: ctrl}
	mock.recorder = &WatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Watcher) EXPECT() *WatcherMockRecorder {
	return m.recorder
}

// Next mocks base method
func (m *Watcher) Next(res k8s.Resource) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", res)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next
func (mr *WatcherMockRecorder) Next(res interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*Watcher)(nil).Next), res)
}

// Close mocks base method
func (m *Watcher) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *WatcherMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*Watcher)(nil).Close))
}
