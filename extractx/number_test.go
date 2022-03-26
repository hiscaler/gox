package extractx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumber(t *testing.T) {
	testCases := []struct {
		tag      string
		string   string
		expected string
	}{
		{"t1", "123", "123"},
		{"t2", "12.3", "12.3"},
		{"t3", "1,234.3", "1234.3"},
		{"t4", "          ab 1 123", "1"},
		{"t5", ".", ""},
		{"t6", ",", ""},
		{"t7", ".,", ""},
		{"t8", "100 23.", "100"},
		{"t9", "$100 $23.", "100"},
		{"t9", "-1", "-1"},
		{"t10", "-1-1", "-1"}, // todo maybe is empty
		{"t11", "1.0 out of 5 stars", "1.0"},
	}
	for _, testCase := range testCases {
		n := Number(testCase.string)
		assert.Equal(t, testCase.expected, n, testCase.tag)
	}
}

func TestNumbers(t *testing.T) {
	testCases := []struct {
		tag      string
		string   string
		expected []string
	}{
		{"t1", "123", []string{"123"}},
		{"t2", "12.3", []string{"12.3"}},
		{"t3", "1,234.3", []string{"1234.3"}},
		{"t4", "          ab 1 123", []string{"1", "123"}},
		{"t5", "          ab .1 123", []string{"1", "123"}},
		{"t5", "          ab ,1 123", []string{"1", "123"}},
		{"t6", "          ab 1. 123", []string{"1", "123"}},
		{"t7", "$100,$200", []string{"100", "200"}},
		{"t8", "1,2.3,4", []string{"12.34"}},
		{"t9", "1, 2.3, 4", []string{"1", "2.3", "4"}},
		{"t10", "-123,4", []string{"-1234"}},
		{"t11", "1-1", []string{"1", "-1"}}, // todo May be return empty string
	}
	for _, testCase := range testCases {
		n := Numbers(testCase.string)
		assert.Equal(t, testCase.expected, n, testCase.tag)
	}
}
