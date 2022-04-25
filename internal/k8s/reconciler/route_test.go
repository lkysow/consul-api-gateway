package reconciler

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/require"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	gwv1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	gwv1beta1 "sigs.k8s.io/gateway-api/apis/v1beta1"

	clientMocks "github.com/hashicorp/consul-api-gateway/internal/k8s/gatewayclient/mocks"
	rstate "github.com/hashicorp/consul-api-gateway/internal/k8s/reconciler/state"
	"github.com/hashicorp/consul-api-gateway/internal/k8s/service"
	"github.com/hashicorp/consul-api-gateway/internal/k8s/service/mocks"
)

func TestRouteID(t *testing.T) {
	t.Parallel()

	factory := NewFactory(FactoryConfig{
		Logger: hclog.NewNullLogger(),
	})

	config := rstate.NewRouteState()

	meta := meta.ObjectMeta{
		Name:      "name",
		Namespace: "namespace",
	}

	require.Equal(t, "http-namespace/name", factory.NewRoute(&gwv1alpha2.HTTPRoute{
		ObjectMeta: meta,
	}, config).ID())
	require.Equal(t, "tcp-namespace/name", factory.NewRoute(&gwv1alpha2.TCPRoute{
		ObjectMeta: meta,
	}, config).ID())
	require.Equal(t, "", factory.NewRoute(&core.Pod{
		ObjectMeta: meta,
	}, config).ID())
}

func TestRouteCommonRouteSpec(t *testing.T) {
	t.Parallel()

	factory := NewFactory(FactoryConfig{
		Logger: hclog.NewNullLogger(),
	})

	config := rstate.NewRouteState()

	expected := gwv1alpha2.CommonRouteSpec{
		ParentRefs: []gwv1alpha2.ParentReference{{
			Name: "expected",
		}},
	}

	require.Equal(t, expected, factory.NewRoute(&gwv1alpha2.HTTPRoute{
		Spec: gwv1alpha2.HTTPRouteSpec{
			CommonRouteSpec: expected,
		},
	}, config).CommonRouteSpec())
	require.Equal(t, expected, factory.NewRoute(&gwv1alpha2.TCPRoute{
		Spec: gwv1alpha2.TCPRouteSpec{
			CommonRouteSpec: expected,
		},
	}, config).CommonRouteSpec())
	require.Equal(t, gwv1alpha2.CommonRouteSpec{}, factory.NewRoute(&core.Pod{}, config).CommonRouteSpec())
}

func TestRouteSetStatus(t *testing.T) {
	t.Parallel()

	factory := NewFactory(FactoryConfig{
		Logger: hclog.NewNullLogger(),
	})

	config := rstate.NewRouteState()

	expected := gwv1alpha2.RouteStatus{
		Parents: []gwv1alpha2.RouteParentStatus{{
			ParentRef: gwv1alpha2.ParentReference{
				Name: "expected",
			},
		}},
	}

	httpRoute := &gwv1alpha2.HTTPRoute{}
	route := factory.NewRoute(httpRoute, config)
	route.SetStatus(expected)
	require.Equal(t, expected, httpRoute.Status.RouteStatus)
	require.Equal(t, expected, route.routeStatus())

	tcpRoute := &gwv1alpha2.TCPRoute{}
	route = factory.NewRoute(tcpRoute, config)
	route.SetStatus(expected)
	require.Equal(t, expected, tcpRoute.Status.RouteStatus)
	require.Equal(t, expected, route.routeStatus())

	route = factory.NewRoute(&core.Pod{}, config)
	route.SetStatus(expected)
	require.Equal(t, gwv1alpha2.RouteStatus{}, route.routeStatus())
}

func TestRouteParents(t *testing.T) {
	t.Parallel()

	factory := NewFactory(FactoryConfig{
		Logger: hclog.NewNullLogger(),
	})

	config := rstate.NewRouteState()

	expected := gwv1alpha2.CommonRouteSpec{
		ParentRefs: []gwv1alpha2.ParentReference{{
			Name: "expected",
		}},
	}

	parents := factory.NewRoute(&gwv1alpha2.HTTPRoute{Spec: gwv1alpha2.HTTPRouteSpec{CommonRouteSpec: expected}}, config).Parents()
	require.Equal(t, expected.ParentRefs, parents)

	parents = factory.NewRoute(&gwv1alpha2.TCPRoute{Spec: gwv1alpha2.TCPRouteSpec{CommonRouteSpec: expected}}, config).Parents()
	require.Equal(t, expected.ParentRefs, parents)

	require.Nil(t, factory.NewRoute(&core.Pod{}, config).Parents())
}

func TestRouteMatchesHostname(t *testing.T) {
	t.Parallel()

	factory := NewFactory(FactoryConfig{
		Logger: hclog.NewNullLogger(),
	})

	hostname := gwv1beta1.Hostname("domain.test")

	require.True(t, factory.NewRoute(&gwv1alpha2.HTTPRoute{
		Spec: gwv1alpha2.HTTPRouteSpec{
			Hostnames: []gwv1alpha2.Hostname{"*"},
		},
	}, rstate.NewRouteState()).MatchesHostname(&hostname))

	require.False(t, factory.NewRoute(&gwv1alpha2.HTTPRoute{
		Spec: gwv1alpha2.HTTPRouteSpec{
			Hostnames: []gwv1alpha2.Hostname{"other.text"},
		},
	}, rstate.NewRouteState()).MatchesHostname(&hostname))

	// check where the underlying route doesn't implement
	// a matching routine
	require.True(t, factory.NewRoute(&gwv1alpha2.TCPRoute{}, rstate.NewRouteState()).MatchesHostname(&hostname))
}

func TestRouteValidate(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	resolver := mocks.NewMockBackendResolver(ctrl)

	factory := NewFactory(FactoryConfig{
		Logger:   hclog.NewNullLogger(),
		Resolver: resolver,
	})

	require.NoError(t, factory.NewRoute(&core.Pod{}, rstate.NewRouteState()).Validate(context.Background()))

	require.True(t, factory.NewRoute(&gwv1alpha2.HTTPRoute{}, rstate.NewRouteState()).IsValid())

	reference := gwv1alpha2.BackendObjectReference{
		Name: "expected",
	}
	resolved := &service.ResolvedReference{
		Type:      service.ConsulServiceReference,
		Reference: &service.BackendReference{},
	}

	resolver.EXPECT().Resolve(gomock.Any(), gomock.Any(), reference).Return(resolved, nil)

	route := factory.NewRoute(&gwv1alpha2.HTTPRoute{
		Spec: gwv1alpha2.HTTPRouteSpec{
			Rules: []gwv1alpha2.HTTPRouteRule{{
				BackendRefs: []gwv1alpha2.HTTPBackendRef{{
					BackendRef: gwv1alpha2.BackendRef{
						BackendObjectReference: reference,
					},
				}},
			}},
		},
	}, rstate.NewRouteState())
	require.NoError(t, route.Validate(context.Background()))
	require.True(t, route.IsValid())

	expected := errors.New("expected")
	resolver.EXPECT().Resolve(gomock.Any(), gomock.Any(), reference).Return(nil, expected)
	require.Equal(t, expected, route.Validate(context.Background()))

	resolver.EXPECT().Resolve(gomock.Any(), gomock.Any(), reference).Return(nil, service.NewK8sResolutionError("error"))
	require.NoError(t, route.Validate(context.Background()))
	require.False(t, route.IsValid())
}

func TestRouteValidateDontAllowCrossNamespace(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	resolver := mocks.NewMockBackendResolver(ctrl)
	client := clientMocks.NewMockClient(ctrl)

	factory := NewFactory(FactoryConfig{
		Logger:   hclog.NewNullLogger(),
		Resolver: resolver,
		Client:   client,
	})

	//set up backend ref with a different namespace
	namespace := gwv1alpha2.Namespace("test")
	route := factory.NewRoute(&gwv1alpha2.HTTPRoute{
		Spec: gwv1alpha2.HTTPRouteSpec{
			Rules: []gwv1alpha2.HTTPRouteRule{{
				BackendRefs: []gwv1alpha2.HTTPBackendRef{{
					BackendRef: gwv1alpha2.BackendRef{
						BackendObjectReference: gwv1alpha2.BackendObjectReference{
							Name:      "expected",
							Namespace: &namespace,
						},
					},
				}},
			}},
		},
	}, rstate.NewRouteState())

	client.EXPECT().
		GetReferenceGrantsInNamespace(gomock.Any(), gomock.Any()).
		Return([]gwv1alpha2.ReferenceGrant{
			{
				Spec: gwv1alpha2.ReferenceGrantSpec{
					From: []gwv1alpha2.ReferenceGrantFrom{},
					To:   []gwv1alpha2.ReferenceGrantTo{},
				},
			},
		}, nil)

	// FUTURE Assert appropriate status set on route and !route.IsValid() once ReferencePolicy requirement is enforced
	_ = route.Validate(context.Background())
}

// TestRouteValidateAllowCrossNamespaceWithReferencePolicy verifies that a cross-namespace
// route + backend combination is allowed if an applicable ReferencePolicy is found.
func TestRouteValidateAllowCrossNamespaceWithReferencePolicy(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	resolver := mocks.NewMockBackendResolver(ctrl)
	client := clientMocks.NewMockClient(ctrl)

	factory := NewFactory(FactoryConfig{
		Logger:   hclog.NewNullLogger(),
		Resolver: resolver,
		Client:   client,
	})

	//set up backend ref with a different namespace
	backendGroup := gwv1alpha2.Group("")
	backendKind := gwv1alpha2.Kind("Service")
	backendNamespace := gwv1alpha2.Namespace("namespace2")
	backendName := gwv1alpha2.ObjectName("backend2")
	route := factory.NewRoute(&gwv1alpha2.HTTPRoute{
		ObjectMeta: meta.ObjectMeta{Namespace: "namespace1"},
		TypeMeta:   meta.TypeMeta{APIVersion: "gateway.networking.k8s.io/v1alpha2", Kind: "HTTPRoute"},
		Spec: gwv1alpha2.HTTPRouteSpec{
			Rules: []gwv1alpha2.HTTPRouteRule{{
				BackendRefs: []gwv1alpha2.HTTPBackendRef{{
					BackendRef: gwv1alpha2.BackendRef{
						BackendObjectReference: gwv1alpha2.BackendObjectReference{
							Group:     &backendGroup,
							Kind:      &backendKind,
							Name:      backendName,
							Namespace: &backendNamespace,
						},
					},
				}},
			}},
		},
	}, rstate.NewRouteState())

	referenceGrant := gwv1alpha2.ReferenceGrant{
		TypeMeta:   meta.TypeMeta{},
		ObjectMeta: meta.ObjectMeta{Namespace: "namespace2"},
		Spec: gwv1alpha2.ReferenceGrantSpec{
			From: []gwv1alpha2.ReferenceGrantFrom{{
				Group:     "gateway.networking.k8s.io",
				Kind:      "HTTPRoute",
				Namespace: "namespace1",
			}},
			To: []gwv1alpha2.ReferenceGrantTo{{
				Group: "",
				Kind:  "Service",
				Name:  &backendName,
			}},
		},
	}

	client.EXPECT().
		GetReferenceGrantsInNamespace(gomock.Any(), gomock.Any()).
		Return([]gwv1alpha2.ReferenceGrant{referenceGrant}, nil)

	resolver.EXPECT().
		Resolve(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&service.ResolvedReference{Type: service.ConsulServiceReference, Reference: &service.BackendReference{}}, nil)

	require.NoError(t, route.Validate(context.Background()))
}

func TestRouteResolve(t *testing.T) {
	t.Parallel()

	factory := NewFactory(FactoryConfig{
		Logger: hclog.NewNullLogger(),
	})

	gateway := &gwv1beta1.Gateway{
		ObjectMeta: meta.ObjectMeta{
			Name: "expected",
		},
	}
	listener := gwv1beta1.Listener{}

	require.Nil(t, factory.NewRoute(&gwv1alpha2.HTTPRoute{}, rstate.NewRouteState()).Resolve(nil))

	require.Nil(t, factory.NewRoute(&core.Pod{}, rstate.NewRouteState()).Resolve(NewK8sListener(gateway, listener, K8sListenerConfig{
		Logger: hclog.NewNullLogger(),
	})))

	require.NotNil(t, factory.NewRoute(&gwv1alpha2.HTTPRoute{}, rstate.NewRouteState()).Resolve(NewK8sListener(gateway, listener, K8sListenerConfig{
		Logger: hclog.NewNullLogger(),
	})))
}

func TestRouteSyncStatus(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := clientMocks.NewMockClient(ctrl)

	factory := NewFactory(FactoryConfig{
		Logger: hclog.NewNullLogger(),
		Client: client,
	})

	g := &gwv1beta1.Gateway{
		ObjectMeta: meta.ObjectMeta{
			Name: "expected",
		},
	}
	gateway := factory.NewGateway(NewGatewayConfig{
		Gateway: g,
		State:   rstate.InitialGatewayState(g),
	})
	inner := &gwv1alpha2.TCPRoute{
		Spec: gwv1alpha2.TCPRouteSpec{
			CommonRouteSpec: gwv1alpha2.CommonRouteSpec{
				ParentRefs: []gwv1alpha2.ParentReference{{
					Name: "expected",
				}, {
					Name: "other",
				}},
			},
		},
		Status: gwv1alpha2.TCPRouteStatus{
			RouteStatus: gwv1alpha2.RouteStatus{
				Parents: []gwv1alpha2.RouteParentStatus{{
					ParentRef: gwv1alpha2.ParentReference{
						Name: "expected",
					},
					ControllerName: "expected",
				}, {
					ParentRef: gwv1alpha2.ParentReference{
						Name: "expected",
					},
					ControllerName: "other",
				}, {
					ParentRef: gwv1alpha2.ParentReference{
						Name: "other",
					},
					ControllerName: "other",
				}},
			},
		},
	}
	route := factory.NewRoute(inner, rstate.NewRouteState())
	route.OnBound(gateway)

	expected := errors.New("expected")
	client.EXPECT().UpdateStatus(gomock.Any(), inner).Return(expected)
	require.True(t, errors.Is(route.SyncStatus(context.Background()), expected))

	require.NoError(t, route.SyncStatus(context.Background()))
}
