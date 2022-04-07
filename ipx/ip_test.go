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
		{
			"t5", map[string][]string{
				"X-Real-IP":       {"::1"},
				"X-Forwarded-For": {"::1"},
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
	ip := LocalAddr()
	if ip == "" {
		t.Error("LocalAddr() return empty value")
	}
}

func TestIsPrivate(t *testing.T) {
	testCases := []struct {
		tag      string
		ip       string
		expected bool
		hasError bool
	}{
		{"t1", "127.0.0.1", true, false},
		{"t2", "::1", true, false},
		{"t3", "xxx", false, true},
	}
	for _, testCase := range testCases {
		v, err := IsPrivate(testCase.ip)
		assert.Equal(t, testCase.expected, v, testCase.tag)
		assert.Equal(t, testCase.hasError, err != nil, testCase.tag+" error")
	}
}

func TestIsPublic(t *testing.T) {
	testCases := []struct {
		tag      string
		ip       string
		expected bool
		hasError bool
	}{
		{"t1", "127.0.0.1", false, false},
		{"t2", "::1", false, false},
		{"t3", "xxx", false, true},
		{"t4", "120.228.142.126", true, false},
	}
	for _, testCase := range testCases {
		v, err := IsPublic(testCase.ip)
		assert.Equal(t, testCase.expected, v, testCase.tag)
		assert.Equal(t, testCase.hasError, err != nil, testCase.tag+" error")
	}
}
