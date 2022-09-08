package slicex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	a := []int{0, 1}
	b := Map(a, func(v int) int {
		return v + 1
	})
	assert.Equal(t, []int{0, 1}, a, "map.int")
	assert.Equal(t, []int{1, 2}, b, "map.int")

	a1 := []string{"a", "b"}
	b1 := Map(a1, func(v string) string {
		return v + "1"
	})
	assert.Equal(t, []string{"a", "b"}, a1, "map.string")
	assert.Equal(t, []string{"a1", "b1"}, b1, "map.string")
}

func TestFilter(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, Filter([]int{0, 1, 2, 3}, func(v int) bool {
		return v > 0
	}), "map.int")
	assert.Equal(t, []string{"c"}, Filter([]string{"a", "b", "c", "a", "b"}, func(v string) bool {
		return v == "c"
	}), "map.string")
}

func TestStringToInterface(t *testing.T) {
	tests := []struct {
		tag      string
		input    []string
		expected []interface{}
	}{
		{"t0", []string{"a", "b", "c"}, []interface{}{"a", "b", "c"}},
		{"t1", nil, []interface{}{}},
	}
	for _, test := range tests {
		v := StringToInterface(test.input)
		assert.Equal(t, test.expected, v, test.tag)
	}
}

func BenchmarkStringToInterface(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringToInterface([]string{"a", "b", "c"})
	}
}

func TestStringSliceEqual(t *testing.T) {
	testCases := []struct {
		A             []string
		B             []string
		CaseSensitive bool
		IgnoreEmpty   bool
		Trim          bool
		Except        bool
	}{
		{[]string{}, []string{}, false, true, true, true},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, true, false, true, true},
		{[]string{"a", "b", "c"}, []string{"a", "b   ", "     c"}, true, false, true, true},
		{[]string{"a", "b", "c"}, []string{"a", "b   ", "     c"}, true, false, false, false},
		{[]string{"a", "b", "c", ""}, []string{"a", "b   ", "     c"}, true, false, true, false},
		{[]string{"a", "b", "c", ""}, []string{"a", "b   ", "     c"}, true, true, true, true},
		{[]string{"a", "b", "c", ""}, []string{"a", "b   ", "     c", ""}, true, false, true, true},
		{[]string{"A", "B", "C"}, []string{"a", "b", "c"}, true, true, true, false},
		{[]string{"A", "B", "C"}, []string{"a", "b", "c"}, false, true, true, true},
		{[]string{"A", "B", "C"}, []string{"b", "c", "a"}, false, true, true, true},
		{[]string{"   ", "", " "}, []string{""}, false, true, true, true},
		{[]string{}, []string{" ", ""}, false, true, true, true},
		{[]string{}, []string{"a", "b"}, false, true, true, false},
		{nil, []string{}, false, true, true, false},
		{[]string{}, nil, false, true, true, false},
		{nil, nil, false, true, true, true},
	}
	for i, testCase := range testCases {
		equal := StringSliceEqual(testCase.A, testCase.B, testCase.CaseSensitive, testCase.IgnoreEmpty, testCase.Trim)
		if equal != testCase.Except {
			t.Errorf("%d except %v actual %v", i, testCase.Except, equal)
		}
	}
}

func TestIntSliceEqual(t *testing.T) {
	testCases := []struct {
		A      []int
		B      []int
		Except bool
	}{
		{[]int{}, []int{}, true},
		{[]int{0, 1, 2}, []int{0, 1, 2}, true},
		{[]int{0, 1, 2}, []int{2, 1, 0}, true},
		{[]int{0, 1, 2}, []int{1, 2}, false},
		{[]int{0, 1, 1, 2}, []int{0, 1, 2}, false},
		{[]int{0, 1, 1, 2}, []int{0, 1, 2, 1}, true},
		{nil, []int{}, false},
		{nil, nil, true},
		{[]int{}, nil, false},
	}

	for i, testCase := range testCases {
		equal := IntSliceEqual(testCase.A, testCase.B)
		if equal != testCase.Except {
			t.Errorf("%d except %v actual %v", i, testCase.Except, equal)
		}
	}
}

func TestStringSliceReverse(t *testing.T) {
	testCases := []struct {
		Before []string
		After  []string
	}{
		{[]string{"a"}, []string{"a"}},
		{[]string{"a", "b"}, []string{"b", "a"}},
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}},
	}

	for _, testCase := range testCases {
		values := StringSliceReverse(testCase.Before)
		if len(values) != len(testCase.After) {
			t.Errorf("%#v reverse after value except: %#v, actual: %#v", testCase.Before, testCase.After, values)
		} else {
			for j, v := range values {
				if testCase.After[j] != v {
					t.Errorf("%#v reverse after value except: %#v, actual: %#v", testCase.Before, testCase.After, values)
					break
				}
			}
		}
	}
}

func TestStringSliceDiff(t *testing.T) {
	testCases := []struct {
		Number         int
		OriginalValues [][]string
		DiffValue      []string
	}{
		{1, [][]string{{"a", "b", "c"}, {"a", "b", "d"}}, []string{"c"}},
		{1, [][]string{{"a", "b", "c"}, {"a"}}, []string{"b", "c"}},
		{2, [][]string{{"a", "b", "d"}, {"a", "b", "c"}}, []string{"d"}},
		{3, [][]string{{"a", "b", "c"}, {"a", "b", "c"}}, []string{}},
		{4, [][]string{{"a", "b", ""}, {"a", "b", "c"}}, []string{""}},
		{5, [][]string{{"a", "b", "c"}, {"a", "b"}, {"c"}}, []string{}},
		{6, [][]string{{"a"}, {"b"}, {"c", "c1"}, {"d"}}, []string{"a"}},
		{7, [][]string{nil, {"a"}, {"b"}, {"c", "c1"}, {"d"}}, []string{}},
		{8, [][]string{nil}, []string{}},
		{9, [][]string{nil, nil, nil}, []string{}},
	}

	for _, testCase := range testCases {
		values := StringSliceDiff(testCase.OriginalValues...)
		if !StringSliceEqual(values, testCase.DiffValue, true, false, true) {
			t.Errorf("%d: diff values except: %#v, actual: %#v", testCase.Number, testCase.DiffValue, values)
		}
	}
}

func TestIntSliceDiff(t *testing.T) {
	testCases := []struct {
		Number         int
		OriginalValues [][]int
		DiffValue      []int
	}{
		{1, [][]int{{1, 2, 3}, {1, 2, 4}}, []int{3}},
		{2, [][]int{{1, 2, 3}, {1, 2, 2, 3}, {3, 4, 5}}, []int{}},
		{3, [][]int{{1, 2, 3}, {1}, {2}, {3}}, []int{}},
		{4, [][]int{{1, 2, 3}, {1, 2, 4, 0, 2, 1}}, []int{3}},
		{5, [][]int{{1, 2, 2, 3}, {1}}, []int{2, 2, 3}},
		{6, [][]int{}, []int{}},
		{7, [][]int{nil, {1, 2, 3}}, []int{}},
		{8, [][]int{nil, nil, {1, 2, 3}}, []int{}},
	}

	for _, testCase := range testCases {
		values := IntSliceDiff(testCase.OriginalValues...)
		if !IntSliceEqual(values, testCase.DiffValue) {
			t.Errorf("%d: diff values except: %#v, actual: %#v", testCase.Number, testCase.DiffValue, values)
		}
	}
}

func TestToInterface(t *testing.T) {
	type args struct {
		values interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{"t1", args{[]int{1, 2, 3}}, []interface{}{1, 2, 3}},
		{"t2", args{[]string{"A", "B", "C"}}, []interface{}{"A", "B", "C"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var value interface{}
			switch tt.args.values.(type) {
			case []string:
				value = ToInterface(tt.args.values.([]string))
			case []int:
				value = ToInterface(tt.args.values.([]int))
			}
			assert.Equalf(t, tt.want, value, "ToInterface(%v)", tt.args.values)
		})
	}
}

func TestChunk(t *testing.T) {
	assert.ElementsMatch(t, Chunk([]int{1, 2, 3}, 0), [][]int{{1, 2, 3}}, "int0")
	assert.ElementsMatch(t, Chunk([]int{1, 2, 3}, 1), [][]int{{1}, {2}, {3}}, "int1")
	assert.ElementsMatch(t, Chunk([]int{1, 2, 3}, 2), [][]int{{1, 2}, {3}}, "int2")
	assert.ElementsMatch(t, Chunk([]int{1, 2, 3}, 3), [][]int{{1, 2, 3}}, "int3")
}
