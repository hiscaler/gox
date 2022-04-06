package ipx

import (
	"fmt"
	"math"
	"net"
	"net/http"
	"strings"
)

func RemoteAddr(r *http.Request, mustPublic bool) string {
	if r == nil {
		return ""
	}

	for _, key := range []string{"X-Forwarded-For", "X-Real-IP"} {
		ip := r.Header.Get(key)
		if ip != "" {
			for _, singleIP := range strings.Split(ip, ", ") {
				if !mustPublic || IsPublic(singleIP) {
					return singleIP
				}
			}
		}
	}
	return r.RemoteAddr
}

func IsPrivate(ip string) bool {
	addr := net.ParseIP(ip)
	return addr.IsLoopback() || addr.IsPrivate()
}

func IsPublic(ip string) bool {
	return !IsPrivate(ip)
}

func Number(ip string) (uint, error) {
	addr := net.ParseIP(ip)
	if addr == nil {
		return 0, fmt.Errorf("%s is invalid ip", ip)
	}
	return uint(addr[3]) | uint(addr[2])<<8 | uint(addr[1])<<16 | uint(addr[0])<<24, nil
}

func String(ip uint) (string, error) {
	if ip > math.MaxUint32 {
		return "", fmt.Errorf("%d is not valid ipv4", ip)
	}

	addr := make(net.IP, net.IPv4len)
	addr[0] = byte(ip >> 24)
	addr[1] = byte(ip >> 16)
	addr[2] = byte(ip >> 8)
	addr[3] = byte(ip)
	return addr.String(), nil
}
