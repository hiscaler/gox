package setx

import "testing"

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
