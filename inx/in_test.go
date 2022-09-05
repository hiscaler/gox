package inx

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"strconv"
	"testing"
)

func TestIn(t *testing.T) {
	assert.Equal(t, true, In(1, []int{1, 2, 3, 4}), "int1")
	assert.Equal(t, false, In(1, []int{2, 3, 4, 5}), "int2")
	assert.Equal(t, false, In(1, nil), "int3")
	assert.Equal(t, false, In(1, []int{}), "int4")
	assert.Equal(t, true, In(1, []float64{1.0, 2.0, 3.0}), "float1")
	assert.Equal(t, false, In(1.1, []float64{1.0, 2.0, 3.0}), "float2")
	assert.Equal(t, true, In(true, []bool{true, false, false}), "bool1")
	assert.Equal(t, false, In(true, []bool{false, false, false}), "bool2")
}

func BenchmarkIn(b *testing.B) {
	b.StopTimer()
	ss := make([]string, 100000)
	for i := 0; i < 100000; i++ {
		ss[i] = strconv.Itoa(100000 - i)
	}
	sort.Strings(ss)
	b.StartTimer()
	In("1", ss)
}

func BenchmarkStringIn(b *testing.B) {
	b.StopTimer()
	ss := make([]string, 100000)
	for i := 0; i < 100000; i++ {
		ss[i] = strconv.Itoa(100000 - i)
	}
	b.StartTimer()
	StringIn("1", ss...)
}

func TestIntIn(t *testing.T) {
	testCases := []struct {
		tag      string
		i        int
		ii       []int
		expected bool
	}{
		{"t1", 1, []int{1, 2, 3, 4, 4}, true},
		{"t2", 1, []int{2, 3, 4, 5}, false},
		{"t3", 1, nil, false},
		{"t4", 1, []int{}, false},
	}
	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, IntIn(testCase.i, testCase.ii...), testCase.tag)
	}
}

func BenchmarkIntIn(b *testing.B) {
	b.StopTimer()
	ii := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		ii[i] = 100000 - i
	}
	b.StartTimer()
	IntIn(1, ii...)
}
