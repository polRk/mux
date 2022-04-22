package mux

import (
	"context"
	"errors"
	"github.com/polRk/mux/events"
	"strings"
)

type Route struct {
	handler Handler

	matchers []matcher
}

func (r *Route) GetMethods() ([]string, error) {
	for _, m := range r.matchers {
		if methods, ok := m.(methodMatcher); ok {
			return methods, nil
		}
	}

	return nil, errors.New("mux: route doesn't have methods")
}

func (r *Route) Match(req *events.APIGatewayProxyRequest) bool {
	for _, m := range r.matchers {
		if m.Match(req) {
			return true
		}
	}

	return false
}

// matcher --------------------------------------------------------------------

type matcher interface {
	Match(*events.APIGatewayProxyRequest) bool
}

func (r *Route) addMatcher(m matcher) *Route {
	r.matchers = append(r.matchers, m)

	return r
}

// Methods --------------------------------------------------------------------

type methodMatcher []string

func (m methodMatcher) Match(e *events.APIGatewayProxyRequest) bool {
	for _, v := range m {
		if v == e.HTTPMethod {
			return true
		}
	}

	return false
}

func (r *Route) Methods(methods ...string) *Route {
	for k, v := range methods {
		methods[k] = strings.ToUpper(v)
	}

	return r.addMatcher(methodMatcher(methods))
}

// Resource -----------------------------------------------------------------------

type resourceMatcher string

func (m resourceMatcher) Match(e *events.APIGatewayProxyRequest) bool {
	return string(m) == e.Resource
}

func (r *Route) Resource(resource string) *Route {
	return r.addMatcher(resourceMatcher(resource))
}

// Handler --------------------------------------------------------------------

// Handler sets a handler for the route.
func (r *Route) Handler(handler Handler) *Route {
	r.handler = handler

	return r
}

func (r *Route) Handle(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return r.handler.Handle(ctx, request)
}
