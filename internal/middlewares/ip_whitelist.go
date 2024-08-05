package middlewares

import (
	"fmt"
	"net"
	"net/http"
)

func IPWhiteListMiddleware(n *net.IPNet) func(next http.Handler) http.Handler {
	whiteLister := NewIPWhiteLister(n)
	return whiteLister.Handler
}

type IPWhiteList struct {
	subnet *net.IPNet
}

func NewIPWhiteLister(subnet *net.IPNet) *IPWhiteList {
	return &IPWhiteList{subnet: subnet}
}

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
