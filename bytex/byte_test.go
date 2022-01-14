package bytex

import "testing"

func TestIsEmpty(t *testing.T) {
	testCases := []struct {
		Number int
		Byte   []byte
		Except bool
	}{
		{1, []byte("a"), false},
		{2, []byte(""), true},
		{3, []byte("   "), true},
	}

	for _, testCase := range testCases {
		isEmpty := IsEmpty(testCase.Byte)
		if isEmpty != testCase.Except {
			t.Errorf("%d except: %#v, actual: %#v", testCase.Number, testCase.Except, isEmpty)
		}
	}
}
