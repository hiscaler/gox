package mapx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringKeys(t *testing.T) {
	tests := []struct {
		tag   string
		value map[string]interface{}
		keys  []string
	}{
		{"t0", nil, []string{}},
		{"t1", map[string]interface{}{"a": 1, "b": 2}, []string{"a", "b"}},
		{"t2", map[string]interface{}{"b": 1, "a": 2}, []string{"a", "b"}},
		{"t3", map[string]interface{}{"a": 1, "b": 2, "": 3}, []string{"", "a", "b"}},
	}

	for _, test := range tests {
		keys := StringKeys(test.value)
		assert.Equal(t, test.keys, keys, test.tag)
	}
}

func BenchmarkStringKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringKeys(map[string]interface{}{"a": 1, "b": 2, "c": "cValue", "d": "dValue"})
	}
}
