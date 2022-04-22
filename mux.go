package mux

import (
	"context"
	"github.com/polRk/mux/events"
	"net/http"
)

func NewRouter() *Router {
	return &Router{}
}

type Router struct {
	routes []*Route
}

func (r *Router) Match(req *events.APIGatewayProxyRequest) *Route {
	for _, route := range r.routes {
		if route.Match(req) {
			return route
		}
	}

	return nil
}

func (r *Router) NewRoute() *Route {
	route := &Route{}
	r.routes = append(r.routes, route)
	return route
}

func (r *Router) Handle(resource string, handler Handler) *Route {
	return r.NewRoute().Resource(resource).Handler(handler)
}

func (r *Router) Methods(methods ...string) *Route {
	return r.NewRoute().Methods(methods...)
}

func (r *Router) Serve(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	handler := r.Match(request)

	if handler != nil {
		return handler.Handle(ctx, request)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode:      http.StatusNotFound,
		Body:            "404 page not found",
		IsBase64Encoded: false,
	}, nil
}
