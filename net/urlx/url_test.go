package urlx

import (
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
