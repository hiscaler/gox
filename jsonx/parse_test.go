package jsonx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFind(t *testing.T) {
	testCases := []struct {
		tag          string
		json         string
		path         string
		defaultValue interface{}
		Except       interface{}
	}{
		{"t1", "", "a", "", ""},
		{"t2", `{"a":1}`, "a", 2, "1"},
		{"t2", `{"a":true}`, "a", 2, "true"},
		{"t2", `{"a":true}`, "a.b", false, "false"},
	}
	for _, testCase := range testCases {
		v := Find(testCase.json, testCase.path, testCase.defaultValue).ToString()
		assert.Equal(t, testCase.Except, v, testCase.tag)
	}
}
