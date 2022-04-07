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

	for _, key := range []string{"X-Forwarded-For", "X-Real-IP", "X-Appengine-Remote-Addr"} {
		value := r.Header.Get(key)
		if value != "" {
			for _, item := range strings.Split(value, ",") {
				var ip string
				if strings.ContainsRune(item, ':') {
					if host, _, err := net.SplitHostPort(strings.TrimSpace(item)); err != nil {
						continue
					} else {
						ip = host
					}
				} else {
					ip = strings.TrimSpace(item)
				}

				if mustPublic {
					if v, e := IsPublic(ip); e != nil && v {
						return ip
					}
				} else {
					return ip
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
		for _, address := range []string{"114.114.114.114:53", "8.8.8.8:53"} {
			var conn net.Conn
			conn, err = net.Dial("udp", address)
			if err == nil {
				conn.Close()
				localAddr := conn.LocalAddr().(*net.UDPAddr)
				addr = strings.Split(localAddr.String(), ":")[0]
				break
			}
		}

	}
	if addr == "" && err == nil {
		err = errors.New("ipx: local address not found")
	}
	return
}

func IsPrivate(ip string) (v bool, err error) {
	addr := net.ParseIP(ip)
	if addr == nil {
		err = fmt.Errorf("ipx: %s address is invalid", ip)
	} else {
		v = addr.IsLoopback() || addr.IsPrivate() || addr.IsLinkLocalUnicast()
	}
	return
}

func IsPublic(ip string) (v bool, err error) {
	v, err = IsPrivate(ip)
	v = err == nil && !v
	return
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
