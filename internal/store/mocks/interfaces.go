// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	core "github.com/hashicorp/consul-api-gateway/internal/core"
	store "github.com/hashicorp/consul-api-gateway/internal/store"
)

// MockGateway is a mock of Gateway interface.
type MockGateway struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayMockRecorder
}

// MockGatewayMockRecorder is the mock recorder for MockGateway.
type MockGatewayMockRecorder struct {
	mock *MockGateway
}

// NewMockGateway creates a new mock instance.
func NewMockGateway(ctrl *gomock.Controller) *MockGateway {
	mock := &MockGateway{ctrl: ctrl}
	mock.recorder = &MockGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGateway) EXPECT() *MockGatewayMockRecorder {
	return m.recorder
}

// CanFetchSecrets mocks base method.
func (m *MockGateway) CanFetchSecrets(secrets []string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanFetchSecrets", secrets)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanFetchSecrets indicates an expected call of CanFetchSecrets.
func (mr *MockGatewayMockRecorder) CanFetchSecrets(secrets interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanFetchSecrets", reflect.TypeOf((*MockGateway)(nil).CanFetchSecrets), secrets)
}

// ID mocks base method.
func (m *MockGateway) ID() core.GatewayID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(core.GatewayID)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockGatewayMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockGateway)(nil).ID))
}

// Resolve mocks base method.
func (m *MockGateway) Resolve() core.ResolvedGateway {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resolve")
	ret0, _ := ret[0].(core.ResolvedGateway)
	return ret0
}

// Resolve indicates an expected call of Resolve.
func (mr *MockGatewayMockRecorder) Resolve() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockGateway)(nil).Resolve))
}

// MockStatusTrackingRoute is a mock of StatusTrackingRoute interface.
type MockStatusTrackingRoute struct {
	ctrl     *gomock.Controller
	recorder *MockStatusTrackingRouteMockRecorder
}

// MockStatusTrackingRouteMockRecorder is the mock recorder for MockStatusTrackingRoute.
type MockStatusTrackingRouteMockRecorder struct {
	mock *MockStatusTrackingRoute
}

// NewMockStatusTrackingRoute creates a new mock instance.
func NewMockStatusTrackingRoute(ctrl *gomock.Controller) *MockStatusTrackingRoute {
	mock := &MockStatusTrackingRoute{ctrl: ctrl}
	mock.recorder = &MockStatusTrackingRouteMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatusTrackingRoute) EXPECT() *MockStatusTrackingRouteMockRecorder {
	return m.recorder
}

// ID mocks base method.
func (m *MockStatusTrackingRoute) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockStatusTrackingRouteMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockStatusTrackingRoute)(nil).ID))
}

// OnGatewayRemoved mocks base method.
func (m *MockStatusTrackingRoute) OnGatewayRemoved(gateway store.Gateway) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnGatewayRemoved", gateway)
}

// OnGatewayRemoved indicates an expected call of OnGatewayRemoved.
func (mr *MockStatusTrackingRouteMockRecorder) OnGatewayRemoved(gateway interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnGatewayRemoved", reflect.TypeOf((*MockStatusTrackingRoute)(nil).OnGatewayRemoved), gateway)
}

// SyncStatus mocks base method.
func (m *MockStatusTrackingRoute) SyncStatus(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatus", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncStatus indicates an expected call of SyncStatus.
func (mr *MockStatusTrackingRouteMockRecorder) SyncStatus(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatus", reflect.TypeOf((*MockStatusTrackingRoute)(nil).SyncStatus), ctx)
}

// MockRoute is a mock of Route interface.
type MockRoute struct {
	ctrl     *gomock.Controller
	recorder *MockRouteMockRecorder
}

// MockRouteMockRecorder is the mock recorder for MockRoute.
type MockRouteMockRecorder struct {
	mock *MockRoute
}

// NewMockRoute creates a new mock instance.
func NewMockRoute(ctrl *gomock.Controller) *MockRoute {
	mock := &MockRoute{ctrl: ctrl}
	mock.recorder = &MockRouteMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoute) EXPECT() *MockRouteMockRecorder {
	return m.recorder
}

// ID mocks base method.
func (m *MockRoute) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockRouteMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockRoute)(nil).ID))
}

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// DeleteGateway mocks base method.
func (m *MockStore) DeleteGateway(ctx context.Context, id core.GatewayID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGateway", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGateway indicates an expected call of DeleteGateway.
func (mr *MockStoreMockRecorder) DeleteGateway(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGateway", reflect.TypeOf((*MockStore)(nil).DeleteGateway), ctx, id)
}

// DeleteRoute mocks base method.
func (m *MockStore) DeleteRoute(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRoute", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRoute indicates an expected call of DeleteRoute.
func (mr *MockStoreMockRecorder) DeleteRoute(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoute", reflect.TypeOf((*MockStore)(nil).DeleteRoute), ctx, id)
}

// GetGateway mocks base method.
func (m *MockStore) GetGateway(ctx context.Context, id core.GatewayID) (store.Gateway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGateway", ctx, id)
	ret0, _ := ret[0].(store.Gateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGateway indicates an expected call of GetGateway.
func (mr *MockStoreMockRecorder) GetGateway(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGateway", reflect.TypeOf((*MockStore)(nil).GetGateway), ctx, id)
}

// Sync mocks base method.
func (m *MockStore) Sync(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync.
func (mr *MockStoreMockRecorder) Sync(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockStore)(nil).Sync), ctx)
}

// UpsertGateway mocks base method.
func (m *MockStore) UpsertGateway(ctx context.Context, gateway store.Gateway, updateConditionFn func(store.Gateway) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertGateway", ctx, gateway, updateConditionFn)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertGateway indicates an expected call of UpsertGateway.
func (mr *MockStoreMockRecorder) UpsertGateway(ctx, gateway, updateConditionFn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertGateway", reflect.TypeOf((*MockStore)(nil).UpsertGateway), ctx, gateway, updateConditionFn)
}

// UpsertRoute mocks base method.
func (m *MockStore) UpsertRoute(ctx context.Context, route store.Route, updateConditionFn func(store.Route) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertRoute", ctx, route, updateConditionFn)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRoute indicates an expected call of UpsertRoute.
func (mr *MockStoreMockRecorder) UpsertRoute(ctx, route, updateConditionFn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRoute", reflect.TypeOf((*MockStore)(nil).UpsertRoute), ctx, route, updateConditionFn)
}

// MockBackend is a mock of Backend interface.
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend.
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance.
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// DeleteGateway mocks base method.
func (m *MockBackend) DeleteGateway(ctx context.Context, id core.GatewayID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGateway", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGateway indicates an expected call of DeleteGateway.
func (mr *MockBackendMockRecorder) DeleteGateway(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGateway", reflect.TypeOf((*MockBackend)(nil).DeleteGateway), ctx, id)
}

// DeleteRoute mocks base method.
func (m *MockBackend) DeleteRoute(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRoute", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRoute indicates an expected call of DeleteRoute.
func (mr *MockBackendMockRecorder) DeleteRoute(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoute", reflect.TypeOf((*MockBackend)(nil).DeleteRoute), ctx, id)
}

// GetGateway mocks base method.
func (m *MockBackend) GetGateway(ctx context.Context, id core.GatewayID) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGateway", ctx, id)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGateway indicates an expected call of GetGateway.
func (mr *MockBackendMockRecorder) GetGateway(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGateway", reflect.TypeOf((*MockBackend)(nil).GetGateway), ctx, id)
}

// GetRoute mocks base method.
func (m *MockBackend) GetRoute(ctx context.Context, id string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoute", ctx, id)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoute indicates an expected call of GetRoute.
func (mr *MockBackendMockRecorder) GetRoute(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoute", reflect.TypeOf((*MockBackend)(nil).GetRoute), ctx, id)
}

// ListGateways mocks base method.
func (m *MockBackend) ListGateways(ctx context.Context) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGateways", ctx)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGateways indicates an expected call of ListGateways.
func (mr *MockBackendMockRecorder) ListGateways(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGateways", reflect.TypeOf((*MockBackend)(nil).ListGateways), ctx)
}

// ListRoutes mocks base method.
func (m *MockBackend) ListRoutes(ctx context.Context) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoutes", ctx)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoutes indicates an expected call of ListRoutes.
func (mr *MockBackendMockRecorder) ListRoutes(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoutes", reflect.TypeOf((*MockBackend)(nil).ListRoutes), ctx)
}

// UpsertGateways mocks base method.
func (m *MockBackend) UpsertGateways(ctx context.Context, gateways ...store.GatewayRecord) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range gateways {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertGateways", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertGateways indicates an expected call of UpsertGateways.
func (mr *MockBackendMockRecorder) UpsertGateways(ctx interface{}, gateways ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, gateways...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertGateways", reflect.TypeOf((*MockBackend)(nil).UpsertGateways), varargs...)
}

// UpsertRoutes mocks base method.
func (m *MockBackend) UpsertRoutes(ctx context.Context, routes ...store.RouteRecord) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range routes {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertRoutes", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertRoutes indicates an expected call of UpsertRoutes.
func (mr *MockBackendMockRecorder) UpsertRoutes(ctx interface{}, routes ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, routes...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertRoutes", reflect.TypeOf((*MockBackend)(nil).UpsertRoutes), varargs...)
}

// MockMarshaler is a mock of Marshaler interface.
type MockMarshaler struct {
	ctrl     *gomock.Controller
	recorder *MockMarshalerMockRecorder
}

// MockMarshalerMockRecorder is the mock recorder for MockMarshaler.
type MockMarshalerMockRecorder struct {
	mock *MockMarshaler
}

// NewMockMarshaler creates a new mock instance.
func NewMockMarshaler(ctrl *gomock.Controller) *MockMarshaler {
	mock := &MockMarshaler{ctrl: ctrl}
	mock.recorder = &MockMarshalerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMarshaler) EXPECT() *MockMarshalerMockRecorder {
	return m.recorder
}

// MarshalGateway mocks base method.
func (m *MockMarshaler) MarshalGateway(arg0 store.Gateway) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalGateway", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalGateway indicates an expected call of MarshalGateway.
func (mr *MockMarshalerMockRecorder) MarshalGateway(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalGateway", reflect.TypeOf((*MockMarshaler)(nil).MarshalGateway), arg0)
}

// MarshalRoute mocks base method.
func (m *MockMarshaler) MarshalRoute(arg0 store.Route) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalRoute", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalRoute indicates an expected call of MarshalRoute.
func (mr *MockMarshalerMockRecorder) MarshalRoute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalRoute", reflect.TypeOf((*MockMarshaler)(nil).MarshalRoute), arg0)
}

// UnmarshalGateway mocks base method.
func (m *MockMarshaler) UnmarshalGateway(data []byte) (store.Gateway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnmarshalGateway", data)
	ret0, _ := ret[0].(store.Gateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnmarshalGateway indicates an expected call of UnmarshalGateway.
func (mr *MockMarshalerMockRecorder) UnmarshalGateway(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalGateway", reflect.TypeOf((*MockMarshaler)(nil).UnmarshalGateway), data)
}

// UnmarshalRoute mocks base method.
func (m *MockMarshaler) UnmarshalRoute(data []byte) (store.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnmarshalRoute", data)
	ret0, _ := ret[0].(store.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnmarshalRoute indicates an expected call of UnmarshalRoute.
func (mr *MockMarshalerMockRecorder) UnmarshalRoute(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalRoute", reflect.TypeOf((*MockMarshaler)(nil).UnmarshalRoute), data)
}
