package isx

import "testing"

func TestNumber(t *testing.T) {
	testCases := []struct {
		Value    string
		IsNumber bool
	}{
		{"a", false},
		{"111", true},
		{"1.23", true},
		{"1,234.5", true},
		{"1234.5,", false},
		{"12345.", false},
		{" 12345.6   ", true},
		{" 12345. 6   ", false},
	}
	for _, testCase := range testCases {
		v := Number(testCase.Value)
		if v != testCase.IsNumber {
			t.Errorf("%s except %v actual %v", testCase.Value, testCase.IsNumber, v)
		}
	}
}
