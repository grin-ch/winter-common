package grpc_util

import (
	"context"
	"net"

	"google.golang.org/grpc/peer"
)

// CtxIp
func CtxIp(ctx context.Context) string {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		return ""
	}
	if pr.Addr == net.Addr(nil) {
		return ""
	}
	return pr.Addr.String()
}
