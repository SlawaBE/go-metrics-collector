package utils

import (
	"net"

	"github.com/SlawaBE/go-metrics-collector/internal/logger"
	"go.uber.org/zap"
)

// LocalIP returns the local IP address for the given remote target address.
// Works with both IPv4 and IPv6.
func LocalIP(targetAddr string) (string, error) {
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
	return localAddr.IP.String(), nil
}
