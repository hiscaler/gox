package setx

import "testing"

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
