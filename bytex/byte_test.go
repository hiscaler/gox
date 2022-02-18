package bytex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func TestToString(t *testing.T) {
	tests := []struct {
		tag        string
		bytesValue []byte
		string     string
	}{
		{"t1", []byte{'a'}, "a"},
		{"t2", []byte("abc"), "abc"},
		{"t3", []byte("a b c "), "a b c "},
	}
	for _, test := range tests {
		s := ToString(test.bytesValue)
		assert.Equal(t, test.string, s, test.tag)
	}
}
