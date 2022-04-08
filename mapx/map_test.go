package mapx

import (
	"fmt"
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
		v := assert.Equal(t, test.keys, keys, test.tag)
		if v {
			for k, value := range test.keys {
				assert.Equal(t, value, keys[k], fmt.Sprintf("keys[%d]", k))
			}
		}
	}
}

func BenchmarkStringKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringKeys(map[string]interface{}{"a": 1, "b": 2, "c": "cValue", "d": "dValue"})
	}
}
