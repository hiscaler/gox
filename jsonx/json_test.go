package jsonx

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestToJson(t *testing.T) {
	testCases := []struct {
		Number       int
		Value        interface{}
		DefaultValue string
		Except       string
	}{
		{1, []string{}, "[]", "[]"},
		{2, struct{}{}, "", "{}"},
		{3, struct {
			Name string
			Age  int
		}{"Hello", 12}, "", `{"Name":"hello","Age":12}`},
		{4, struct {
			Name string `json:"a"`
			Age  int    `json:"b"`
		}{"Hello", 12}, "", `{"a":"hello","b":12}`},
		{5, nil, "abc", "null"},
		{6, []int{1, 2}, "null", "[1,2]"},
		{7, []string{"a", "b"}, "null", `["a","b"]`},
		{8, 1, "[]", "1"},
		{9, "abc", "[]", `"abc"`},
	}
	for _, testCase := range testCases {
		s := ToJson(testCase.Value, testCase.DefaultValue)
		if !strings.EqualFold(s, testCase.Except) {
			t.Errorf("%d %#v except: %s actual: %s", testCase.Number, testCase.Value, testCase.Except, s)
		}
	}
}

func TestEmptyObject(t *testing.T) {
	result := "{}"
	if b, err := EmptyObjectRawMessage().MarshalJSON(); err == nil {
		eValue := string(b)
		if !strings.EqualFold(eValue, result) {
			t.Errorf("Excepted value: %s, actual value: %s", eValue, result)
		}
	} else {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestEmptyArray(t *testing.T) {
	result := "[]"
	if b, err := EmptyArrayRawMessage().MarshalJSON(); err == nil {
		eValue := string(b)
		if !strings.EqualFold(eValue, result) {
			t.Errorf("Excepted value: %s, actual value: %s", eValue, result)
		}
	} else {
		t.Errorf("Error: %s", err.Error())
	}
}

func TestIsEmpty(t *testing.T) {
	type testCase struct {
		Value json.RawMessage
		Empty bool
	}
	v1, _ := ToRawMessage([]string{}, "[]")
	v2, _ := ToRawMessage([]string{"a", "b"}, "[]")
	v3, _ := ToRawMessage([]int{1, 2, 3}, "[]")
	v4, _ := ToRawMessage(struct {
		Name string
		Age  int
	}{"John", 10}, "{}")
	v5, _ := ToRawMessage(nil, "[]")
	a := json.RawMessage{}
	a.UnmarshalJSON([]byte("null"))
	b := json.RawMessage{}
	b.UnmarshalJSON([]byte(""))
	testCases := []testCase{
		{json.RawMessage{}, true},
		{EmptyObjectRawMessage(), true},
		{EmptyArrayRawMessage(), true},
		{v1, true},
		{v2, false},
		{v3, false},
		{v4, false},
		{v5, true},
		{a, true},
		{b, true},
	}

	for _, c := range testCases {
		v := IsEmptyRawMessage(c.Value)
		if v != c.Empty {
			t.Errorf("except: %v, actual: %v", c.Empty, v)
		}
	}
}
