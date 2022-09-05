package setx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToSet(t *testing.T) {
	assert.ElementsMatch(t, []int{1, 2, 3}, ToSet([]int{1, 2, 3}), "int1")
	assert.ElementsMatch(t, []int{1, 2, 3}, ToSet([]int{1, 1, 2, 2, 3, 3}), "int2")
	assert.ElementsMatch(t, []float32{1, 2, 3}, ToSet([]float32{1, 2, 3}), "float321")
	assert.ElementsMatch(t, []float64{1, 2, 3}, ToSet([]float64{1, 1, 2, 2, 3, 3}), "float641")
	assert.ElementsMatch(t, []string{"A", "B", "C"}, ToSet([]string{"A", "B", "C", "C", "B", "A"}), "string1")
	assert.ElementsMatch(t, []string{"A", " A ", "B", "C"}, ToSet([]string{" A ", "B", "C", "C", "B", "A"}), "string2")
}

func TestToStringSet(t *testing.T) {
	testCases := []struct {
		A      []string
		B      []string
		Len    int
		Values []string
	}{
		{[]string{"1", "2", "3"}, []string{"0", "1", "4"}, 5, []string{"0", "1", "2", "3", "4"}},
		{[]string{"1", "1", "1"}, []string{"0", "1", "2"}, 3, []string{"0", "1", "2"}},
		{[]string{"   ", "1", "1", "1"}, []string{"0", "1", "2"}, 3, []string{"0", "1", "2"}},
		{[]string{"\tabc\t", " abc ", "1", "1", "1"}, []string{"0", "1", "2"}, 4, []string{"0", "1", "2", "abc"}},
		{[]string{"\tabc\t", " abc ", "1", "1", "1", "ABC"}, []string{"0", "1", "2"}, 5, []string{"0", "1", "2", "abc", "ABC"}},
	}

	for _, testCase := range testCases {
		c := ToStringSet(append(testCase.A, testCase.B...), true)
		if len(c) != testCase.Len {
			t.Errorf("Except %d, actual %d", testCase.Len, len(c))
		}
		for _, value := range testCase.Values {
			exists := false
			for _, v := range c {
				if v == value {
					exists = true
					break
				}
			}
			if !exists {
				t.Errorf("%s not in %#v", value, testCase.Values)
			}
		}
	}
}

func BenchmarkToStringSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToStringSet([]string{"A", "B", "c", "C", "a", "d", "d", "e", "fgh", "FGH", "fGH", "fgH"}, false)
	}
}

func TestIntSliceToSet(t *testing.T) {
	testCases := []struct {
		A      []int
		B      []int
		Len    int
		Values []int
	}{
		{[]int{1, 2, 3}, []int{0, 1, 4}, 5, []int{0, 1, 2, 3, 4}},
		{[]int{1, 1, 1}, []int{0, 1, 2}, 3, []int{0, 1, 2}},
	}

	for _, testCase := range testCases {
		c := ToIntSet(append(testCase.A, testCase.B...))
		if len(c) != testCase.Len {
			t.Errorf("Except %d, actual %d", testCase.Len, len(c))
		}
		for _, value := range testCase.Values {
			exists := false
			for _, v := range c {
				if v == value {
					exists = true
					break
				}
			}
			if !exists {
				t.Errorf("%d not in %#v", value, testCase.Values)
			}
		}
	}
}

func BenchmarkToIntSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToIntSet([]int{1, 2, 3, 3, 45, 5, 6, 56, 56, 56, 77, 6, 7, 67, 678, 78, 78, 8, 78})
	}
}
