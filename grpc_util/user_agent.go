package grpc_util

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const UserAgentKey = "user-agent"

func CtxUserAgent(ctx context.Context) string {
	val, _ := ctx.Value(UserAgentKey).(string)
	return val
}

func InterceptorUserAgent(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		userAgent := md.Get(UserAgentKey)
		if len(userAgent) > 0 {
			ctx = context.WithValue(ctx, UserAgentKey, userAgent[0])
		}
	}
	return handler(ctx, req)
}
