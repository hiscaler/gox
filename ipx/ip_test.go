package ipx

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRemoteAddr(t *testing.T) {
	request := &http.Request{
		Header: map[string][]string{},
	}
	testCases := []struct {
		tag        string
		headers    map[string][]string
		mustPublic bool
		expected   string
	}{
		{
			"t1", map[string][]string{
				"X-Real-IP":       {"127.0.0.1"},
				"X-Forwarded-For": {"127.0.0.1"},
			}, false, "127.0.0.1",
		},
		{
			"t2", map[string][]string{
				"X-Real-IP":       {"127.0.0.1:8080"},
				"X-Forwarded-For": {"127.0.0.1:8080"},
			}, false, "127.0.0.1",
		},
		{
			"t3", map[string][]string{
				"X-Real-IP":       {"127.0.0.1"},
				"X-Forwarded-For": {"127.0.0.1"},
			}, true, "",
		},
		{
			"t4", map[string][]string{
				"X-Real-IP":       {"127.0.0.1:8080"},
				"X-Forwarded-For": {"127.0.0.1:8080"},
			}, true, "",
		},
	}
	for _, testCase := range testCases {
		request.Header = testCase.headers
		addr := RemoteAddr(request, testCase.mustPublic)
		assert.Equal(t, testCase.expected, addr, testCase.tag)
	}
}

func TestLocalAddr(t *testing.T) {
	_, err := LocalAddr()
	if err != nil {
		t.Errorf("LocalAddr() error: %s", err.Error())
	}
}
