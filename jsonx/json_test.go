package jsonx

import (
	"encoding/json"
	"strings"
	"testing"
)

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
