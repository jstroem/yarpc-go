// Code generated by thriftrw-plugin-yarpc
// @generated

// Copyright (c) 2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package secondservicetest

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	yarpc "go.uber.org/yarpc"
	secondserviceclient "go.uber.org/yarpc/internal/crossdock/thrift/gauntlet/secondserviceclient"
)

// MockClient implements a gomock-compatible mock client for service
// SecondService.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

var _ secondserviceclient.Interface = (*MockClient)(nil)

type _MockClientRecorder struct {
	mock *MockClient
}

// Build a new mock client for service SecondService.
//
// 	mockCtrl := gomock.NewController(t)
// 	client := secondservicetest.NewMockClient(mockCtrl)
//
// Use EXPECT() to set expectations on the mock.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

// EXPECT returns an object that allows you to define an expectation on the
// SecondService mock client.
func (m *MockClient) EXPECT() *_MockClientRecorder {
	return m.recorder
}

// BlahBlah responds to a BlahBlah call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().BlahBlah(gomock.Any(), ...).Return(...)
// 	... := client.BlahBlah(...)
func (m *MockClient) BlahBlah(
	ctx context.Context,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "BlahBlah", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) BlahBlah(
	ctx interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "BlahBlah", args...)
}

// SecondtestString responds to a SecondtestString call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().SecondtestString(gomock.Any(), ...).Return(...)
// 	... := client.SecondtestString(...)
func (m *MockClient) SecondtestString(
	ctx context.Context,
	_Thing *string,
	opts ...yarpc.CallOption,
) (success string, err error) {

	args := []interface{}{ctx, _Thing}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "SecondtestString", args...)
	success, _ = ret[i].(string)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) SecondtestString(
	ctx interface{},
	_Thing interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Thing}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "SecondtestString", args...)
}
