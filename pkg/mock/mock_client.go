// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/meroxa/meroxa.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	http "net/http"
	url "net/url"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	meroxa "github.com/meroxa/meroxa-go/pkg/meroxa"
)

// Mockrequester is a mock of requester interface.
type Mockrequester struct {
	ctrl     *gomock.Controller
	recorder *MockrequesterMockRecorder
}

// MockrequesterMockRecorder is the mock recorder for Mockrequester.
type MockrequesterMockRecorder struct {
	mock *Mockrequester
}

// NewMockrequester creates a new mock instance.
func NewMockrequester(ctrl *gomock.Controller) *Mockrequester {
	mock := &Mockrequester{ctrl: ctrl}
	mock.recorder = &MockrequesterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockrequester) EXPECT() *MockrequesterMockRecorder {
	return m.recorder
}

// AddHeader mocks base method.
func (m *Mockrequester) AddHeader(key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddHeader", key, value)
}

// AddHeader indicates an expected call of AddHeader.
func (mr *MockrequesterMockRecorder) AddHeader(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHeader", reflect.TypeOf((*Mockrequester)(nil).AddHeader), key, value)
}

// MakeRequest mocks base method.
func (m *Mockrequester) MakeRequest(ctx context.Context, method, path string, body interface{}, params url.Values, headers http.Header) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeRequest", ctx, method, path, body, params, headers)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeRequest indicates an expected call of MakeRequest.
func (mr *MockrequesterMockRecorder) MakeRequest(ctx, method, path, body, params, headers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRequest", reflect.TypeOf((*Mockrequester)(nil).MakeRequest), ctx, method, path, body, params, headers)
}

// Mockaccount is a mock of account interface.
type Mockaccount struct {
	ctrl     *gomock.Controller
	recorder *MockaccountMockRecorder
}

// MockaccountMockRecorder is the mock recorder for Mockaccount.
type MockaccountMockRecorder struct {
	mock *Mockaccount
}

// NewMockaccount creates a new mock instance.
func NewMockaccount(ctrl *gomock.Controller) *Mockaccount {
	mock := &Mockaccount{ctrl: ctrl}
	mock.recorder = &MockaccountMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockaccount) EXPECT() *MockaccountMockRecorder {
	return m.recorder
}

// ListAccounts mocks base method.
func (m *Mockaccount) ListAccounts(ctx context.Context) ([]*meroxa.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", ctx)
	ret0, _ := ret[0].([]*meroxa.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockaccountMockRecorder) ListAccounts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*Mockaccount)(nil).ListAccounts), ctx)
}

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AddHeader mocks base method.
func (m *MockClient) AddHeader(key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddHeader", key, value)
}

// AddHeader indicates an expected call of AddHeader.
func (mr *MockClientMockRecorder) AddHeader(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHeader", reflect.TypeOf((*MockClient)(nil).AddHeader), key, value)
}

// CreateApplication mocks base method.
func (m *MockClient) CreateApplication(ctx context.Context, input *meroxa.CreateApplicationInput) (*meroxa.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApplication", ctx, input)
	ret0, _ := ret[0].(*meroxa.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApplication indicates an expected call of CreateApplication.
func (mr *MockClientMockRecorder) CreateApplication(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApplication", reflect.TypeOf((*MockClient)(nil).CreateApplication), ctx, input)
}

// CreateApplicationV2 mocks base method.
func (m *MockClient) CreateApplicationV2(ctx context.Context, input *meroxa.CreateApplicationInput) (*meroxa.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApplicationV2", ctx, input)
	ret0, _ := ret[0].(*meroxa.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApplicationV2 indicates an expected call of CreateApplicationV2.
func (mr *MockClientMockRecorder) CreateApplicationV2(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApplicationV2", reflect.TypeOf((*MockClient)(nil).CreateApplicationV2), ctx, input)
}

// CreateBuild mocks base method.
func (m *MockClient) CreateBuild(ctx context.Context, input *meroxa.CreateBuildInput) (*meroxa.Build, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBuild", ctx, input)
	ret0, _ := ret[0].(*meroxa.Build)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBuild indicates an expected call of CreateBuild.
func (mr *MockClientMockRecorder) CreateBuild(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBuild", reflect.TypeOf((*MockClient)(nil).CreateBuild), ctx, input)
}

// CreateConnector mocks base method.
func (m *MockClient) CreateConnector(ctx context.Context, input *meroxa.CreateConnectorInput) (*meroxa.Connector, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConnector", ctx, input)
	ret0, _ := ret[0].(*meroxa.Connector)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateConnector indicates an expected call of CreateConnector.
func (mr *MockClientMockRecorder) CreateConnector(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConnector", reflect.TypeOf((*MockClient)(nil).CreateConnector), ctx, input)
}

// CreateDeployment mocks base method.
func (m *MockClient) CreateDeployment(ctx context.Context, input *meroxa.CreateDeploymentInput) (*meroxa.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeployment", ctx, input)
	ret0, _ := ret[0].(*meroxa.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDeployment indicates an expected call of CreateDeployment.
func (mr *MockClientMockRecorder) CreateDeployment(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeployment", reflect.TypeOf((*MockClient)(nil).CreateDeployment), ctx, input)
}

// CreateEnvironment mocks base method.
func (m *MockClient) CreateEnvironment(ctx context.Context, input *meroxa.CreateEnvironmentInput) (*meroxa.Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEnvironment", ctx, input)
	ret0, _ := ret[0].(*meroxa.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEnvironment indicates an expected call of CreateEnvironment.
func (mr *MockClientMockRecorder) CreateEnvironment(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEnvironment", reflect.TypeOf((*MockClient)(nil).CreateEnvironment), ctx, input)
}

// CreateFunction mocks base method.
func (m *MockClient) CreateFunction(ctx context.Context, input *meroxa.CreateFunctionInput) (*meroxa.Function, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFunction", ctx, input)
	ret0, _ := ret[0].(*meroxa.Function)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFunction indicates an expected call of CreateFunction.
func (mr *MockClientMockRecorder) CreateFunction(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFunction", reflect.TypeOf((*MockClient)(nil).CreateFunction), ctx, input)
}

// CreatePipeline mocks base method.
func (m *MockClient) CreatePipeline(ctx context.Context, input *meroxa.CreatePipelineInput) (*meroxa.Pipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePipeline", ctx, input)
	ret0, _ := ret[0].(*meroxa.Pipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePipeline indicates an expected call of CreatePipeline.
func (mr *MockClientMockRecorder) CreatePipeline(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePipeline", reflect.TypeOf((*MockClient)(nil).CreatePipeline), ctx, input)
}

// CreateResource mocks base method.
func (m *MockClient) CreateResource(ctx context.Context, input *meroxa.CreateResourceInput) (*meroxa.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateResource", ctx, input)
	ret0, _ := ret[0].(*meroxa.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateResource indicates an expected call of CreateResource.
func (mr *MockClientMockRecorder) CreateResource(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateResource", reflect.TypeOf((*MockClient)(nil).CreateResource), ctx, input)
}

// CreateSource mocks base method.
func (m *MockClient) CreateSource(ctx context.Context) (*meroxa.Source, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSource", ctx)
	ret0, _ := ret[0].(*meroxa.Source)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSource indicates an expected call of CreateSource.
func (mr *MockClientMockRecorder) CreateSource(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSource", reflect.TypeOf((*MockClient)(nil).CreateSource), ctx)
}

// DeleteApplication mocks base method.
func (m *MockClient) DeleteApplication(ctx context.Context, nameOrUUID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApplication", ctx, nameOrUUID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteApplication indicates an expected call of DeleteApplication.
func (mr *MockClientMockRecorder) DeleteApplication(ctx, nameOrUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApplication", reflect.TypeOf((*MockClient)(nil).DeleteApplication), ctx, nameOrUUID)
}

// DeleteApplicationEntities mocks base method.
func (m *MockClient) DeleteApplicationEntities(ctx context.Context, nameOrUUID string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApplicationEntities", ctx, nameOrUUID)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteApplicationEntities indicates an expected call of DeleteApplicationEntities.
func (mr *MockClientMockRecorder) DeleteApplicationEntities(ctx, nameOrUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApplicationEntities", reflect.TypeOf((*MockClient)(nil).DeleteApplicationEntities), ctx, nameOrUUID)
}

// DeleteConnector mocks base method.
func (m *MockClient) DeleteConnector(ctx context.Context, nameOrID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteConnector", ctx, nameOrID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteConnector indicates an expected call of DeleteConnector.
func (mr *MockClientMockRecorder) DeleteConnector(ctx, nameOrID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConnector", reflect.TypeOf((*MockClient)(nil).DeleteConnector), ctx, nameOrID)
}

// DeleteEnvironment mocks base method.
func (m *MockClient) DeleteEnvironment(ctx context.Context, nameOrUUID string) (*meroxa.Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEnvironment", ctx, nameOrUUID)
	ret0, _ := ret[0].(*meroxa.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEnvironment indicates an expected call of DeleteEnvironment.
func (mr *MockClientMockRecorder) DeleteEnvironment(ctx, nameOrUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEnvironment", reflect.TypeOf((*MockClient)(nil).DeleteEnvironment), ctx, nameOrUUID)
}

// DeleteFunction mocks base method.
func (m *MockClient) DeleteFunction(ctx context.Context, nameOrUUID string) (*meroxa.Function, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFunction", ctx, nameOrUUID)
	ret0, _ := ret[0].(*meroxa.Function)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFunction indicates an expected call of DeleteFunction.
func (mr *MockClientMockRecorder) DeleteFunction(ctx, nameOrUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFunction", reflect.TypeOf((*MockClient)(nil).DeleteFunction), ctx, nameOrUUID)
}

// DeletePipeline mocks base method.
func (m *MockClient) DeletePipeline(ctx context.Context, nameOrID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePipeline", ctx, nameOrID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePipeline indicates an expected call of DeletePipeline.
func (mr *MockClientMockRecorder) DeletePipeline(ctx, nameOrID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePipeline", reflect.TypeOf((*MockClient)(nil).DeletePipeline), ctx, nameOrID)
}

// DeleteResource mocks base method.
func (m *MockClient) DeleteResource(ctx context.Context, nameOrID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResource", ctx, nameOrID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteResource indicates an expected call of DeleteResource.
func (mr *MockClientMockRecorder) DeleteResource(ctx, nameOrID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResource", reflect.TypeOf((*MockClient)(nil).DeleteResource), ctx, nameOrID)
}

// GetApplication mocks base method.
func (m *MockClient) GetApplication(ctx context.Context, nameOrUUID string) (*meroxa.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplication", ctx, nameOrUUID)
	ret0, _ := ret[0].(*meroxa.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplication indicates an expected call of GetApplication.
func (mr *MockClientMockRecorder) GetApplication(ctx, nameOrUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplication", reflect.TypeOf((*MockClient)(nil).GetApplication), ctx, nameOrUUID)
}

// GetApplicationLogs mocks base method.
func (m *MockClient) GetApplicationLogs(ctx context.Context, nameOrUUID string) (*meroxa.ApplicationLogs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplicationLogs", ctx, nameOrUUID)
	ret0, _ := ret[0].(*meroxa.ApplicationLogs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplicationLogs indicates an expected call of GetApplicationLogs.
func (mr *MockClientMockRecorder) GetApplicationLogs(ctx, nameOrUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplicationLogs", reflect.TypeOf((*MockClient)(nil).GetApplicationLogs), ctx, nameOrUUID)
}

// GetBuild mocks base method.
func (m *MockClient) GetBuild(ctx context.Context, uuid string) (*meroxa.Build, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBuild", ctx, uuid)
	ret0, _ := ret[0].(*meroxa.Build)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBuild indicates an expected call of GetBuild.
func (mr *MockClientMockRecorder) GetBuild(ctx, uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBuild", reflect.TypeOf((*MockClient)(nil).GetBuild), ctx, uuid)
}

// GetBuildLogs mocks base method.
func (m *MockClient) GetBuildLogs(ctx context.Context, uuid string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBuildLogs", ctx, uuid)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBuildLogs indicates an expected call of GetBuildLogs.
func (mr *MockClientMockRecorder) GetBuildLogs(ctx, uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBuildLogs", reflect.TypeOf((*MockClient)(nil).GetBuildLogs), ctx, uuid)
}

// GetConnectorByNameOrID mocks base method.
func (m *MockClient) GetConnectorByNameOrID(ctx context.Context, nameOrID string) (*meroxa.Connector, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnectorByNameOrID", ctx, nameOrID)
	ret0, _ := ret[0].(*meroxa.Connector)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConnectorByNameOrID indicates an expected call of GetConnectorByNameOrID.
func (mr *MockClientMockRecorder) GetConnectorByNameOrID(ctx, nameOrID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnectorByNameOrID", reflect.TypeOf((*MockClient)(nil).GetConnectorByNameOrID), ctx, nameOrID)
}

// GetConnectorLogs mocks base method.
func (m *MockClient) GetConnectorLogs(ctx context.Context, nameOrID string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnectorLogs", ctx, nameOrID)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConnectorLogs indicates an expected call of GetConnectorLogs.
func (mr *MockClientMockRecorder) GetConnectorLogs(ctx, nameOrID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnectorLogs", reflect.TypeOf((*MockClient)(nil).GetConnectorLogs), ctx, nameOrID)
}

// GetDeployment mocks base method.
func (m *MockClient) GetDeployment(ctx context.Context, appIdentifier, depUUID string) (*meroxa.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployment", ctx, appIdentifier, depUUID)
	ret0, _ := ret[0].(*meroxa.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployment indicates an expected call of GetDeployment.
func (mr *MockClientMockRecorder) GetDeployment(ctx, appIdentifier, depUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployment", reflect.TypeOf((*MockClient)(nil).GetDeployment), ctx, appIdentifier, depUUID)
}

// GetEnvironment mocks base method.
func (m *MockClient) GetEnvironment(ctx context.Context, nameOrUUID string) (*meroxa.Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvironment", ctx, nameOrUUID)
	ret0, _ := ret[0].(*meroxa.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvironment indicates an expected call of GetEnvironment.
func (mr *MockClientMockRecorder) GetEnvironment(ctx, nameOrUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironment", reflect.TypeOf((*MockClient)(nil).GetEnvironment), ctx, nameOrUUID)
}

// GetFunction mocks base method.
func (m *MockClient) GetFunction(ctx context.Context, nameOrUUID string) (*meroxa.Function, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFunction", ctx, nameOrUUID)
	ret0, _ := ret[0].(*meroxa.Function)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFunction indicates an expected call of GetFunction.
func (mr *MockClientMockRecorder) GetFunction(ctx, nameOrUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFunction", reflect.TypeOf((*MockClient)(nil).GetFunction), ctx, nameOrUUID)
}

// GetFunctionLogs mocks base method.
func (m *MockClient) GetFunctionLogs(ctx context.Context, nameOrUUID string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFunctionLogs", ctx, nameOrUUID)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFunctionLogs indicates an expected call of GetFunctionLogs.
func (mr *MockClientMockRecorder) GetFunctionLogs(ctx, nameOrUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFunctionLogs", reflect.TypeOf((*MockClient)(nil).GetFunctionLogs), ctx, nameOrUUID)
}

// GetLatestDeployment mocks base method.
func (m *MockClient) GetLatestDeployment(ctx context.Context, appIdentifier string) (*meroxa.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestDeployment", ctx, appIdentifier)
	ret0, _ := ret[0].(*meroxa.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestDeployment indicates an expected call of GetLatestDeployment.
func (mr *MockClientMockRecorder) GetLatestDeployment(ctx, appIdentifier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestDeployment", reflect.TypeOf((*MockClient)(nil).GetLatestDeployment), ctx, appIdentifier)
}

// GetPipeline mocks base method.
func (m *MockClient) GetPipeline(ctx context.Context, pipelineID int) (*meroxa.Pipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPipeline", ctx, pipelineID)
	ret0, _ := ret[0].(*meroxa.Pipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPipeline indicates an expected call of GetPipeline.
func (mr *MockClientMockRecorder) GetPipeline(ctx, pipelineID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPipeline", reflect.TypeOf((*MockClient)(nil).GetPipeline), ctx, pipelineID)
}

// GetPipelineByName mocks base method.
func (m *MockClient) GetPipelineByName(ctx context.Context, name string) (*meroxa.Pipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPipelineByName", ctx, name)
	ret0, _ := ret[0].(*meroxa.Pipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPipelineByName indicates an expected call of GetPipelineByName.
func (mr *MockClientMockRecorder) GetPipelineByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPipelineByName", reflect.TypeOf((*MockClient)(nil).GetPipelineByName), ctx, name)
}

// GetResourceByNameOrID mocks base method.
func (m *MockClient) GetResourceByNameOrID(ctx context.Context, nameOrID string) (*meroxa.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceByNameOrID", ctx, nameOrID)
	ret0, _ := ret[0].(*meroxa.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceByNameOrID indicates an expected call of GetResourceByNameOrID.
func (mr *MockClientMockRecorder) GetResourceByNameOrID(ctx, nameOrID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceByNameOrID", reflect.TypeOf((*MockClient)(nil).GetResourceByNameOrID), ctx, nameOrID)
}

// GetUser mocks base method.
func (m *MockClient) GetUser(ctx context.Context) (*meroxa.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx)
	ret0, _ := ret[0].(*meroxa.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockClientMockRecorder) GetUser(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockClient)(nil).GetUser), ctx)
}

// IntrospectResource mocks base method.
func (m *MockClient) IntrospectResource(ctx context.Context, nameOrID string) (*meroxa.ResourceIntrospection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntrospectResource", ctx, nameOrID)
	ret0, _ := ret[0].(*meroxa.ResourceIntrospection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IntrospectResource indicates an expected call of IntrospectResource.
func (mr *MockClientMockRecorder) IntrospectResource(ctx, nameOrID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntrospectResource", reflect.TypeOf((*MockClient)(nil).IntrospectResource), ctx, nameOrID)
}

// ListAccounts mocks base method.
func (m *MockClient) ListAccounts(ctx context.Context) ([]*meroxa.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", ctx)
	ret0, _ := ret[0].([]*meroxa.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockClientMockRecorder) ListAccounts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockClient)(nil).ListAccounts), ctx)
}

// ListApplications mocks base method.
func (m *MockClient) ListApplications(ctx context.Context) ([]*meroxa.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApplications", ctx)
	ret0, _ := ret[0].([]*meroxa.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApplications indicates an expected call of ListApplications.
func (mr *MockClientMockRecorder) ListApplications(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplications", reflect.TypeOf((*MockClient)(nil).ListApplications), ctx)
}

// ListConnectors mocks base method.
func (m *MockClient) ListConnectors(ctx context.Context) ([]*meroxa.Connector, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListConnectors", ctx)
	ret0, _ := ret[0].([]*meroxa.Connector)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConnectors indicates an expected call of ListConnectors.
func (mr *MockClientMockRecorder) ListConnectors(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConnectors", reflect.TypeOf((*MockClient)(nil).ListConnectors), ctx)
}

// ListEnvironments mocks base method.
func (m *MockClient) ListEnvironments(ctx context.Context) ([]*meroxa.Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEnvironments", ctx)
	ret0, _ := ret[0].([]*meroxa.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnvironments indicates an expected call of ListEnvironments.
func (mr *MockClientMockRecorder) ListEnvironments(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnvironments", reflect.TypeOf((*MockClient)(nil).ListEnvironments), ctx)
}

// ListFunctions mocks base method.
func (m *MockClient) ListFunctions(ctx context.Context) ([]*meroxa.Function, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFunctions", ctx)
	ret0, _ := ret[0].([]*meroxa.Function)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFunctions indicates an expected call of ListFunctions.
func (mr *MockClientMockRecorder) ListFunctions(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFunctions", reflect.TypeOf((*MockClient)(nil).ListFunctions), ctx)
}

// ListPipelineConnectors mocks base method.
func (m *MockClient) ListPipelineConnectors(ctx context.Context, pipelineNameOrID string) ([]*meroxa.Connector, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPipelineConnectors", ctx, pipelineNameOrID)
	ret0, _ := ret[0].([]*meroxa.Connector)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPipelineConnectors indicates an expected call of ListPipelineConnectors.
func (mr *MockClientMockRecorder) ListPipelineConnectors(ctx, pipelineNameOrID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPipelineConnectors", reflect.TypeOf((*MockClient)(nil).ListPipelineConnectors), ctx, pipelineNameOrID)
}

// ListPipelines mocks base method.
func (m *MockClient) ListPipelines(ctx context.Context) ([]*meroxa.Pipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPipelines", ctx)
	ret0, _ := ret[0].([]*meroxa.Pipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPipelines indicates an expected call of ListPipelines.
func (mr *MockClientMockRecorder) ListPipelines(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPipelines", reflect.TypeOf((*MockClient)(nil).ListPipelines), ctx)
}

// ListResourceTypes mocks base method.
func (m *MockClient) ListResourceTypes(ctx context.Context) ([]meroxa.ResourceType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResourceTypes", ctx)
	ret0, _ := ret[0].([]meroxa.ResourceType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourceTypes indicates an expected call of ListResourceTypes.
func (mr *MockClientMockRecorder) ListResourceTypes(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceTypes", reflect.TypeOf((*MockClient)(nil).ListResourceTypes), ctx)
}

// ListResources mocks base method.
func (m *MockClient) ListResources(ctx context.Context) ([]*meroxa.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResources", ctx)
	ret0, _ := ret[0].([]*meroxa.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResources indicates an expected call of ListResources.
func (mr *MockClientMockRecorder) ListResources(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResources", reflect.TypeOf((*MockClient)(nil).ListResources), ctx)
}

// ListTransforms mocks base method.
func (m *MockClient) ListTransforms(ctx context.Context) ([]*meroxa.Transform, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTransforms", ctx)
	ret0, _ := ret[0].([]*meroxa.Transform)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTransforms indicates an expected call of ListTransforms.
func (mr *MockClientMockRecorder) ListTransforms(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTransforms", reflect.TypeOf((*MockClient)(nil).ListTransforms), ctx)
}

// MakeRequest mocks base method.
func (m *MockClient) MakeRequest(ctx context.Context, method, path string, body interface{}, params url.Values, headers http.Header) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeRequest", ctx, method, path, body, params, headers)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeRequest indicates an expected call of MakeRequest.
func (mr *MockClientMockRecorder) MakeRequest(ctx, method, path, body, params, headers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRequest", reflect.TypeOf((*MockClient)(nil).MakeRequest), ctx, method, path, body, params, headers)
}

// PerformActionOnEnvironment mocks base method.
func (m *MockClient) PerformActionOnEnvironment(ctx context.Context, nameOrUUID string, input *meroxa.RepairEnvironmentInput) (*meroxa.Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PerformActionOnEnvironment", ctx, nameOrUUID, input)
	ret0, _ := ret[0].(*meroxa.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PerformActionOnEnvironment indicates an expected call of PerformActionOnEnvironment.
func (mr *MockClientMockRecorder) PerformActionOnEnvironment(ctx, nameOrUUID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformActionOnEnvironment", reflect.TypeOf((*MockClient)(nil).PerformActionOnEnvironment), ctx, nameOrUUID, input)
}

// RotateTunnelKeyForResource mocks base method.
func (m *MockClient) RotateTunnelKeyForResource(ctx context.Context, nameOrID string) (*meroxa.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RotateTunnelKeyForResource", ctx, nameOrID)
	ret0, _ := ret[0].(*meroxa.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RotateTunnelKeyForResource indicates an expected call of RotateTunnelKeyForResource.
func (mr *MockClientMockRecorder) RotateTunnelKeyForResource(ctx, nameOrID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotateTunnelKeyForResource", reflect.TypeOf((*MockClient)(nil).RotateTunnelKeyForResource), ctx, nameOrID)
}

// UpdateConnector mocks base method.
func (m *MockClient) UpdateConnector(ctx context.Context, nameOrID string, input *meroxa.UpdateConnectorInput) (*meroxa.Connector, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConnector", ctx, nameOrID, input)
	ret0, _ := ret[0].(*meroxa.Connector)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateConnector indicates an expected call of UpdateConnector.
func (mr *MockClientMockRecorder) UpdateConnector(ctx, nameOrID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConnector", reflect.TypeOf((*MockClient)(nil).UpdateConnector), ctx, nameOrID, input)
}

// UpdateConnectorStatus mocks base method.
func (m *MockClient) UpdateConnectorStatus(ctx context.Context, nameOrID string, state meroxa.Action) (*meroxa.Connector, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConnectorStatus", ctx, nameOrID, state)
	ret0, _ := ret[0].(*meroxa.Connector)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateConnectorStatus indicates an expected call of UpdateConnectorStatus.
func (mr *MockClientMockRecorder) UpdateConnectorStatus(ctx, nameOrID, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConnectorStatus", reflect.TypeOf((*MockClient)(nil).UpdateConnectorStatus), ctx, nameOrID, state)
}

// UpdateEnvironment mocks base method.
func (m *MockClient) UpdateEnvironment(ctx context.Context, nameOrUUID string, input *meroxa.UpdateEnvironmentInput) (*meroxa.Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEnvironment", ctx, nameOrUUID, input)
	ret0, _ := ret[0].(*meroxa.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEnvironment indicates an expected call of UpdateEnvironment.
func (mr *MockClientMockRecorder) UpdateEnvironment(ctx, nameOrUUID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEnvironment", reflect.TypeOf((*MockClient)(nil).UpdateEnvironment), ctx, nameOrUUID, input)
}

// UpdatePipeline mocks base method.
func (m *MockClient) UpdatePipeline(ctx context.Context, pipelineNameOrID string, input *meroxa.UpdatePipelineInput) (*meroxa.Pipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePipeline", ctx, pipelineNameOrID, input)
	ret0, _ := ret[0].(*meroxa.Pipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePipeline indicates an expected call of UpdatePipeline.
func (mr *MockClientMockRecorder) UpdatePipeline(ctx, pipelineNameOrID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePipeline", reflect.TypeOf((*MockClient)(nil).UpdatePipeline), ctx, pipelineNameOrID, input)
}

// UpdatePipelineStatus mocks base method.
func (m *MockClient) UpdatePipelineStatus(ctx context.Context, pipelineNameOrID string, action meroxa.Action) (*meroxa.Pipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePipelineStatus", ctx, pipelineNameOrID, action)
	ret0, _ := ret[0].(*meroxa.Pipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePipelineStatus indicates an expected call of UpdatePipelineStatus.
func (mr *MockClientMockRecorder) UpdatePipelineStatus(ctx, pipelineNameOrID, action interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePipelineStatus", reflect.TypeOf((*MockClient)(nil).UpdatePipelineStatus), ctx, pipelineNameOrID, action)
}

// UpdateResource mocks base method.
func (m *MockClient) UpdateResource(ctx context.Context, nameOrID string, input *meroxa.UpdateResourceInput) (*meroxa.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateResource", ctx, nameOrID, input)
	ret0, _ := ret[0].(*meroxa.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateResource indicates an expected call of UpdateResource.
func (mr *MockClientMockRecorder) UpdateResource(ctx, nameOrID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResource", reflect.TypeOf((*MockClient)(nil).UpdateResource), ctx, nameOrID, input)
}

// ValidateResource mocks base method.
func (m *MockClient) ValidateResource(ctx context.Context, nameOrID string) (*meroxa.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateResource", ctx, nameOrID)
	ret0, _ := ret[0].(*meroxa.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateResource indicates an expected call of ValidateResource.
func (mr *MockClientMockRecorder) ValidateResource(ctx, nameOrID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateResource", reflect.TypeOf((*MockClient)(nil).ValidateResource), ctx, nameOrID)
}
