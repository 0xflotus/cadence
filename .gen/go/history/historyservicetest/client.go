// The MIT License (MIT)
// 
// Copyright (c) 2017 Uber Technologies, Inc.
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by thriftrw-plugin-yarpc
// @generated

package historyservicetest

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	history "github.com/uber/cadence/.gen/go/history"
	historyserviceclient "github.com/uber/cadence/.gen/go/history/historyserviceclient"
	replicator "github.com/uber/cadence/.gen/go/replicator"
	shared "github.com/uber/cadence/.gen/go/shared"
	yarpc "go.uber.org/yarpc"
)

// MockClient implements a gomock-compatible mock client for service
// HistoryService.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

var _ historyserviceclient.Interface = (*MockClient)(nil)

type _MockClientRecorder struct {
	mock *MockClient
}

// Build a new mock client for service HistoryService.
//
// 	mockCtrl := gomock.NewController(t)
// 	client := historyservicetest.NewMockClient(mockCtrl)
//
// Use EXPECT() to set expectations on the mock.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

// EXPECT returns an object that allows you to define an expectation on the
// HistoryService mock client.
func (m *MockClient) EXPECT() *_MockClientRecorder {
	return m.recorder
}

// CloseShard responds to a CloseShard call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().CloseShard(gomock.Any(), ...).Return(...)
// 	... := client.CloseShard(...)
func (m *MockClient) CloseShard(
	ctx context.Context,
	_Request *shared.CloseShardRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _Request}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "CloseShard", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) CloseShard(
	ctx interface{},
	_Request interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Request}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "CloseShard", args...)
}

// DescribeHistoryHost responds to a DescribeHistoryHost call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().DescribeHistoryHost(gomock.Any(), ...).Return(...)
// 	... := client.DescribeHistoryHost(...)
func (m *MockClient) DescribeHistoryHost(
	ctx context.Context,
	_Request *shared.DescribeHistoryHostRequest,
	opts ...yarpc.CallOption,
) (success *shared.DescribeHistoryHostResponse, err error) {

	args := []interface{}{ctx, _Request}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "DescribeHistoryHost", args...)
	success, _ = ret[i].(*shared.DescribeHistoryHostResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) DescribeHistoryHost(
	ctx interface{},
	_Request interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Request}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "DescribeHistoryHost", args...)
}

// DescribeMutableState responds to a DescribeMutableState call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().DescribeMutableState(gomock.Any(), ...).Return(...)
// 	... := client.DescribeMutableState(...)
func (m *MockClient) DescribeMutableState(
	ctx context.Context,
	_Request *history.DescribeMutableStateRequest,
	opts ...yarpc.CallOption,
) (success *history.DescribeMutableStateResponse, err error) {

	args := []interface{}{ctx, _Request}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "DescribeMutableState", args...)
	success, _ = ret[i].(*history.DescribeMutableStateResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) DescribeMutableState(
	ctx interface{},
	_Request interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Request}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "DescribeMutableState", args...)
}

// DescribeWorkflowExecution responds to a DescribeWorkflowExecution call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().DescribeWorkflowExecution(gomock.Any(), ...).Return(...)
// 	... := client.DescribeWorkflowExecution(...)
func (m *MockClient) DescribeWorkflowExecution(
	ctx context.Context,
	_DescribeRequest *history.DescribeWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (success *shared.DescribeWorkflowExecutionResponse, err error) {

	args := []interface{}{ctx, _DescribeRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "DescribeWorkflowExecution", args...)
	success, _ = ret[i].(*shared.DescribeWorkflowExecutionResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) DescribeWorkflowExecution(
	ctx interface{},
	_DescribeRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _DescribeRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "DescribeWorkflowExecution", args...)
}

// GetMutableState responds to a GetMutableState call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().GetMutableState(gomock.Any(), ...).Return(...)
// 	... := client.GetMutableState(...)
func (m *MockClient) GetMutableState(
	ctx context.Context,
	_GetRequest *history.GetMutableStateRequest,
	opts ...yarpc.CallOption,
) (success *history.GetMutableStateResponse, err error) {

	args := []interface{}{ctx, _GetRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "GetMutableState", args...)
	success, _ = ret[i].(*history.GetMutableStateResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) GetMutableState(
	ctx interface{},
	_GetRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _GetRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "GetMutableState", args...)
}

// GetReplicationMessages responds to a GetReplicationMessages call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().GetReplicationMessages(gomock.Any(), ...).Return(...)
// 	... := client.GetReplicationMessages(...)
func (m *MockClient) GetReplicationMessages(
	ctx context.Context,
	_Request *replicator.GetReplicationMessagesRequest,
	opts ...yarpc.CallOption,
) (success *replicator.GetReplicationMessagesResponse, err error) {

	args := []interface{}{ctx, _Request}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "GetReplicationMessages", args...)
	success, _ = ret[i].(*replicator.GetReplicationMessagesResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) GetReplicationMessages(
	ctx interface{},
	_Request interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Request}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "GetReplicationMessages", args...)
}

// QueryWorkflow responds to a QueryWorkflow call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().QueryWorkflow(gomock.Any(), ...).Return(...)
// 	... := client.QueryWorkflow(...)
func (m *MockClient) QueryWorkflow(
	ctx context.Context,
	_QueryRequest *history.QueryWorkflowRequest,
	opts ...yarpc.CallOption,
) (success *history.QueryWorkflowResponse, err error) {

	args := []interface{}{ctx, _QueryRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "QueryWorkflow", args...)
	success, _ = ret[i].(*history.QueryWorkflowResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) QueryWorkflow(
	ctx interface{},
	_QueryRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _QueryRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "QueryWorkflow", args...)
}

// RecordActivityTaskHeartbeat responds to a RecordActivityTaskHeartbeat call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().RecordActivityTaskHeartbeat(gomock.Any(), ...).Return(...)
// 	... := client.RecordActivityTaskHeartbeat(...)
func (m *MockClient) RecordActivityTaskHeartbeat(
	ctx context.Context,
	_HeartbeatRequest *history.RecordActivityTaskHeartbeatRequest,
	opts ...yarpc.CallOption,
) (success *shared.RecordActivityTaskHeartbeatResponse, err error) {

	args := []interface{}{ctx, _HeartbeatRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "RecordActivityTaskHeartbeat", args...)
	success, _ = ret[i].(*shared.RecordActivityTaskHeartbeatResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) RecordActivityTaskHeartbeat(
	ctx interface{},
	_HeartbeatRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _HeartbeatRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "RecordActivityTaskHeartbeat", args...)
}

// RecordActivityTaskStarted responds to a RecordActivityTaskStarted call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().RecordActivityTaskStarted(gomock.Any(), ...).Return(...)
// 	... := client.RecordActivityTaskStarted(...)
func (m *MockClient) RecordActivityTaskStarted(
	ctx context.Context,
	_AddRequest *history.RecordActivityTaskStartedRequest,
	opts ...yarpc.CallOption,
) (success *history.RecordActivityTaskStartedResponse, err error) {

	args := []interface{}{ctx, _AddRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "RecordActivityTaskStarted", args...)
	success, _ = ret[i].(*history.RecordActivityTaskStartedResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) RecordActivityTaskStarted(
	ctx interface{},
	_AddRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _AddRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "RecordActivityTaskStarted", args...)
}

// RecordChildExecutionCompleted responds to a RecordChildExecutionCompleted call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().RecordChildExecutionCompleted(gomock.Any(), ...).Return(...)
// 	... := client.RecordChildExecutionCompleted(...)
func (m *MockClient) RecordChildExecutionCompleted(
	ctx context.Context,
	_CompletionRequest *history.RecordChildExecutionCompletedRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _CompletionRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "RecordChildExecutionCompleted", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) RecordChildExecutionCompleted(
	ctx interface{},
	_CompletionRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _CompletionRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "RecordChildExecutionCompleted", args...)
}

// RecordDecisionTaskStarted responds to a RecordDecisionTaskStarted call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().RecordDecisionTaskStarted(gomock.Any(), ...).Return(...)
// 	... := client.RecordDecisionTaskStarted(...)
func (m *MockClient) RecordDecisionTaskStarted(
	ctx context.Context,
	_AddRequest *history.RecordDecisionTaskStartedRequest,
	opts ...yarpc.CallOption,
) (success *history.RecordDecisionTaskStartedResponse, err error) {

	args := []interface{}{ctx, _AddRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "RecordDecisionTaskStarted", args...)
	success, _ = ret[i].(*history.RecordDecisionTaskStartedResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) RecordDecisionTaskStarted(
	ctx interface{},
	_AddRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _AddRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "RecordDecisionTaskStarted", args...)
}

// RemoveSignalMutableState responds to a RemoveSignalMutableState call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().RemoveSignalMutableState(gomock.Any(), ...).Return(...)
// 	... := client.RemoveSignalMutableState(...)
func (m *MockClient) RemoveSignalMutableState(
	ctx context.Context,
	_RemoveRequest *history.RemoveSignalMutableStateRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _RemoveRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "RemoveSignalMutableState", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) RemoveSignalMutableState(
	ctx interface{},
	_RemoveRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _RemoveRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "RemoveSignalMutableState", args...)
}

// RemoveTask responds to a RemoveTask call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().RemoveTask(gomock.Any(), ...).Return(...)
// 	... := client.RemoveTask(...)
func (m *MockClient) RemoveTask(
	ctx context.Context,
	_Request *shared.RemoveTaskRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _Request}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "RemoveTask", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) RemoveTask(
	ctx interface{},
	_Request interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Request}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "RemoveTask", args...)
}

// ReplicateEvents responds to a ReplicateEvents call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().ReplicateEvents(gomock.Any(), ...).Return(...)
// 	... := client.ReplicateEvents(...)
func (m *MockClient) ReplicateEvents(
	ctx context.Context,
	_ReplicateRequest *history.ReplicateEventsRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _ReplicateRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "ReplicateEvents", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) ReplicateEvents(
	ctx interface{},
	_ReplicateRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _ReplicateRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "ReplicateEvents", args...)
}

// ReplicateRawEvents responds to a ReplicateRawEvents call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().ReplicateRawEvents(gomock.Any(), ...).Return(...)
// 	... := client.ReplicateRawEvents(...)
func (m *MockClient) ReplicateRawEvents(
	ctx context.Context,
	_ReplicateRequest *history.ReplicateRawEventsRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _ReplicateRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "ReplicateRawEvents", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) ReplicateRawEvents(
	ctx interface{},
	_ReplicateRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _ReplicateRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "ReplicateRawEvents", args...)
}

// RequestCancelWorkflowExecution responds to a RequestCancelWorkflowExecution call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().RequestCancelWorkflowExecution(gomock.Any(), ...).Return(...)
// 	... := client.RequestCancelWorkflowExecution(...)
func (m *MockClient) RequestCancelWorkflowExecution(
	ctx context.Context,
	_CancelRequest *history.RequestCancelWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _CancelRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "RequestCancelWorkflowExecution", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) RequestCancelWorkflowExecution(
	ctx interface{},
	_CancelRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _CancelRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "RequestCancelWorkflowExecution", args...)
}

// ResetStickyTaskList responds to a ResetStickyTaskList call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().ResetStickyTaskList(gomock.Any(), ...).Return(...)
// 	... := client.ResetStickyTaskList(...)
func (m *MockClient) ResetStickyTaskList(
	ctx context.Context,
	_ResetRequest *history.ResetStickyTaskListRequest,
	opts ...yarpc.CallOption,
) (success *history.ResetStickyTaskListResponse, err error) {

	args := []interface{}{ctx, _ResetRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "ResetStickyTaskList", args...)
	success, _ = ret[i].(*history.ResetStickyTaskListResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) ResetStickyTaskList(
	ctx interface{},
	_ResetRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _ResetRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "ResetStickyTaskList", args...)
}

// ResetWorkflowExecution responds to a ResetWorkflowExecution call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().ResetWorkflowExecution(gomock.Any(), ...).Return(...)
// 	... := client.ResetWorkflowExecution(...)
func (m *MockClient) ResetWorkflowExecution(
	ctx context.Context,
	_ResetRequest *history.ResetWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (success *shared.ResetWorkflowExecutionResponse, err error) {

	args := []interface{}{ctx, _ResetRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "ResetWorkflowExecution", args...)
	success, _ = ret[i].(*shared.ResetWorkflowExecutionResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) ResetWorkflowExecution(
	ctx interface{},
	_ResetRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _ResetRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "ResetWorkflowExecution", args...)
}

// RespondActivityTaskCanceled responds to a RespondActivityTaskCanceled call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().RespondActivityTaskCanceled(gomock.Any(), ...).Return(...)
// 	... := client.RespondActivityTaskCanceled(...)
func (m *MockClient) RespondActivityTaskCanceled(
	ctx context.Context,
	_CanceledRequest *history.RespondActivityTaskCanceledRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _CanceledRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "RespondActivityTaskCanceled", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) RespondActivityTaskCanceled(
	ctx interface{},
	_CanceledRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _CanceledRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "RespondActivityTaskCanceled", args...)
}

// RespondActivityTaskCompleted responds to a RespondActivityTaskCompleted call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().RespondActivityTaskCompleted(gomock.Any(), ...).Return(...)
// 	... := client.RespondActivityTaskCompleted(...)
func (m *MockClient) RespondActivityTaskCompleted(
	ctx context.Context,
	_CompleteRequest *history.RespondActivityTaskCompletedRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _CompleteRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "RespondActivityTaskCompleted", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) RespondActivityTaskCompleted(
	ctx interface{},
	_CompleteRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _CompleteRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "RespondActivityTaskCompleted", args...)
}

// RespondActivityTaskFailed responds to a RespondActivityTaskFailed call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().RespondActivityTaskFailed(gomock.Any(), ...).Return(...)
// 	... := client.RespondActivityTaskFailed(...)
func (m *MockClient) RespondActivityTaskFailed(
	ctx context.Context,
	_FailRequest *history.RespondActivityTaskFailedRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _FailRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "RespondActivityTaskFailed", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) RespondActivityTaskFailed(
	ctx interface{},
	_FailRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _FailRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "RespondActivityTaskFailed", args...)
}

// RespondDecisionTaskCompleted responds to a RespondDecisionTaskCompleted call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().RespondDecisionTaskCompleted(gomock.Any(), ...).Return(...)
// 	... := client.RespondDecisionTaskCompleted(...)
func (m *MockClient) RespondDecisionTaskCompleted(
	ctx context.Context,
	_CompleteRequest *history.RespondDecisionTaskCompletedRequest,
	opts ...yarpc.CallOption,
) (success *history.RespondDecisionTaskCompletedResponse, err error) {

	args := []interface{}{ctx, _CompleteRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "RespondDecisionTaskCompleted", args...)
	success, _ = ret[i].(*history.RespondDecisionTaskCompletedResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) RespondDecisionTaskCompleted(
	ctx interface{},
	_CompleteRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _CompleteRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "RespondDecisionTaskCompleted", args...)
}

// RespondDecisionTaskFailed responds to a RespondDecisionTaskFailed call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().RespondDecisionTaskFailed(gomock.Any(), ...).Return(...)
// 	... := client.RespondDecisionTaskFailed(...)
func (m *MockClient) RespondDecisionTaskFailed(
	ctx context.Context,
	_FailedRequest *history.RespondDecisionTaskFailedRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _FailedRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "RespondDecisionTaskFailed", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) RespondDecisionTaskFailed(
	ctx interface{},
	_FailedRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _FailedRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "RespondDecisionTaskFailed", args...)
}

// ScheduleDecisionTask responds to a ScheduleDecisionTask call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().ScheduleDecisionTask(gomock.Any(), ...).Return(...)
// 	... := client.ScheduleDecisionTask(...)
func (m *MockClient) ScheduleDecisionTask(
	ctx context.Context,
	_ScheduleRequest *history.ScheduleDecisionTaskRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _ScheduleRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "ScheduleDecisionTask", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) ScheduleDecisionTask(
	ctx interface{},
	_ScheduleRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _ScheduleRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "ScheduleDecisionTask", args...)
}

// SignalWithStartWorkflowExecution responds to a SignalWithStartWorkflowExecution call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().SignalWithStartWorkflowExecution(gomock.Any(), ...).Return(...)
// 	... := client.SignalWithStartWorkflowExecution(...)
func (m *MockClient) SignalWithStartWorkflowExecution(
	ctx context.Context,
	_SignalWithStartRequest *history.SignalWithStartWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (success *shared.StartWorkflowExecutionResponse, err error) {

	args := []interface{}{ctx, _SignalWithStartRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "SignalWithStartWorkflowExecution", args...)
	success, _ = ret[i].(*shared.StartWorkflowExecutionResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) SignalWithStartWorkflowExecution(
	ctx interface{},
	_SignalWithStartRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _SignalWithStartRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "SignalWithStartWorkflowExecution", args...)
}

// SignalWorkflowExecution responds to a SignalWorkflowExecution call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().SignalWorkflowExecution(gomock.Any(), ...).Return(...)
// 	... := client.SignalWorkflowExecution(...)
func (m *MockClient) SignalWorkflowExecution(
	ctx context.Context,
	_SignalRequest *history.SignalWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _SignalRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "SignalWorkflowExecution", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) SignalWorkflowExecution(
	ctx interface{},
	_SignalRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _SignalRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "SignalWorkflowExecution", args...)
}

// StartWorkflowExecution responds to a StartWorkflowExecution call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().StartWorkflowExecution(gomock.Any(), ...).Return(...)
// 	... := client.StartWorkflowExecution(...)
func (m *MockClient) StartWorkflowExecution(
	ctx context.Context,
	_StartRequest *history.StartWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (success *shared.StartWorkflowExecutionResponse, err error) {

	args := []interface{}{ctx, _StartRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "StartWorkflowExecution", args...)
	success, _ = ret[i].(*shared.StartWorkflowExecutionResponse)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) StartWorkflowExecution(
	ctx interface{},
	_StartRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _StartRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "StartWorkflowExecution", args...)
}

// SyncActivity responds to a SyncActivity call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().SyncActivity(gomock.Any(), ...).Return(...)
// 	... := client.SyncActivity(...)
func (m *MockClient) SyncActivity(
	ctx context.Context,
	_SyncActivityRequest *history.SyncActivityRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _SyncActivityRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "SyncActivity", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) SyncActivity(
	ctx interface{},
	_SyncActivityRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _SyncActivityRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "SyncActivity", args...)
}

// SyncShardStatus responds to a SyncShardStatus call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().SyncShardStatus(gomock.Any(), ...).Return(...)
// 	... := client.SyncShardStatus(...)
func (m *MockClient) SyncShardStatus(
	ctx context.Context,
	_SyncShardStatusRequest *history.SyncShardStatusRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _SyncShardStatusRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "SyncShardStatus", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) SyncShardStatus(
	ctx interface{},
	_SyncShardStatusRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _SyncShardStatusRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "SyncShardStatus", args...)
}

// TerminateWorkflowExecution responds to a TerminateWorkflowExecution call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
// 	client.EXPECT().TerminateWorkflowExecution(gomock.Any(), ...).Return(...)
// 	... := client.TerminateWorkflowExecution(...)
func (m *MockClient) TerminateWorkflowExecution(
	ctx context.Context,
	_TerminateRequest *history.TerminateWorkflowExecutionRequest,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _TerminateRequest}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TerminateWorkflowExecution", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TerminateWorkflowExecution(
	ctx interface{},
	_TerminateRequest interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _TerminateRequest}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TerminateWorkflowExecution", args...)
}
