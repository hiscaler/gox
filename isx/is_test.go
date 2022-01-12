package isx

import "testing"

func TestNumber(t *testing.T) {
	testCases := []struct {
		Value    string
		IsNumber bool
	}{
		{"a", false},
		{"111", true},
	}
	for _, testCase := range testCases {
		v := Number(testCase.Value)
		if v != testCase.IsNumber {
			t.Errorf("%s except %v actual %v", testCase.Value, testCase.IsNumber, v)
		}
	}
}
