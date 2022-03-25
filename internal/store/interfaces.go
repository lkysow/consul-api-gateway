package store

import (
	"context"

	"github.com/hashicorp/consul-api-gateway/internal/core"
)

//go:generate mockgen -source ./interfaces.go -destination ./mocks/interfaces.go -package mocks StatusTrackingGateway,Gateway,RouteTrackingListener,Listener,StatusTrackingRoute,Route,Store

// StatusTrackingGateway is an optional extension
// of Gateway. If supported by a Store, when
// a Gateway is synced to an external location,
// its corresponding callbacks should
// be called.
type StatusTrackingGateway interface {
	Gateway

	TrackSync(ctx context.Context, sync func() (bool, error)) error
}

// Gateway describes a gateway.
type Gateway interface {
	ID() core.GatewayID
	Bind(ctx context.Context, route Route) []string

	// TODO: get rid of these
	Meta() map[string]string
	Listeners() []Listener
}

// ListenerConfig contains the common configuration
// options of a listener.
type ListenerConfig struct {
	Name     string
	Hostname string
	Port     int
	Protocol string
	TLS      core.TLSParams
}

// Listener describes the basic methods of a gateway
// listener.
// TODO: get rid of this interface
type Listener interface {
	ID() string
	Config() ListenerConfig
	IsValid() bool
	RouteRemoved()
}

// StatusTrackingRoute is an optional extension
// of Route. If supported by a Store, when
// a Route is bound or fails to be bound to
// a gateway, its corresponding callbacks should
// be called. At the end of any methods affecting
// the route's binding, SyncStatus should be called.
type StatusTrackingRoute interface {
	Route

	SyncStatus(ctx context.Context) error
	OnGatewayRemoved(gateway Gateway)
}

// Route should be implemented by all route
// source integrations
type Route interface {
	ID() string

	// TODO: move this to the gateway as
	// Resolve(routes []Route) *core.ResolvedGateway
	Resolve(listener Listener) *core.ResolvedRoute
}

// Store is used for persisting and querying gateways and routes
type Store interface {
	// TODO: make these part of a Backend interface
	GatewayExists(ctx context.Context, id core.GatewayID) (bool, error)
	DeleteGateway(ctx context.Context, id core.GatewayID) error
	UpsertGateway(ctx context.Context, gateway Gateway, updateConditionFn func(current Gateway) bool) error
	DeleteRoute(ctx context.Context, id string) error
	UpsertRoute(ctx context.Context, route Route, updateConditionFn func(current Route) bool) error

	// TODO: move this into a concrete implementation of store
	Sync(ctx context.Context) error

	// TODO: move this to the gateway
	CanFetchSecrets(ctx context.Context, id core.GatewayID, secrets []string) (bool, error)
}
