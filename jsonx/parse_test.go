package jsonx

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestParser_Find(t *testing.T) {
	testCases := []struct {
		tag          string
		json         string
		path         string
		defaultValue interface{}
		valueKind    reflect.Kind
		Except       interface{}
	}{
		{"string1", "", "a", "", reflect.String, ""},
		{"string2", `{"a":1}`, "a", 2, reflect.String, "1"},
		{"string3", `{"a":true}`, "a", 2, reflect.String, "true"},
		{"string4", `{"a":true}`, "a.b", false, reflect.String, "false"},
		{"string5", `{"a":{"b": {"c": 123}}}`, "a.b", "{}", reflect.String, `{"c":123}`},
		{"string5", `{"a":{"b": {"c": 123}}}`, "a.b.c", "", reflect.String, "123"},
		{"string6", `{"a":{"b": {"c": [1,2,3]}}}`, "a.b.c.0", "", reflect.String, "1"},
		{"string6", `{"a":{"b": {"c": [1,2,3]}}}`, "a.b.c.2", "", reflect.String, "3"},
		{"int1", `{"a":1}`, "a", 2, reflect.Int, 1},
		{"int2", `{"a":1}`, "aa", 2, reflect.Int, 2},
		{"int641", `{"a":1}`, "a", 2, reflect.Int64, int64(1)},
		{"int641", `{"a":1}`, "aa", 2, reflect.Int64, int64(2)},
		{"bool1", `{"a":true}`, "a", false, reflect.Bool, true},
		{"bool2", `{"a":true}`, "a.b", false, reflect.Bool, false},
		{"float321", `{"a":1.23}`, "a", 0, reflect.Float32, float32(1.23)},
		{"float322", `{"a":1.23}`, "b", 0, reflect.Float32, float32(0)},
		{"float641", `{"a":1.23}`, "a", 0, reflect.Float64, 1.23},
		{"float642", `{"a":1.23}`, "b", 0, reflect.Float64, 0.0},
	}
	for _, testCase := range testCases {
		var v interface{}
		switch testCase.valueKind {
		case reflect.String:
			v = NewParser(testCase.json).Find(testCase.path, testCase.defaultValue).ToString()
		case reflect.Int:
			v = NewParser(testCase.json).Find(testCase.path, testCase.defaultValue).ToInt()
		case reflect.Int64:
			v = NewParser(testCase.json).Find(testCase.path, testCase.defaultValue).ToInt64()
		case reflect.Float32:
			v = NewParser(testCase.json).Find(testCase.path, testCase.defaultValue).ToFloat32()
		case reflect.Float64:
			v = NewParser(testCase.json).Find(testCase.path, testCase.defaultValue).ToFloat64()
		case reflect.Bool:
			v = NewParser(testCase.json).Find(testCase.path, testCase.defaultValue).ToBool()
		}
		assert.Equal(t, testCase.Except, v, testCase.tag)
	}
}

func TestParser_Exists(t *testing.T) {
	testCases := []struct {
		tag    string
		json   string
		path   string
		Except bool
	}{
		{"string1", "", "a", false},
		{"string1", `{"a"}`, "a", false},
		{"string1", `{"a":1}`, "a", true},
		{"string1", `{"a":[0,1,2]}`, "a.1", true},
	}
	for _, testCase := range testCases {
		v := NewParser(testCase.json).Exists(testCase.path)
		assert.Equal(t, testCase.Except, v, testCase.tag)
	}
}
