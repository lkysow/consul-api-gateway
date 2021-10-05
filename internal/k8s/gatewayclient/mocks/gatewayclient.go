// Code generated by MockGen. DO NOT EDIT.
// Source: ./gatewayclient.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/hashicorp/consul-api-gateway/pkg/apis/v1alpha1"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/api/core/v1"
	types "k8s.io/apimachinery/pkg/types"
	client "sigs.k8s.io/controller-runtime/pkg/client"
	v1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

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

// CreateOrUpdateDeployment mocks base method.
func (m *MockClient) CreateOrUpdateDeployment(ctx context.Context, deployment *v1.Deployment, mutators ...func() error) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, deployment}
	for _, a := range mutators {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOrUpdateDeployment", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateDeployment indicates an expected call of CreateOrUpdateDeployment.
func (mr *MockClientMockRecorder) CreateOrUpdateDeployment(ctx, deployment interface{}, mutators ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, deployment}, mutators...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateDeployment", reflect.TypeOf((*MockClient)(nil).CreateOrUpdateDeployment), varargs...)
}

// CreateOrUpdateService mocks base method.
func (m *MockClient) CreateOrUpdateService(ctx context.Context, service *v10.Service, mutators ...func() error) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, service}
	for _, a := range mutators {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOrUpdateService", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateService indicates an expected call of CreateOrUpdateService.
func (mr *MockClientMockRecorder) CreateOrUpdateService(ctx, service interface{}, mutators ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, service}, mutators...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateService", reflect.TypeOf((*MockClient)(nil).CreateOrUpdateService), varargs...)
}

// DeleteService mocks base method.
func (m *MockClient) DeleteService(ctx context.Context, service *v10.Service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", ctx, service)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService.
func (mr *MockClientMockRecorder) DeleteService(ctx, service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockClient)(nil).DeleteService), ctx, service)
}

// DeploymentForGateway mocks base method.
func (m *MockClient) DeploymentForGateway(ctx context.Context, gw *v1alpha2.Gateway) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeploymentForGateway", ctx, gw)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeploymentForGateway indicates an expected call of DeploymentForGateway.
func (mr *MockClientMockRecorder) DeploymentForGateway(ctx, gw interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploymentForGateway", reflect.TypeOf((*MockClient)(nil).DeploymentForGateway), ctx, gw)
}

// EnsureFinalizer mocks base method.
func (m *MockClient) EnsureFinalizer(ctx context.Context, object client.Object, finalizer string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureFinalizer", ctx, object, finalizer)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureFinalizer indicates an expected call of EnsureFinalizer.
func (mr *MockClientMockRecorder) EnsureFinalizer(ctx, object, finalizer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureFinalizer", reflect.TypeOf((*MockClient)(nil).EnsureFinalizer), ctx, object, finalizer)
}

// GatewayClassConfigForGatewayClass mocks base method.
func (m *MockClient) GatewayClassConfigForGatewayClass(ctx context.Context, gc *v1alpha2.GatewayClass) (*v1alpha1.GatewayClassConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GatewayClassConfigForGatewayClass", ctx, gc)
	ret0, _ := ret[0].(*v1alpha1.GatewayClassConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GatewayClassConfigForGatewayClass indicates an expected call of GatewayClassConfigForGatewayClass.
func (mr *MockClientMockRecorder) GatewayClassConfigForGatewayClass(ctx, gc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GatewayClassConfigForGatewayClass", reflect.TypeOf((*MockClient)(nil).GatewayClassConfigForGatewayClass), ctx, gc)
}

// GatewayClassConfigInUse mocks base method.
func (m *MockClient) GatewayClassConfigInUse(ctx context.Context, gcc *v1alpha1.GatewayClassConfig) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GatewayClassConfigInUse", ctx, gcc)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GatewayClassConfigInUse indicates an expected call of GatewayClassConfigInUse.
func (mr *MockClientMockRecorder) GatewayClassConfigInUse(ctx, gcc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GatewayClassConfigInUse", reflect.TypeOf((*MockClient)(nil).GatewayClassConfigInUse), ctx, gcc)
}

// GatewayClassForGateway mocks base method.
func (m *MockClient) GatewayClassForGateway(ctx context.Context, gw *v1alpha2.Gateway) (*v1alpha2.GatewayClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GatewayClassForGateway", ctx, gw)
	ret0, _ := ret[0].(*v1alpha2.GatewayClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GatewayClassForGateway indicates an expected call of GatewayClassForGateway.
func (mr *MockClientMockRecorder) GatewayClassForGateway(ctx, gw interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GatewayClassForGateway", reflect.TypeOf((*MockClient)(nil).GatewayClassForGateway), ctx, gw)
}

// GatewayClassInUse mocks base method.
func (m *MockClient) GatewayClassInUse(ctx context.Context, gc *v1alpha2.GatewayClass) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GatewayClassInUse", ctx, gc)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GatewayClassInUse indicates an expected call of GatewayClassInUse.
func (mr *MockClientMockRecorder) GatewayClassInUse(ctx, gc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GatewayClassInUse", reflect.TypeOf((*MockClient)(nil).GatewayClassInUse), ctx, gc)
}

// GetGateway mocks base method.
func (m *MockClient) GetGateway(ctx context.Context, key types.NamespacedName) (*v1alpha2.Gateway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGateway", ctx, key)
	ret0, _ := ret[0].(*v1alpha2.Gateway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGateway indicates an expected call of GetGateway.
func (mr *MockClientMockRecorder) GetGateway(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGateway", reflect.TypeOf((*MockClient)(nil).GetGateway), ctx, key)
}

// GetGatewayClass mocks base method.
func (m *MockClient) GetGatewayClass(ctx context.Context, key types.NamespacedName) (*v1alpha2.GatewayClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGatewayClass", ctx, key)
	ret0, _ := ret[0].(*v1alpha2.GatewayClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGatewayClass indicates an expected call of GetGatewayClass.
func (mr *MockClientMockRecorder) GetGatewayClass(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGatewayClass", reflect.TypeOf((*MockClient)(nil).GetGatewayClass), ctx, key)
}

// GetGatewayClassConfig mocks base method.
func (m *MockClient) GetGatewayClassConfig(ctx context.Context, key types.NamespacedName) (*v1alpha1.GatewayClassConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGatewayClassConfig", ctx, key)
	ret0, _ := ret[0].(*v1alpha1.GatewayClassConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGatewayClassConfig indicates an expected call of GetGatewayClassConfig.
func (mr *MockClientMockRecorder) GetGatewayClassConfig(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGatewayClassConfig", reflect.TypeOf((*MockClient)(nil).GetGatewayClassConfig), ctx, key)
}

// GetHTTPRoute mocks base method.
func (m *MockClient) GetHTTPRoute(ctx context.Context, key types.NamespacedName) (*v1alpha2.HTTPRoute, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHTTPRoute", ctx, key)
	ret0, _ := ret[0].(*v1alpha2.HTTPRoute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHTTPRoute indicates an expected call of GetHTTPRoute.
func (mr *MockClientMockRecorder) GetHTTPRoute(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHTTPRoute", reflect.TypeOf((*MockClient)(nil).GetHTTPRoute), ctx, key)
}

// GetMeshService mocks base method.
func (m *MockClient) GetMeshService(ctx context.Context, key types.NamespacedName) (*v1alpha1.MeshService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeshService", ctx, key)
	ret0, _ := ret[0].(*v1alpha1.MeshService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeshService indicates an expected call of GetMeshService.
func (mr *MockClientMockRecorder) GetMeshService(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeshService", reflect.TypeOf((*MockClient)(nil).GetMeshService), ctx, key)
}

// GetSecret mocks base method.
func (m *MockClient) GetSecret(ctx context.Context, key types.NamespacedName) (*v10.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", ctx, key)
	ret0, _ := ret[0].(*v10.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockClientMockRecorder) GetSecret(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockClient)(nil).GetSecret), ctx, key)
}

// GetService mocks base method.
func (m *MockClient) GetService(ctx context.Context, key types.NamespacedName) (*v10.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", ctx, key)
	ret0, _ := ret[0].(*v10.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService.
func (mr *MockClientMockRecorder) GetService(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockClient)(nil).GetService), ctx, key)
}

// IsManagedRoute mocks base method.
func (m *MockClient) IsManagedRoute(ctx context.Context, spec v1alpha2.CommonRouteSpec, routeNamespace, controllerName string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsManagedRoute", ctx, spec, routeNamespace, controllerName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsManagedRoute indicates an expected call of IsManagedRoute.
func (mr *MockClientMockRecorder) IsManagedRoute(ctx, spec, routeNamespace, controllerName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsManagedRoute", reflect.TypeOf((*MockClient)(nil).IsManagedRoute), ctx, spec, routeNamespace, controllerName)
}

// IsValidGatewayClass mocks base method.
func (m *MockClient) IsValidGatewayClass(ctx context.Context, gc *v1alpha2.GatewayClass) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidGatewayClass", ctx, gc)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsValidGatewayClass indicates an expected call of IsValidGatewayClass.
func (mr *MockClientMockRecorder) IsValidGatewayClass(ctx, gc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidGatewayClass", reflect.TypeOf((*MockClient)(nil).IsValidGatewayClass), ctx, gc)
}

// PodWithLabels mocks base method.
func (m *MockClient) PodWithLabels(ctx context.Context, labels map[string]string) (*v10.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PodWithLabels", ctx, labels)
	ret0, _ := ret[0].(*v10.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PodWithLabels indicates an expected call of PodWithLabels.
func (mr *MockClientMockRecorder) PodWithLabels(ctx, labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PodWithLabels", reflect.TypeOf((*MockClient)(nil).PodWithLabels), ctx, labels)
}

// RemoveFinalizer mocks base method.
func (m *MockClient) RemoveFinalizer(ctx context.Context, object client.Object, finalizer string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFinalizer", ctx, object, finalizer)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveFinalizer indicates an expected call of RemoveFinalizer.
func (mr *MockClientMockRecorder) RemoveFinalizer(ctx, object, finalizer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFinalizer", reflect.TypeOf((*MockClient)(nil).RemoveFinalizer), ctx, object, finalizer)
}

// SetControllerOwnership mocks base method.
func (m *MockClient) SetControllerOwnership(owner, object client.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetControllerOwnership", owner, object)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetControllerOwnership indicates an expected call of SetControllerOwnership.
func (mr *MockClientMockRecorder) SetControllerOwnership(owner, object interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetControllerOwnership", reflect.TypeOf((*MockClient)(nil).SetControllerOwnership), owner, object)
}

// UpdateStatus mocks base method.
func (m *MockClient) UpdateStatus(ctx context.Context, obj client.Object) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", ctx, obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockClientMockRecorder) UpdateStatus(ctx, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockClient)(nil).UpdateStatus), ctx, obj)
}