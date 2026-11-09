package utils

import (
	"net"

	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"go.uber.org/zap"
)

// LocalIPv4 returns the first non-loopback IPv4 address found on the host.
// It is used by the agent to populate the X-Real-IP header when sending
// metrics to the server
func LocalIPv4(targetAddr string) (string, error) {
	conn, err := net.Dial("udp", targetAddr)
	if err != nil {
		return "", err
	}
	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			logger.Log.Error("failed to close dial connection", zap.Error(closeErr))
		}
	}()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.To4().String(), nil
}
