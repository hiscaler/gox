package slicex

import (
	"testing"
)

func TestStringSliceEqual(t *testing.T) {
	testCases := []struct {
		A           []string
		B           []string
		IgnoreCase  bool
		IgnoreEmpty bool
		Trim        bool
		Except      bool
	}{
		{[]string{}, []string{}, true, true, true, true},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, false, false, true, true},
		{[]string{"a", "b", "c"}, []string{"a", "b   ", "     c"}, false, false, true, true},
		{[]string{"a", "b", "c"}, []string{"a", "b   ", "     c"}, false, false, false, false},
		{[]string{"a", "b", "c", ""}, []string{"a", "b   ", "     c"}, false, false, true, false},
		{[]string{"a", "b", "c", ""}, []string{"a", "b   ", "     c"}, false, true, true, true},
		{[]string{"a", "b", "c", ""}, []string{"a", "b   ", "     c", ""}, false, false, true, true},
		{[]string{"A", "B", "C"}, []string{"a", "b", "c"}, false, true, true, false},
		{[]string{"A", "B", "C"}, []string{"a", "b", "c"}, true, true, true, true},
		{[]string{"A", "B", "C"}, []string{"b", "c", "a"}, true, true, true, true},
		{[]string{"   ", "", " "}, []string{""}, true, true, true, true},
		{[]string{}, []string{" ", ""}, true, true, true, true},
		{[]string{}, []string{"a", "b"}, true, true, true, false},
	}
	for i, testCase := range testCases {
		equal := StringSliceEqual(testCase.A, testCase.B, testCase.IgnoreCase, testCase.IgnoreEmpty, testCase.Trim)
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
	}

	for i, testCase := range testCases {
		equal := IntSliceEqual(testCase.A, testCase.B)
		if equal != testCase.Except {
			t.Errorf("%d except %v actual %v", i, testCase.Except, equal)
		}
	}
}
