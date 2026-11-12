package middleware

import (
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrustedSubnetMiddleware(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	tests := []struct {
		name           string
		trustedSubnet  string
		xRealIP        string
		expectedStatus int
	}{
		{
			name:           "IP in CIDR allowed",
			trustedSubnet:  "192.168.1.0/24",
			xRealIP:        "192.168.1.50",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "IP outside CIDR forbidden",
			trustedSubnet:  "192.168.1.0/24",
			xRealIP:        "10.0.0.1",
			expectedStatus: http.StatusForbidden,
		},
		{
			name:           "missing X-Real-IP forbidden",
			trustedSubnet:  "192.168.1.0/24",
			xRealIP:        "",
			expectedStatus: http.StatusForbidden,
		},
		{
			name:           "invalid IP forbidden",
			trustedSubnet:  "192.168.1.0/24",
			xRealIP:        "not-an-ip",
			expectedStatus: http.StatusForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ipNet, netErr := net.ParseCIDR(tt.trustedSubnet)
			assert.NoError(t, netErr)

			middleware := TrustedSubnetMiddleware(ipNet, nextHandler)
			req := httptest.NewRequest(http.MethodPost, "/updates", nil)
			req.Header.Set("X-Real-IP", tt.xRealIP)
			rec := httptest.NewRecorder()

			middleware.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}
