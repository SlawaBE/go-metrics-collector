// Package interceptor contains GRPC interceptors for the metrics server
package interceptor

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func NewTrustedSubnetInterceptor(ipNet *net.IPNet) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		var ip string
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			values := md.Get("x-real-ip")
			if len(values) > 0 {
				ip = values[0]
			}
		}
		if len(ip) == 0 {
			return nil, status.Error(codes.PermissionDenied, "missing x-real-ip")
		}

		addr := net.ParseIP(ip)
		if addr == nil {
			return nil, status.Errorf(codes.PermissionDenied, "x-real-ip %s not an ip", ip)
		}

		if !ipNet.Contains(addr) {
			return nil, status.Errorf(codes.PermissionDenied, "x-real-ip %s is not trusted", ip)
		}

		return handler(ctx, req)
	}
}
