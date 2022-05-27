package spreedsheetx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColumnName(t *testing.T) {
	testCases := []struct {
		tag      string
		name     string
		offset   int
		expected string
		hasError bool
	}{
		{"tag0", "", 0, "", true},
		{"tag0.1", "1", 0, "", true},
		{"tag0.2", "A1", 0, "", true},
		{"tag1", "Z", 0, "Z", false},
		{"tag1.1", "Z", 1, "AA", false},
		{"tag1.2", "A", 1, "B", false},
		{"tag1.3", "a", 3, "D", false},
		{"tag2.1", "ZZ", 1, "AAA", false},
		{"tag3.1", "ZZZ", 1, "AAAA", false},
		{"tag3.2", "ZZX", 1, "ZZY", false},
		{"tag3.3", "ZZX", 2, "ZZZ", false},
		{"tag3.4", "ZZX", 3, "AAAA", false},
		{"tag3.5", "AAA", 1, "AAB", false},
	}
	for _, testCase := range testCases {
		index, err := ColumnName(testCase.name, testCase.offset)
		assert.Equal(t, testCase.expected, index, testCase.tag)
		assert.Equal(t, testCase.hasError, err != nil, testCase.tag)
	}
}
