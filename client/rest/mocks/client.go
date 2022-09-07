// Copyright (C) 2022 Specter Ops, Inc.
//
// This file is part of AzureHound.
//
// AzureHound is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// AzureHound is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.


// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRestClient is a mock of RestClient interface.
type MockRestClient struct {
	ctrl     *gomock.Controller
	recorder *MockRestClientMockRecorder
}

// MockRestClientMockRecorder is the mock recorder for MockRestClient.
type MockRestClientMockRecorder struct {
	mock *MockRestClient
}

// NewMockRestClient creates a new mock instance.
func NewMockRestClient(ctrl *gomock.Controller) *MockRestClient {
	mock := &MockRestClient{ctrl: ctrl}
	mock.recorder = &MockRestClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRestClient) EXPECT() *MockRestClientMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockRestClient) Authenticate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockRestClientMockRecorder) Authenticate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockRestClient)(nil).Authenticate))
}

// Delete mocks base method.
func (m *MockRestClient) Delete(arg0 context.Context, arg1 string, arg2 interface{}, arg3, arg4 map[string]string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockRestClientMockRecorder) Delete(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRestClient)(nil).Delete), arg0, arg1, arg2, arg3, arg4)
}

// Get mocks base method.
func (m *MockRestClient) Get(arg0 context.Context, arg1 string, arg2, arg3 map[string]string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRestClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRestClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// Patch mocks base method.
func (m *MockRestClient) Patch(arg0 context.Context, arg1 string, arg2 interface{}, arg3, arg4 map[string]string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Patch", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockRestClientMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockRestClient)(nil).Patch), arg0, arg1, arg2, arg3, arg4)
}

// Post mocks base method.
func (m *MockRestClient) Post(arg0 context.Context, arg1 string, arg2 interface{}, arg3, arg4 map[string]string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockRestClientMockRecorder) Post(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockRestClient)(nil).Post), arg0, arg1, arg2, arg3, arg4)
}

// Put mocks base method.
func (m *MockRestClient) Put(arg0 context.Context, arg1 string, arg2 interface{}, arg3, arg4 map[string]string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockRestClientMockRecorder) Put(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockRestClient)(nil).Put), arg0, arg1, arg2, arg3, arg4)
}

// Send mocks base method.
func (m *MockRestClient) Send(arg0 *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockRestClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockRestClient)(nil).Send), arg0)
}
