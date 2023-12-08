package ctx_util

import (
	"context"
	"net"

	"google.golang.org/grpc/peer"
)

// ClientIp 通过从metadata中获取远程地址信息
func ClientIp(ctx context.Context) string {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		return ""
	}
	if pr.Addr == net.Addr(nil) {
		return ""
	}
	return pr.Addr.String()
}
