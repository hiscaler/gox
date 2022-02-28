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

func TestStartsWith(t *testing.T) {
	tests := []struct {
		tag           string
		string        []byte
		words         [][]byte
		caseSensitive bool
		except        bool
	}{
		{"t1", []byte("Hello world!"), [][]byte{[]byte("he"), []byte("He")}, false, true},
		{"t2", []byte("Hello world!"), [][]byte{[]byte("he"), []byte("He")}, true, true},
		{"t3", []byte("Hello world!"), [][]byte{[]byte("he")}, true, false},
		{"t4", []byte(""), [][]byte{[]byte("")}, true, true},
		{"t5", []byte(""), nil, true, true},
		{"t6", []byte(""), [][]byte{}, true, true},
		{"t7", []byte("Hello world!"), [][]byte{[]byte("")}, true, true},
	}
	for _, test := range tests {
		b := StartsWith(test.string, test.words, test.caseSensitive)
		assert.Equal(t, test.except, b, test.tag)
	}
}

func TestEndsWith(t *testing.T) {
	tests := []struct {
		tag           string
		string        []byte
		words         [][]byte
		caseSensitive bool
		except        bool
	}{
		{"t1", []byte("Hello world!"), [][]byte{[]byte("he"), []byte("He")}, false, false},
		{"t2", []byte("Hello world!"), [][]byte{[]byte("he"), []byte("He")}, true, false},
		{"t3", []byte("Hello world!"), [][]byte{[]byte("d!"), []byte("!")}, true, true},
		{"t4", []byte("Hello world!"), [][]byte{[]byte("WORLD!")}, false, true},
		{"t5", []byte(""), [][]byte{[]byte("")}, true, true},
		{"t6", []byte(""), nil, true, true},
		{"t7", []byte(""), [][]byte{}, true, true},
		{"t8", []byte("Hello world!"), [][]byte{[]byte("")}, true, true},
	}
	for _, test := range tests {
		b := EndsWith(test.string, test.words, test.caseSensitive)
		assert.Equal(t, test.except, b, test.tag)
	}
}
