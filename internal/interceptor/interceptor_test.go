package interceptor

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func trustedSubnetCIDR(t *testing.T, cidr string) *net.IPNet {
	t.Helper()
	_, ipNet, err := net.ParseCIDR(cidr)
	require.NoError(t, err)
	return ipNet
}

func newHandler() grpc.UnaryHandler {
	return func(ctx context.Context, req any) (interface{}, error) {
		return "handled", nil
	}
}

// withIncomingMD builds a context with incoming metadata for server-side interceptor tests.
func withIncomingMD(ctx context.Context, md metadata.MD) context.Context {
	return metadata.NewIncomingContext(ctx, md)
}

func TestNewTrustedSubnetInterceptor_Contains(t *testing.T) {
	ipNet := trustedSubnetCIDR(t, "10.0.0.0/8")
	interp := NewTrustedSubnetInterceptor(ipNet)

	md := metadata.New(map[string]string{"x-real-ip": "10.5.3.1"})
	ctx := withIncomingMD(context.Background(), md)

	resp, err := interp(ctx, nil, &grpc.UnaryServerInfo{FullMethod: "/Metrics/UpdateMetrics"}, newHandler())
	require.NoError(t, err)
	assert.Equal(t, "handled", resp)
}

func TestNewTrustedSubnetInterceptor_IPv6InCIDR(t *testing.T) {
	ipNet := trustedSubnetCIDR(t, "::1/128")
	interp := NewTrustedSubnetInterceptor(ipNet)

	md := metadata.New(map[string]string{"x-real-ip": "fc00::dead:beef"})
	ctx := withIncomingMD(context.Background(), md)

	_, err := interp(ctx, nil, &grpc.UnaryServerInfo{FullMethod: "/foo"}, newHandler())
	assert.Error(t, err)
	assert.Equal(t, codes.PermissionDenied, status.Code(err))
	assert.Contains(t, err.Error(), "not trusted")
}

func TestNewTrustedSubnetInterceptor_NotContained(t *testing.T) {
	ipNet := trustedSubnetCIDR(t, "192.168.0.0/24")
	interp := NewTrustedSubnetInterceptor(ipNet)

	md := metadata.New(map[string]string{"x-real-ip": "10.0.0.1"})
	ctx := withIncomingMD(context.Background(), md)

	_, err := interp(ctx, nil, &grpc.UnaryServerInfo{FullMethod: "/Metrics/UpdateMetrics"}, newHandler())
	assert.Error(t, err)
	assert.Equal(t, codes.PermissionDenied, status.Code(err))
	assert.Contains(t, err.Error(), "not trusted")
}

func TestNewTrustedSubnetInterceptor_MissingIP(t *testing.T) {
	ipNet := trustedSubnetCIDR(t, "10.0.0.0/8")
	interp := NewTrustedSubnetInterceptor(ipNet)

	ctx := context.Background()

	_, err := interp(ctx, nil, &grpc.UnaryServerInfo{FullMethod: "/Metrics/UpdateMetrics"}, newHandler())
	assert.Error(t, err)
	assert.Equal(t, codes.PermissionDenied, status.Code(err))
	assert.Contains(t, err.Error(), "missing x-real-ip")
}

func TestNewTrustedSubnetInterceptor_EmptyXRealIP(t *testing.T) {
	ipNet := trustedSubnetCIDR(t, "10.0.0.0/8")
	interp := NewTrustedSubnetInterceptor(ipNet)

	md := metadata.New(map[string]string{})
	ctx := withIncomingMD(context.Background(), md)

	_, err := interp(ctx, nil, &grpc.UnaryServerInfo{FullMethod: "/Metrics/UpdateMetrics"}, newHandler())
	assert.Error(t, err)
	assert.Equal(t, codes.PermissionDenied, status.Code(err))
	assert.Contains(t, err.Error(), "missing x-real-ip")
}

func TestNewTrustedSubnetInterceptor_InvalidIP(t *testing.T) {
	ipNet := trustedSubnetCIDR(t, "10.0.0.0/8")
	interp := NewTrustedSubnetInterceptor(ipNet)

	md := metadata.New(map[string]string{"x-real-ip": "not-an-ip"})
	ctx := withIncomingMD(context.Background(), md)

	_, err := interp(ctx, nil, &grpc.UnaryServerInfo{FullMethod: "/Metrics/UpdateMetrics"}, newHandler())
	assert.Error(t, err)
	assert.Equal(t, codes.PermissionDenied, status.Code(err))
	assert.Contains(t, err.Error(), "not an ip")
}

func TestNewTrustedSubnetInterceptor_MultipleXRealIP_UsesFirst(t *testing.T) {
	ipNet := trustedSubnetCIDR(t, "10.0.0.0/8")
	interp := NewTrustedSubnetInterceptor(ipNet)

	md := metadata.New(map[string]string{
		"x-real-ip": "10.1.1.1",
	})
	ctx := withIncomingMD(context.Background(), md)

	resp, err := interp(ctx, nil, &grpc.UnaryServerInfo{FullMethod: "/Metrics/UpdateMetrics"}, newHandler())
	require.NoError(t, err)
	assert.Equal(t, "handled", resp)
}

func TestNewTrustedSubnetInterceptor_BannedPublicIP(t *testing.T) {
	ipNet := trustedSubnetCIDR(t, "10.0.0.0/8")
	interp := NewTrustedSubnetInterceptor(ipNet)

	md := metadata.New(map[string]string{"x-real-ip": "8.8.8.8"})
	ctx := withIncomingMD(context.Background(), md)

	_, err := interp(ctx, nil, &grpc.UnaryServerInfo{FullMethod: "/Metrics/UpdateMetrics"}, newHandler())
	assert.Error(t, err)
	assert.Equal(t, codes.PermissionDenied, status.Code(err))
	assert.Contains(t, err.Error(), "not trusted")
}

func TestNewTrustedSubnetInterceptor_NetworkBoundary(t *testing.T) {
	ipNet := trustedSubnetCIDR(t, "172.16.0.0/12")
	interp := NewTrustedSubnetInterceptor(ipNet)

	md := metadata.New(map[string]string{"x-real-ip": "172.16.0.1"})
	ctx := withIncomingMD(context.Background(), md)

	resp, err := interp(ctx, nil, &grpc.UnaryServerInfo{FullMethod: "/foo"}, newHandler())
	require.NoError(t, err)
	assert.Equal(t, "handled", resp)
}

func TestNewTrustedSubnetInterceptor_SubnetOutside(t *testing.T) {
	ipNet := trustedSubnetCIDR(t, "172.16.0.0/12")
	interp := NewTrustedSubnetInterceptor(ipNet)

	md := metadata.New(map[string]string{"x-real-ip": "172.32.0.1"})
	ctx := withIncomingMD(context.Background(), md)

	_, err := interp(ctx, nil, &grpc.UnaryServerInfo{FullMethod: "/foo"}, newHandler())
	assert.Error(t, err)
	assert.Equal(t, codes.PermissionDenied, status.Code(err))
}
