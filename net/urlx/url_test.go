package urlx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestURL_AddValue(t *testing.T) {
	type testCase struct {
		Number int
		Path   string
		Values map[string]string
		Except string
	}
	testCases := []testCase{
		{1, "https://www.example.com/a/b/c/1.txt?a=1&b=2", map[string]string{"a": "11", "b": "22"}, "https://www.example.com/a/b/c/1.txt?a=11&b=22"},
		{1, "https://www.example.com/a/b/c/1.txt?a=1&b=2&c=3", map[string]string{"a": "11", "c": "33"}, "https://www.example.com/a/b/c/1.txt?a=11&b=2&c=33"},
		{2, "https://www.example.com/a/b/c/1.txt?a=1&b=2#abc", map[string]string{"a": "11"}, "https://www.example.com/a/b/c/1.txt?a=11&b=2#abc"},
		{3, "https://www.example.com/a/b/c/1.txt?a=1&b=2#abc", map[string]string{"A": "11"}, "https://www.example.com/a/b/c/1.txt?A=11&a=1&b=2#abc"},
		{4, "https://www.example.com/a/b/c/1.txt?b=1&a=2#abc", map[string]string{"A": "11"}, "https://www.example.com/a/b/c/1.txt?A=11&a=2&b=1#abc"},
		{5, "https://www.example.com", map[string]string{"A": "11"}, "https://www.example.com?A=11"},
		{6, "https://www.example.com/", map[string]string{"A": "11"}, "https://www.example.com/?A=11"},
	}

	for _, tc := range testCases {
		url := NewURL(tc.Path)
		for k, v := range tc.Values {
			url.SetValue(k, v)
		}
		s := url.String()
		if s != tc.Except {
			t.Errorf("%d except: %s, actual: %s", tc.Number, tc.Except, s)
		}
	}
}

func TestURL_DeleteValue(t *testing.T) {
	type testCase struct {
		Number     int
		Path       string
		DeleteKeys []string
		Except     string
	}
	testCases := []testCase{
		{1, "https://www.example.com/a/b/c/1.txt?a=1&b=2#abc", []string{"a", "b"}, "https://www.example.com/a/b/c/1.txt?#abc"},
		{1, "https://www.example.com/a/b/c/1.txt?a=1&b=2#abc", []string{"a"}, "https://www.example.com/a/b/c/1.txt?b=2#abc"},
		{2, "https://www.example.com/a/b/c/1.txt", []string{"a", "b"}, "https://www.example.com/a/b/c/1.txt"},
		{2, "https://www/a/b/c/1.txt", []string{"a", "b"}, "https://www/a/b/c/1.txt"},
	}

	for _, tc := range testCases {
		url := NewURL(tc.Path)
		for _, v := range tc.DeleteKeys {
			url.DelKey(v)
		}
		s := url.String()
		if s != tc.Except {
			t.Errorf("%d except: %s, actual: %s", tc.Number, tc.Except, s)
		}
	}
}

func TestIsAbsolute(t *testing.T) {
	testCases := []struct {
		tag   string
		url   string
		isAbs bool
	}{
		{"t0.1", "https://www.a.com", true},
		{"t0.2", "http://www.a.com", true},
		{"t0.3", "//www.a.com", true},
		{"t0.4", "//a.b", true},
		{"t0.5", "//abc", false},
		{"t0.6", "//abc...", false},
		{"t0.7", "//.a.b", false},
		{"t0.8", "//a.b..", false},

		{"t1.1", "httpa.com", false},
		{"t1.2", "httpa.com//", false},
		{"t1.3", "//", false},
		{"t1.4", "//a", false},
		{"t1.5", "//....a", false},
	}
	for _, testCase := range testCases {
		isAbs := IsAbsolute(testCase.url)
		assert.Equal(t, testCase.isAbs, isAbs, testCase.tag)
	}
}
