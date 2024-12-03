package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type InterceptorClient struct {
	AccessToken string
}

func NewInterceptorClient() *InterceptorClient {
	return &InterceptorClient{}
}

func (i *InterceptorClient) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		return invoker(metadata.AppendToOutgoingContext(ctx, "authorization", i.AccessToken), method, req, reply, cc, opts...)
	}
}
