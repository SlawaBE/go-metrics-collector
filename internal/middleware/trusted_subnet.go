package middleware

import (
	"net"
	"net/http"
)

// TrustedSubnetMiddleware checks that the X-Real-IP header value falls within
// the configured CIDR.
func TrustedSubnetMiddleware(ipNet *net.IPNet, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.Header.Get("X-Real-IP")
		if clientIP == "" {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}

		addr := net.ParseIP(clientIP)
		if addr == nil {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}

		if !ipNet.Contains(addr) {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
