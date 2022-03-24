package reconciler

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	clientMocks "github.com/hashicorp/consul-api-gateway/internal/k8s/gatewayclient/mocks"
	"github.com/hashicorp/consul-api-gateway/internal/k8s/service"
	"github.com/hashicorp/consul-api-gateway/internal/k8s/service/mocks"
	"github.com/hashicorp/consul-api-gateway/internal/store"
	storeMocks "github.com/hashicorp/consul-api-gateway/internal/store/mocks"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/require"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	gw "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

func TestRouteID(t *testing.T) {
	t.Parallel()

	factory := NewFactory(FactoryConfig{
		Logger: hclog.NewNullLogger(),
	})

	meta := meta.ObjectMeta{
		Name:      "name",
		Namespace: "namespace",
	}

	require.Equal(t, "http-namespace/name", factory.NewRoute(&gw.HTTPRoute{
		ObjectMeta: meta,
	}).ID())
	require.Equal(t, "udp-namespace/name", factory.NewRoute(&gw.UDPRoute{
		ObjectMeta: meta,
	}).ID())
	require.Equal(t, "tcp-namespace/name", factory.NewRoute(&gw.TCPRoute{
		ObjectMeta: meta,
	}).ID())
	require.Equal(t, "tls-namespace/name", factory.NewRoute(&gw.TLSRoute{
		ObjectMeta: meta,
	}).ID())
	require.Equal(t, "", factory.NewRoute(&core.Pod{
		ObjectMeta: meta,
	}).ID())
}

func TestRouteCommonRouteSpec(t *testing.T) {
	t.Parallel()

	factory := NewFactory(FactoryConfig{
		Logger: hclog.NewNullLogger(),
	})

	expected := gw.CommonRouteSpec{
		ParentRefs: []gw.ParentRef{{
			Name: "expected",
		}},
	}

	require.Equal(t, expected, factory.NewRoute(&gw.HTTPRoute{
		Spec: gw.HTTPRouteSpec{
			CommonRouteSpec: expected,
		},
	}).CommonRouteSpec())
	require.Equal(t, expected, factory.NewRoute(&gw.UDPRoute{
		Spec: gw.UDPRouteSpec{
			CommonRouteSpec: expected,
		},
	}).CommonRouteSpec())
	require.Equal(t, expected, factory.NewRoute(&gw.TCPRoute{
		Spec: gw.TCPRouteSpec{
			CommonRouteSpec: expected,
		},
	}).CommonRouteSpec())
	require.Equal(t, expected, factory.NewRoute(&gw.TLSRoute{
		Spec: gw.TLSRouteSpec{
			CommonRouteSpec: expected,
		},
	}).CommonRouteSpec())
	require.Equal(t, gw.CommonRouteSpec{}, factory.NewRoute(&core.Pod{}).CommonRouteSpec())
}

func TestRouteFilterParentStatuses(t *testing.T) {
	t.Parallel()

	factory := NewFactory(FactoryConfig{
		Logger:         hclog.NewNullLogger(),
		ControllerName: "expected",
	})

	route := factory.NewRoute(&gw.HTTPRoute{
		Spec: gw.HTTPRouteSpec{
			CommonRouteSpec: gw.CommonRouteSpec{
				ParentRefs: []gw.ParentRef{{
					Name: "expected",
				}, {
					Name: "other",
				}},
			},
		},
		Status: gw.HTTPRouteStatus{
			RouteStatus: gw.RouteStatus{
				Parents: []gw.RouteParentStatus{{
					ParentRef: gw.ParentRef{
						Name: "expected",
					},
					ControllerName: "expected",
				}, {
					ParentRef: gw.ParentRef{
						Name: "expected",
					},
					ControllerName: "other",
				}, {
					ParentRef: gw.ParentRef{
						Name: "other",
					},
					ControllerName: "other",
				}},
			},
		},
	})

	route.OnBound(factory.NewGateway(NewGatewayConfig{
		Gateway: &gw.Gateway{
			ObjectMeta: meta.ObjectMeta{
				Name: "expected",
			},
		},
	}))

	statuses := route.FilterParentStatuses()
	require.Len(t, statuses, 2)
	require.Equal(t, "expected", string(statuses[0].ParentRef.Name))
	require.Equal(t, "other", string(statuses[0].ControllerName))
	require.Equal(t, "other", string(statuses[1].ParentRef.Name))
	require.Equal(t, "other", string(statuses[1].ControllerName))
}

func TestRouteMergedStatusAndBinding(t *testing.T) {
	t.Parallel()

	factory := NewFactory(FactoryConfig{
		Logger:         hclog.NewNullLogger(),
		ControllerName: "expected",
	})

	gateway := factory.NewGateway(NewGatewayConfig{
		Gateway: &gw.Gateway{
			ObjectMeta: meta.ObjectMeta{
				Name: "expected",
			},
		},
	})
	inner := &gw.TLSRoute{
		Spec: gw.TLSRouteSpec{
			CommonRouteSpec: gw.CommonRouteSpec{
				ParentRefs: []gw.ParentRef{{
					Name: "expected",
				}, {
					Name: "other",
				}},
			},
		},
		Status: gw.TLSRouteStatus{
			RouteStatus: gw.RouteStatus{
				Parents: []gw.RouteParentStatus{{
					ParentRef: gw.ParentRef{
						Name: "expected",
					},
					ControllerName: "expected",
				}, {
					ParentRef: gw.ParentRef{
						Name: "expected",
					},
					ControllerName: "other",
				}, {
					ParentRef: gw.ParentRef{
						Name: "other",
					},
					ControllerName: "other",
				}},
			},
		},
	}
	route := factory.NewRoute(inner)
	route.OnBound(gateway)

	statuses := route.MergedStatus().Parents
	require.Len(t, statuses, 3)
	require.Equal(t, "expected", string(statuses[0].ParentRef.Name))
	require.Equal(t, "expected", string(statuses[0].ControllerName))
	require.Len(t, statuses[0].Conditions, 2)
	require.Equal(t, "Route accepted.", statuses[0].Conditions[0].Message)
	require.Equal(t, "expected", string(statuses[1].ParentRef.Name))
	require.Equal(t, "other", string(statuses[1].ControllerName))
	require.Equal(t, "other", string(statuses[2].ParentRef.Name))
	require.Equal(t, "other", string(statuses[2].ControllerName))

	route.OnBindFailed(errors.New("expected"), gateway)

	statuses = route.MergedStatus().Parents
	require.Len(t, statuses, 3)
	require.Equal(t, "expected", string(statuses[0].ParentRef.Name))
	require.Equal(t, "expected", string(statuses[0].ControllerName))
	require.Equal(t, "expected", statuses[0].Conditions[0].Message)
	require.Equal(t, RouteConditionReasonBindError, statuses[0].Conditions[0].Reason)

	route.OnBindFailed(NewBindErrorHostnameMismatch("expected"), gateway)

	statuses = route.MergedStatus().Parents
	require.Len(t, statuses, 3)
	require.Equal(t, "expected", string(statuses[0].ParentRef.Name))
	require.Equal(t, "expected", string(statuses[0].ControllerName))
	require.Equal(t, "expected", statuses[0].Conditions[0].Message)
	require.Equal(t, RouteConditionReasonListenerHostnameMismatch, statuses[0].Conditions[0].Reason)

	route.OnBindFailed(NewBindErrorListenerNamespacePolicy("expected"), gateway)

	statuses = route.MergedStatus().Parents
	require.Len(t, statuses, 3)
	require.Equal(t, "expected", string(statuses[0].ParentRef.Name))
	require.Equal(t, "expected", string(statuses[0].ControllerName))
	require.Equal(t, "expected", statuses[0].Conditions[0].Message)
	require.Equal(t, RouteConditionReasonListenerNamespacePolicy, statuses[0].Conditions[0].Reason)

	route.OnBindFailed(NewBindErrorRouteKind("expected"), gateway)

	statuses = route.MergedStatus().Parents
	require.Len(t, statuses, 3)
	require.Equal(t, "expected", string(statuses[0].ParentRef.Name))
	require.Equal(t, "expected", string(statuses[0].ControllerName))
	require.Equal(t, "expected", statuses[0].Conditions[0].Message)
	require.Equal(t, RouteConditionReasonInvalidRouteKind, statuses[0].Conditions[0].Reason)

	route.OnBound(gateway)

	statuses = route.MergedStatus().Parents
	require.Len(t, statuses, 3)
	require.Equal(t, "expected", string(statuses[0].ParentRef.Name))
	require.Equal(t, "expected", string(statuses[0].ControllerName))
	require.Equal(t, "Route accepted.", statuses[0].Conditions[0].Message)

	route.OnGatewayRemoved(gateway)

	statuses = route.MergedStatus().Parents
	require.Len(t, statuses, 2)
	require.Equal(t, "expected", string(statuses[0].ParentRef.Name))
	require.Equal(t, "other", string(statuses[0].ControllerName))
	require.Equal(t, "other", string(statuses[1].ParentRef.Name))
	require.Equal(t, "other", string(statuses[1].ControllerName))

	// check creating a status on bind failure when it's not there
	route = factory.NewRoute(inner)

	route.OnBindFailed(errors.New("expected"), gateway)

	statuses = route.MergedStatus().Parents
	require.Len(t, statuses, 3)
	require.Equal(t, "expected", string(statuses[0].ParentRef.Name))
	require.Equal(t, "expected", string(statuses[0].ControllerName))
	require.Equal(t, "expected", statuses[0].Conditions[0].Message)
	require.Equal(t, RouteConditionReasonBindError, statuses[0].Conditions[0].Reason)

	// check binding for non-existent route
	gateway = factory.NewGateway(NewGatewayConfig{
		Gateway: &gw.Gateway{
			ObjectMeta: meta.ObjectMeta{
				Name: "nothing",
			},
		},
	})
	statuses = route.MergedStatus().Parents
	require.Len(t, statuses, 3)
	route.OnBound(gateway)
	statuses = route.MergedStatus().Parents
	require.Len(t, statuses, 3)
	route.OnBindFailed(errors.New("expected"), gateway)
	statuses = route.MergedStatus().Parents
	require.Len(t, statuses, 3)
}

func TestRouteNeedsStatusUpdate(t *testing.T) {
	t.Parallel()

	factory := NewFactory(FactoryConfig{
		Logger:         hclog.NewNullLogger(),
		ControllerName: "expected",
	})

	route := factory.NewRoute(&gw.TCPRoute{
		Spec: gw.TCPRouteSpec{
			CommonRouteSpec: gw.CommonRouteSpec{
				ParentRefs: []gw.ParentRef{{
					Name: "expected",
				}, {
					Name: "other",
				}},
			},
		},
		Status: gw.TCPRouteStatus{
			RouteStatus: gw.RouteStatus{
				Parents: []gw.RouteParentStatus{{
					ParentRef: gw.ParentRef{
						Name: "expected",
					},
					ControllerName: "expected",
				}, {
					ParentRef: gw.ParentRef{
						Name: "expected",
					},
					ControllerName: "other",
				}, {
					ParentRef: gw.ParentRef{
						Name: "other",
					},
					ControllerName: "other",
				}},
			},
		},
	})
	route.SetStatus(route.MergedStatus())

	require.False(t, route.NeedsStatusUpdate())

	route.OnBound(factory.NewGateway(NewGatewayConfig{
		Gateway: &gw.Gateway{
			ObjectMeta: meta.ObjectMeta{
				Name: "expected",
			},
		},
	}))

	require.True(t, route.NeedsStatusUpdate())

	route.SetStatus(route.MergedStatus())

	require.False(t, route.NeedsStatusUpdate())
}

func TestRouteSetStatus(t *testing.T) {
	t.Parallel()

	factory := NewFactory(FactoryConfig{
		Logger: hclog.NewNullLogger(),
	})

	expected := gw.RouteStatus{
		Parents: []gw.RouteParentStatus{{
			ParentRef: gw.ParentRef{
				Name: "expected",
			},
		}},
	}

	httpRoute := &gw.HTTPRoute{}
	route := factory.NewRoute(httpRoute)
	route.SetStatus(expected)
	require.Equal(t, expected, httpRoute.Status.RouteStatus)
	require.Equal(t, expected, route.routeStatus())

	tcpRoute := &gw.TCPRoute{}
	route = factory.NewRoute(tcpRoute)
	route.SetStatus(expected)
	require.Equal(t, expected, tcpRoute.Status.RouteStatus)
	require.Equal(t, expected, route.routeStatus())

	tlsRoute := &gw.TLSRoute{}
	route = factory.NewRoute(tlsRoute)
	route.SetStatus(expected)
	require.Equal(t, expected, tlsRoute.Status.RouteStatus)
	require.Equal(t, expected, route.routeStatus())

	udpRoute := &gw.UDPRoute{}
	route = factory.NewRoute(udpRoute)
	route.SetStatus(expected)
	require.Equal(t, expected, udpRoute.Status.RouteStatus)
	require.Equal(t, expected, route.routeStatus())

	route = factory.NewRoute(&core.Pod{})
	route.SetStatus(expected)
	require.Equal(t, gw.RouteStatus{}, route.routeStatus())
}

func TestRouteParents(t *testing.T) {
	t.Parallel()

	factory := NewFactory(FactoryConfig{
		Logger: hclog.NewNullLogger(),
	})

	expected := gw.CommonRouteSpec{
		ParentRefs: []gw.ParentRef{{
			Name: "expected",
		}},
	}

	parents := factory.NewRoute(&gw.HTTPRoute{Spec: gw.HTTPRouteSpec{CommonRouteSpec: expected}}).Parents()
	require.Equal(t, expected.ParentRefs, parents)

	parents = factory.NewRoute(&gw.TCPRoute{Spec: gw.TCPRouteSpec{CommonRouteSpec: expected}}).Parents()
	require.Equal(t, expected.ParentRefs, parents)

	parents = factory.NewRoute(&gw.TLSRoute{Spec: gw.TLSRouteSpec{CommonRouteSpec: expected}}).Parents()
	require.Equal(t, expected.ParentRefs, parents)

	parents = factory.NewRoute(&gw.UDPRoute{Spec: gw.UDPRouteSpec{CommonRouteSpec: expected}}).Parents()
	require.Equal(t, expected.ParentRefs, parents)

	require.Nil(t, factory.NewRoute(&core.Pod{}).Parents())
}

func TestRouteMatchesHostname(t *testing.T) {
	t.Parallel()

	hostname := gw.Hostname("domain.test")

	factory := NewFactory(FactoryConfig{
		Logger: hclog.NewNullLogger(),
	})

	require.True(t, factory.NewRoute(&gw.HTTPRoute{
		Spec: gw.HTTPRouteSpec{
			Hostnames: []gw.Hostname{"*"},
		},
	}).MatchesHostname(&hostname))

	require.False(t, factory.NewRoute(&gw.HTTPRoute{
		Spec: gw.HTTPRouteSpec{
			Hostnames: []gw.Hostname{"other.text"},
		},
	}).MatchesHostname(&hostname))

	// check where the underlying route doesn't implement
	// a matching routine
	require.True(t, factory.NewRoute(&gw.TCPRoute{}).MatchesHostname(&hostname))
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

	require.NoError(t, factory.NewRoute(&core.Pod{}).Validate(context.Background()))
	require.True(t, factory.NewRoute(&gw.HTTPRoute{}).IsValid())

	reference := gw.BackendObjectReference{
		Name: "expected",
	}
	resolved := &service.ResolvedReference{
		Type:      service.ConsulServiceReference,
		Reference: &service.BackendReference{},
	}

	resolver.EXPECT().Resolve(gomock.Any(), gomock.Any(), reference).Return(resolved, nil)

	route := factory.NewRoute(&gw.HTTPRoute{
		Spec: gw.HTTPRouteSpec{
			Rules: []gw.HTTPRouteRule{{
				BackendRefs: []gw.HTTPBackendRef{{
					BackendRef: gw.BackendRef{
						BackendObjectReference: reference,
					},
				}},
			}},
		},
	})
	require.NoError(t, route.Validate(context.Background()))
	require.True(t, route.IsValid())

	expected := errors.New("expected")
	resolver.EXPECT().Resolve(gomock.Any(), gomock.Any(), reference).Return(nil, expected)
	require.Equal(t, expected, route.Validate(context.Background()))

	resolver.EXPECT().Resolve(gomock.Any(), gomock.Any(), reference).Return(nil, service.NewK8sResolutionError("error"))
	require.NoError(t, route.Validate(context.Background()))
	require.False(t, route.IsValid())
}

func TestRouteResolve(t *testing.T) {
	t.Parallel()

	gateway := &gw.Gateway{
		ObjectMeta: meta.ObjectMeta{
			Name: "expected",
		},
	}
	listener := gw.Listener{}

	factory := NewFactory(FactoryConfig{
		Logger: hclog.NewNullLogger(),
	})

	require.Nil(t, factory.NewRoute(&gw.HTTPRoute{}).Resolve(nil))

	require.Nil(t, factory.NewRoute(&core.Pod{}).Resolve(NewK8sListener(gateway, listener, K8sListenerConfig{
		Logger: hclog.NewNullLogger(),
	})))

	require.NotNil(t, factory.NewRoute(&gw.HTTPRoute{}).Resolve(NewK8sListener(gateway, listener, K8sListenerConfig{
		Logger: hclog.NewNullLogger(),
	})))
}

func TestRouteCompare(t *testing.T) {
	t.Parallel()

	factory := NewFactory(FactoryConfig{
		Logger: hclog.NewNullLogger(),
	})

	// invalid route comparison
	route := factory.NewRoute(&core.Pod{})
	other := factory.NewRoute(&core.Pod{})

	require.Equal(t, store.CompareResultNotEqual, route.Compare(route))
	require.Equal(t, store.CompareResultInvalid, route.Compare(nil))
	route = nil
	require.Equal(t, store.CompareResultNotEqual, route.Compare(other))

	// http route comparison
	route = factory.NewRoute(&gw.HTTPRoute{})
	other = factory.NewRoute(&gw.HTTPRoute{})
	require.Equal(t, store.CompareResultEqual, route.Compare(other))
	other.resolutionErrors.Add(service.NewConsulResolutionError("error"))
	require.Equal(t, store.CompareResultNotEqual, route.Compare(other))
	route = factory.NewRoute(&gw.HTTPRoute{
		ObjectMeta: meta.ObjectMeta{
			ResourceVersion: "1",
		},
	})
	require.Equal(t, store.CompareResultNewer, route.Compare(other))

	route = factory.NewRoute(&gw.HTTPRoute{})
	other = factory.NewRoute(&gw.TCPRoute{})
	require.Equal(t, store.CompareResultNotEqual, route.Compare(other))

	// tcp route comparison
	route = factory.NewRoute(&gw.TCPRoute{})
	other = factory.NewRoute(&gw.TCPRoute{})
	require.Equal(t, store.CompareResultEqual, route.Compare(other))
	other = factory.NewRoute(&gw.HTTPRoute{})
	require.Equal(t, store.CompareResultNotEqual, route.Compare(other))

	// tls route comparison
	route = factory.NewRoute(&gw.TLSRoute{})
	other = factory.NewRoute(&gw.TLSRoute{})
	require.Equal(t, store.CompareResultEqual, route.Compare(other))
	other = factory.NewRoute(&gw.HTTPRoute{})
	require.Equal(t, store.CompareResultNotEqual, route.Compare(other))

	// udp route comparison
	route = factory.NewRoute(&gw.UDPRoute{})
	other = factory.NewRoute(&gw.UDPRoute{})
	require.Equal(t, store.CompareResultEqual, route.Compare(other))
	other = factory.NewRoute(&gw.HTTPRoute{})
	require.Equal(t, store.CompareResultNotEqual, route.Compare(other))

	// mismatched types
	require.Equal(t, store.CompareResultInvalid, route.Compare(storeMocks.NewMockRoute(nil)))
}

func TestRouteSyncStatus(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := clientMocks.NewMockClient(ctrl)

	factory := NewFactory(FactoryConfig{
		Logger:         hclog.NewNullLogger(),
		Client:         client,
		ControllerName: "expected",
	})

	gateway := factory.NewGateway(NewGatewayConfig{
		Gateway: &gw.Gateway{
			ObjectMeta: meta.ObjectMeta{
				Name: "expected",
			},
		},
	})
	inner := &gw.TLSRoute{
		Spec: gw.TLSRouteSpec{
			CommonRouteSpec: gw.CommonRouteSpec{
				ParentRefs: []gw.ParentRef{{
					Name: "expected",
				}, {
					Name: "other",
				}},
			},
		},
		Status: gw.TLSRouteStatus{
			RouteStatus: gw.RouteStatus{
				Parents: []gw.RouteParentStatus{{
					ParentRef: gw.ParentRef{
						Name: "expected",
					},
					ControllerName: "expected",
				}, {
					ParentRef: gw.ParentRef{
						Name: "expected",
					},
					ControllerName: "other",
				}, {
					ParentRef: gw.ParentRef{
						Name: "other",
					},
					ControllerName: "other",
				}},
			},
		},
	}

	logger := hclog.New(&hclog.LoggerOptions{
		Output: io.Discard,
	})
	logger.SetLevel(hclog.Trace)
	route := factory.NewRoute(inner)
	route.OnBound(gateway)

	expected := errors.New("expected")
	client.EXPECT().UpdateStatus(gomock.Any(), inner).Return(expected)
	require.True(t, errors.Is(route.SyncStatus(context.Background()), expected))

	client.EXPECT().UpdateStatus(gomock.Any(), inner)
	require.NoError(t, route.SyncStatus(context.Background()))

	// sync again, no status update called
	require.NoError(t, route.SyncStatus(context.Background()))
}
