package validators

import (
	"context"
	"errors"

	"github.com/hashicorp/consul-api-gateway/internal/k8s/reconciler/state"
	"github.com/hashicorp/consul-api-gateway/internal/k8s/reconciler/status"
	"github.com/hashicorp/consul-api-gateway/internal/k8s/service"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	gw "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

// all kubernetes routes implement the following two interfaces
type Route interface {
	client.Object
	schema.ObjectKind
}

type RouteValidator struct {
	resolver service.BackendResolver
}

func NewRouteValidator(resolver service.BackendResolver) *RouteValidator {
	return &RouteValidator{
		resolver: resolver,
	}
}

func (r *RouteValidator) Validate(ctx context.Context, route Route) (*state.RouteState, error) {
	state := &state.RouteState{
		ResolutionErrors: service.NewResolutionErrors(),
		References:       make(service.RouteRuleReferenceMap),
		ParentStatuses:   make(status.RouteStatuses),
	}

	switch route := route.(type) {
	case *gw.HTTPRoute:
		return r.validateHTTPRoute(ctx, state, route)
	case *gw.TCPRoute:
		return r.validateTCPRoute(ctx, state, route)
	}

	return nil, nil
}

func (r *RouteValidator) validateHTTPRoute(ctx context.Context, state *state.RouteState, route *gw.HTTPRoute) (*state.RouteState, error) {
	for _, httpRule := range route.Spec.Rules {
		rule := httpRule
		routeRule := service.NewRouteRule(&rule)
		for _, backendRef := range rule.BackendRefs {
			ref := backendRef
			reference, err := r.resolver.Resolve(ctx, route.GetNamespace(), ref.BackendObjectReference)
			if err != nil {
				var resolutionError service.ResolutionError
				if !errors.As(err, &resolutionError) {
					return nil, err
				}
				state.ResolutionErrors.Add(resolutionError)
				continue
			}
			reference.Reference.Set(&ref)
			state.References.Add(routeRule, *reference)
		}
	}
	return state, nil
}

func (r *RouteValidator) validateTCPRoute(ctx context.Context, state *state.RouteState, route *gw.TCPRoute) (*state.RouteState, error) {
	if len(route.Spec.Rules) != 1 {
		err := service.NewResolutionError("a single tcp rule is required")
		state.ResolutionErrors.Add(err)
		return state, nil
	}

	rule := route.Spec.Rules[0]

	if len(rule.BackendRefs) != 1 {
		err := service.NewResolutionError("a single backendRef per tcp rule is required")
		state.ResolutionErrors.Add(err)
		return state, nil
	}

	routeRule := service.NewRouteRule(rule)

	ref := rule.BackendRefs[0]
	reference, err := r.resolver.Resolve(ctx, route.GetNamespace(), ref.BackendObjectReference)
	if err != nil {
		var resolutionError service.ResolutionError
		if !errors.As(err, &resolutionError) {
			return nil, err
		}
		state.ResolutionErrors.Add(resolutionError)
		return state, nil
	}

	reference.Reference.Set(&ref)
	state.References.Add(routeRule, *reference)
	return state, nil
}