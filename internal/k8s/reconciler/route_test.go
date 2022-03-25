package reconciler

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	clientMocks "github.com/hashicorp/consul-api-gateway/internal/k8s/gatewayclient/mocks"
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
	require.Equal(t, "tcp-namespace/name", factory.NewRoute(&gw.TCPRoute{
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

func TestRouteResolve(t *testing.T) {
	t.Parallel()

	gateway := &K8sGateway{
		Gateway: &gw.Gateway{
			ObjectMeta: meta.ObjectMeta{
				Name: "expected",
			},
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
	route.bound(gateway.Gateway)

	expected := errors.New("expected")
	client.EXPECT().UpdateStatus(gomock.Any(), inner).Return(expected)
	require.True(t, errors.Is(route.SyncStatus(context.Background()), expected))

	require.NoError(t, route.SyncStatus(context.Background()))
}
