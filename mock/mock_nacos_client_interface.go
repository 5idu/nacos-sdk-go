/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by MockGen. DO NOT EDIT.
// Source: clients/nacos_client/nacos_client_interface.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	"github.com/5idu/nacos-sdk-go/common/constant"
	"github.com/5idu/nacos-sdk-go/common/http_agent"

	gomock "github.com/golang/mock/gomock"
)

// MockINacosClient is a mock of INacosClient interface
type MockINacosClient struct {
	ctrl     *gomock.Controller
	recorder *MockINacosClientMockRecorder
}

// MockINacosClientMockRecorder is the mock recorder for MockINacosClient
type MockINacosClientMockRecorder struct {
	mock *MockINacosClient
}

// NewMockINacosClient creates a new mock instance
func NewMockINacosClient(ctrl *gomock.Controller) *MockINacosClient {
	mock := &MockINacosClient{ctrl: ctrl}
	mock.recorder = &MockINacosClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockINacosClient) EXPECT() *MockINacosClientMockRecorder {
	return m.recorder
}

// SetClientConfig mocks base method
func (m *MockINacosClient) SetClientConfig(arg0 constant.ClientConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetClientConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetClientConfig indicates an expected call of SetClientConfig
func (mr *MockINacosClientMockRecorder) SetClientConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClientConfig", reflect.TypeOf((*MockINacosClient)(nil).SetClientConfig), arg0)
}

// SetServerConfig mocks base method
func (m *MockINacosClient) SetServerConfig(arg0 []constant.ServerConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetServerConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetServerConfig indicates an expected call of SetServerConfig
func (mr *MockINacosClientMockRecorder) SetServerConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetServerConfig", reflect.TypeOf((*MockINacosClient)(nil).SetServerConfig), arg0)
}

// GetClientConfig mocks base method
func (m *MockINacosClient) GetClientConfig() (constant.ClientConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientConfig")
	ret0, _ := ret[0].(constant.ClientConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientConfig indicates an expected call of GetClientConfig
func (mr *MockINacosClientMockRecorder) GetClientConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientConfig", reflect.TypeOf((*MockINacosClient)(nil).GetClientConfig))
}

// GetServerConfig mocks base method
func (m *MockINacosClient) GetServerConfig() ([]constant.ServerConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerConfig")
	ret0, _ := ret[0].([]constant.ServerConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServerConfig indicates an expected call of GetServerConfig
func (mr *MockINacosClientMockRecorder) GetServerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerConfig", reflect.TypeOf((*MockINacosClient)(nil).GetServerConfig))
}

// SetHttpAgent mocks base method
func (m *MockINacosClient) SetHttpAgent(arg0 http_agent.IHttpAgent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHttpAgent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHttpAgent indicates an expected call of SetHttpAgent
func (mr *MockINacosClientMockRecorder) SetHttpAgent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHttpAgent", reflect.TypeOf((*MockINacosClient)(nil).SetHttpAgent), arg0)
}

// GetHttpAgent mocks base method
func (m *MockINacosClient) GetHttpAgent() (http_agent.IHttpAgent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHttpAgent")
	ret0, _ := ret[0].(http_agent.IHttpAgent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHttpAgent indicates an expected call of GetHttpAgent
func (mr *MockINacosClientMockRecorder) GetHttpAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHttpAgent", reflect.TypeOf((*MockINacosClient)(nil).GetHttpAgent))
}