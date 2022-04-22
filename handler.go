package mux

import (
	"context"
	"github.com/polRk/mux/events"
)

type Handler interface {
	Handle(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)
}

type HandlerFunc func(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)

func (f HandlerFunc) Handle(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return f(ctx, request)
}
