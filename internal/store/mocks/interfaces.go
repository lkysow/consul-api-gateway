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

// MockStatusTrackingGateway is a mock of StatusTrackingGateway interface.
type MockStatusTrackingGateway struct {
	ctrl     *gomock.Controller
	recorder *MockStatusTrackingGatewayMockRecorder
}

// MockStatusTrackingGatewayMockRecorder is the mock recorder for MockStatusTrackingGateway.
type MockStatusTrackingGatewayMockRecorder struct {
	mock *MockStatusTrackingGateway
}

// NewMockStatusTrackingGateway creates a new mock instance.
func NewMockStatusTrackingGateway(ctrl *gomock.Controller) *MockStatusTrackingGateway {
	mock := &MockStatusTrackingGateway{ctrl: ctrl}
	mock.recorder = &MockStatusTrackingGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatusTrackingGateway) EXPECT() *MockStatusTrackingGatewayMockRecorder {
	return m.recorder
}

// ID mocks base method.
func (m *MockStatusTrackingGateway) ID() core.GatewayID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(core.GatewayID)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockStatusTrackingGatewayMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockStatusTrackingGateway)(nil).ID))
}

// Listeners mocks base method.
func (m *MockStatusTrackingGateway) Listeners() []store.Listener {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Listeners")
	ret0, _ := ret[0].([]store.Listener)
	return ret0
}

// Listeners indicates an expected call of Listeners.
func (mr *MockStatusTrackingGatewayMockRecorder) Listeners() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listeners", reflect.TypeOf((*MockStatusTrackingGateway)(nil).Listeners))
}

// Meta mocks base method.
func (m *MockStatusTrackingGateway) Meta() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meta")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Meta indicates an expected call of Meta.
func (mr *MockStatusTrackingGatewayMockRecorder) Meta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meta", reflect.TypeOf((*MockStatusTrackingGateway)(nil).Meta))
}

// ShouldBind mocks base method.
func (m *MockStatusTrackingGateway) ShouldBind(route store.Route) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldBind", route)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ShouldBind indicates an expected call of ShouldBind.
func (mr *MockStatusTrackingGatewayMockRecorder) ShouldBind(route interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldBind", reflect.TypeOf((*MockStatusTrackingGateway)(nil).ShouldBind), route)
}

// ShouldUpdate mocks base method.
func (m *MockStatusTrackingGateway) ShouldUpdate(other store.Gateway) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldUpdate", other)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ShouldUpdate indicates an expected call of ShouldUpdate.
func (mr *MockStatusTrackingGatewayMockRecorder) ShouldUpdate(other interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldUpdate", reflect.TypeOf((*MockStatusTrackingGateway)(nil).ShouldUpdate), other)
}

// TrackSync mocks base method.
func (m *MockStatusTrackingGateway) TrackSync(ctx context.Context, sync func() (bool, error)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrackSync", ctx, sync)
	ret0, _ := ret[0].(error)
	return ret0
}

// TrackSync indicates an expected call of TrackSync.
func (mr *MockStatusTrackingGatewayMockRecorder) TrackSync(ctx, sync interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrackSync", reflect.TypeOf((*MockStatusTrackingGateway)(nil).TrackSync), ctx, sync)
}

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

// Listeners mocks base method.
func (m *MockGateway) Listeners() []store.Listener {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Listeners")
	ret0, _ := ret[0].([]store.Listener)
	return ret0
}

// Listeners indicates an expected call of Listeners.
func (mr *MockGatewayMockRecorder) Listeners() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listeners", reflect.TypeOf((*MockGateway)(nil).Listeners))
}

// Meta mocks base method.
func (m *MockGateway) Meta() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meta")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Meta indicates an expected call of Meta.
func (mr *MockGatewayMockRecorder) Meta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meta", reflect.TypeOf((*MockGateway)(nil).Meta))
}

// ShouldBind mocks base method.
func (m *MockGateway) ShouldBind(route store.Route) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldBind", route)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ShouldBind indicates an expected call of ShouldBind.
func (mr *MockGatewayMockRecorder) ShouldBind(route interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldBind", reflect.TypeOf((*MockGateway)(nil).ShouldBind), route)
}

// ShouldUpdate mocks base method.
func (m *MockGateway) ShouldUpdate(other store.Gateway) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldUpdate", other)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ShouldUpdate indicates an expected call of ShouldUpdate.
func (mr *MockGatewayMockRecorder) ShouldUpdate(other interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldUpdate", reflect.TypeOf((*MockGateway)(nil).ShouldUpdate), other)
}

// MockRouteTrackingListener is a mock of RouteTrackingListener interface.
type MockRouteTrackingListener struct {
	ctrl     *gomock.Controller
	recorder *MockRouteTrackingListenerMockRecorder
}

// MockRouteTrackingListenerMockRecorder is the mock recorder for MockRouteTrackingListener.
type MockRouteTrackingListenerMockRecorder struct {
	mock *MockRouteTrackingListener
}

// NewMockRouteTrackingListener creates a new mock instance.
func NewMockRouteTrackingListener(ctrl *gomock.Controller) *MockRouteTrackingListener {
	mock := &MockRouteTrackingListener{ctrl: ctrl}
	mock.recorder = &MockRouteTrackingListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteTrackingListener) EXPECT() *MockRouteTrackingListenerMockRecorder {
	return m.recorder
}

// CanBind mocks base method.
func (m *MockRouteTrackingListener) CanBind(ctx context.Context, route store.Route) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanBind", ctx, route)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanBind indicates an expected call of CanBind.
func (mr *MockRouteTrackingListenerMockRecorder) CanBind(ctx, route interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanBind", reflect.TypeOf((*MockRouteTrackingListener)(nil).CanBind), ctx, route)
}

// Config mocks base method.
func (m *MockRouteTrackingListener) Config() store.ListenerConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(store.ListenerConfig)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockRouteTrackingListenerMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockRouteTrackingListener)(nil).Config))
}

// ID mocks base method.
func (m *MockRouteTrackingListener) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockRouteTrackingListenerMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockRouteTrackingListener)(nil).ID))
}

// IsValid mocks base method.
func (m *MockRouteTrackingListener) IsValid() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValid")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValid indicates an expected call of IsValid.
func (mr *MockRouteTrackingListenerMockRecorder) IsValid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockRouteTrackingListener)(nil).IsValid))
}

// OnRouteAdded mocks base method.
func (m *MockRouteTrackingListener) OnRouteAdded(route store.Route) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnRouteAdded", route)
}

// OnRouteAdded indicates an expected call of OnRouteAdded.
func (mr *MockRouteTrackingListenerMockRecorder) OnRouteAdded(route interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnRouteAdded", reflect.TypeOf((*MockRouteTrackingListener)(nil).OnRouteAdded), route)
}

// OnRouteRemoved mocks base method.
func (m *MockRouteTrackingListener) OnRouteRemoved(id string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnRouteRemoved", id)
}

// OnRouteRemoved indicates an expected call of OnRouteRemoved.
func (mr *MockRouteTrackingListenerMockRecorder) OnRouteRemoved(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnRouteRemoved", reflect.TypeOf((*MockRouteTrackingListener)(nil).OnRouteRemoved), id)
}

// MockListener is a mock of Listener interface.
type MockListener struct {
	ctrl     *gomock.Controller
	recorder *MockListenerMockRecorder
}

// MockListenerMockRecorder is the mock recorder for MockListener.
type MockListenerMockRecorder struct {
	mock *MockListener
}

// NewMockListener creates a new mock instance.
func NewMockListener(ctrl *gomock.Controller) *MockListener {
	mock := &MockListener{ctrl: ctrl}
	mock.recorder = &MockListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListener) EXPECT() *MockListenerMockRecorder {
	return m.recorder
}

// CanBind mocks base method.
func (m *MockListener) CanBind(ctx context.Context, route store.Route) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanBind", ctx, route)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanBind indicates an expected call of CanBind.
func (mr *MockListenerMockRecorder) CanBind(ctx, route interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanBind", reflect.TypeOf((*MockListener)(nil).CanBind), ctx, route)
}

// Config mocks base method.
func (m *MockListener) Config() store.ListenerConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(store.ListenerConfig)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockListenerMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockListener)(nil).Config))
}

// ID mocks base method.
func (m *MockListener) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockListenerMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockListener)(nil).ID))
}

// IsValid mocks base method.
func (m *MockListener) IsValid() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValid")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValid indicates an expected call of IsValid.
func (mr *MockListenerMockRecorder) IsValid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockListener)(nil).IsValid))
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

// OnBindFailed mocks base method.
func (m *MockStatusTrackingRoute) OnBindFailed(err error, gateway store.Gateway) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnBindFailed", err, gateway)
}

// OnBindFailed indicates an expected call of OnBindFailed.
func (mr *MockStatusTrackingRouteMockRecorder) OnBindFailed(err, gateway interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnBindFailed", reflect.TypeOf((*MockStatusTrackingRoute)(nil).OnBindFailed), err, gateway)
}

// OnBound mocks base method.
func (m *MockStatusTrackingRoute) OnBound(gateway store.Gateway) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnBound", gateway)
}

// OnBound indicates an expected call of OnBound.
func (mr *MockStatusTrackingRouteMockRecorder) OnBound(gateway interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnBound", reflect.TypeOf((*MockStatusTrackingRoute)(nil).OnBound), gateway)
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

// Resolve mocks base method.
func (m *MockStatusTrackingRoute) Resolve(listener store.Listener) *core.ResolvedRoute {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resolve", listener)
	ret0, _ := ret[0].(*core.ResolvedRoute)
	return ret0
}

// Resolve indicates an expected call of Resolve.
func (mr *MockStatusTrackingRouteMockRecorder) Resolve(listener interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockStatusTrackingRoute)(nil).Resolve), listener)
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

// Resolve mocks base method.
func (m *MockRoute) Resolve(listener store.Listener) *core.ResolvedRoute {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Resolve", listener)
	ret0, _ := ret[0].(*core.ResolvedRoute)
	return ret0
}

// Resolve indicates an expected call of Resolve.
func (mr *MockRouteMockRecorder) Resolve(listener interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockRoute)(nil).Resolve), listener)
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

// CanFetchSecrets mocks base method.
func (m *MockStore) CanFetchSecrets(ctx context.Context, id core.GatewayID, secrets []string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanFetchSecrets", ctx, id, secrets)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanFetchSecrets indicates an expected call of CanFetchSecrets.
func (mr *MockStoreMockRecorder) CanFetchSecrets(ctx, id, secrets interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanFetchSecrets", reflect.TypeOf((*MockStore)(nil).CanFetchSecrets), ctx, id, secrets)
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

// GatewayExists mocks base method.
func (m *MockStore) GatewayExists(ctx context.Context, id core.GatewayID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GatewayExists", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GatewayExists indicates an expected call of GatewayExists.
func (mr *MockStoreMockRecorder) GatewayExists(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GatewayExists", reflect.TypeOf((*MockStore)(nil).GatewayExists), ctx, id)
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
