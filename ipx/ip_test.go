package ipx

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRemoteAddr(t *testing.T) {
	request := &http.Request{
		Header: map[string][]string{
			"X-Real-IP": {"127.0.0.1"},
			"X-Forwarded-For": {"127.0.0.1"},
		},
	}
	fmt.Println(RemoteAddr(request, false))
}
