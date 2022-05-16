package mapx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeys(t *testing.T) {
	tests := []struct {
		tag   string
		value interface{}
		keys  []string
	}{
		{"t0", nil, nil},
		{"t1", map[string]interface{}{"a": 1, "b": 2}, []string{"a", "b"}},
		{"t2", map[string]interface{}{"b": 1, "a": 2}, []string{"a", "b"}},
		{"t3", map[string]interface{}{"a": 1, "b": 2, "": 3}, []string{"", "a", "b"}},
		{"t4", map[string]string{"a": "1", "b": "2", "": "3"}, []string{"", "a", "b"}},
		{"t4", map[int]string{1: "1", 3: "3", 2: "2"}, []string{"1", "2", "3"}},
		{"t4", map[float64]string{1.1: "1", 3: "3", 2: "2"}, []string{"1.1", "2", "3"}},
		{"t4", map[bool]string{true: "1", false: "3"}, []string{"0", "1"}},
	}

	for _, test := range tests {
		keys := Keys(test.value)
		v := assert.Equal(t, test.keys, keys, test.tag)
		if v {
			for k, value := range test.keys {
				assert.Equal(t, value, keys[k], fmt.Sprintf("keys[%d]", k))
			}
		}
	}
}

func BenchmarkKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Keys(map[interface{}]interface{}{"a": 1, "b": 2, "c": "cValue", "d": "dValue", 1: 1, 2: 2})
	}
}

func TestStringMapStringEncode(t *testing.T) {
	tests := []struct {
		tag      string
		value    map[string]string
		expected string
	}{
		{"t0", nil, ""},
		{"t1", map[string]string{"a": "1", "b": "2"}, "a=1&b=2"},
		{"t2", map[string]string{"b": "1", "a": "2"}, "a=2&b=1"},
		{"t3", map[string]string{"a": "1", "b": "2", "c": "3"}, "a=1&b=2&c=3"},
		{"t4", map[string]string{"a": "1", "b": "2", "": "3"}, "=3&a=1&b=2"},
		{"t4", map[string]string{"1": "1", "3": "3", "2": "2"}, "1=1&2=2&3=3"},
	}

	for _, test := range tests {
		s := StringMapStringEncode(test.value)
		assert.Equal(t, test.expected, s, test.tag)
	}
}
