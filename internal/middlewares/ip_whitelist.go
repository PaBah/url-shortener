package middlewares

import (
	"fmt"
	"net"
	"net/http"
)

// IPWhiteListMiddleware return middleware handler
func IPWhiteListMiddleware(n *net.IPNet) func(next http.Handler) http.Handler {
	whiteLister := NewIPWhiteLister(n)
	return whiteLister.Handler
}

// IPWhiteList struct with white list middleware data
type IPWhiteList struct {
	subnet *net.IPNet
}

// NewIPWhiteLister creates new IPWhiteList
func NewIPWhiteLister(subnet *net.IPNet) *IPWhiteList {
	return &IPWhiteList{subnet: subnet}
}

// Handler IPWhiteList middlewares handler
func (wl *IPWhiteList) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := net.ParseIP(r.Header.Get("X-Real-IP"))
		if !wl.subnet.Contains(ip) {
			http.Error(w, fmt.Sprintf("ip %s is not from trusted subnet", ip.String()), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
