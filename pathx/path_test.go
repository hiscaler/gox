package pathx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilename(t *testing.T) {
	testCases := []struct {
		tag      string
		path     string
		expected string
	}{
		{"t1", "a.jpg", "a"},
		{"t2", "/a/b/c.jpg", "c"},
		{"t3", "/a/b/c", "c"},
		{"t4", "/a/b/c/", "c"},
		{"t5", "/a/b/c/中文.jpg", "中文"},
		{"t5", "https://www.example.com/a/b/c/中文.jpg", "中文"},
	}
	for _, testCase := range testCases {
		v := Filename(testCase.path)
		assert.Equal(t, testCase.expected, v, testCase.tag)
	}
}
