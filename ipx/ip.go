package ipx

import (
	"errors"
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

func LocalAddr() (addr string, err error) {
	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, address := range addresses {
		if ipNet, ok := address.(*net.IPNet); ok &&
			!ipNet.IP.IsLoopback() &&
			!ipNet.IP.IsPrivate() &&
			!ipNet.IP.IsLinkLocalUnicast() {
			if ipNet.IP.To4() != nil {
				addr = ipNet.IP.String()
				break
			}
		}
	}
	if addr == "" {
		var conn net.Conn
		conn, err = net.Dial("udp", "8.8.8.8:53")
		if err == nil {
			localAddr := conn.LocalAddr().(*net.UDPAddr)
			addr = strings.Split(localAddr.String(), ":")[0]
		}
	}
	if addr == "" && err == nil {
		err = errors.New("ip: not found")
	}
	return
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
