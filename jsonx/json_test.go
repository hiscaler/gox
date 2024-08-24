package jsonx

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestToJson(t *testing.T) {
	var names []string
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
		{5, nil, "abc", "abc"},
		{6, []int{1, 2}, "null", "[1,2]"},
		{7, []string{"a", "b"}, "null", `["a","b"]`},
		{8, 1, "[]", "1"},
		{9, "abc", "[]", `"abc"`},
		{10, nil, "[]", `[]`},
		{11, names, "[]", `[]`},
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

func TestIsEmptyRawMessage(t *testing.T) {
	type testCase struct {
		Number int
		Value  json.RawMessage
		Empty  bool
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
	c := json.RawMessage{}
	c.UnmarshalJSON([]byte("[    ]"))
	testCases := []testCase{
		{1, json.RawMessage{}, true},
		{2, EmptyObjectRawMessage(), true},
		{3, EmptyArrayRawMessage(), true},
		{4, v1, true},
		{5, v2, false},
		{6, v3, false},
		{7, v4, false},
		{8, v5, true},
		{9, a, true},
		{10, b, true},
		{11, c, true},
	}

	for _, tc := range testCases {
		v := IsEmptyRawMessage(tc.Value)
		if v != tc.Empty {
			t.Errorf("%d except: %v, actual: %v", tc.Number, tc.Empty, v)
		}
	}
}

func TestConvert(t *testing.T) {
	testCases := []struct {
		Number int
		From   json.RawMessage
		Except any
	}{
		{1, nil, struct{}{}},
		{2, EmptyArrayRawMessage(), struct{}{}},
		{3, []byte(`{"ID":1,"Name":"hiscaler"}`), struct {
			ID   int
			Name string
		}{}},
		{4, []byte(`{"ID":1,"Name":"hiscaler","age":1}`), struct {
			ID   int
			Name string
			age  int
		}{}},
	}
	for _, testCase := range testCases {
		exceptValue := testCase.Except
		err := Convert(testCase.From, &exceptValue)
		assert.Equalf(t, nil, err, "Test %d", testCase.Number)
		actualValue := ""
		if testCase.From != nil {
			actualValue = ToJson(exceptValue, "null")
		}
		t.Logf(`
#%d %s
    ↓
    %#v`, testCase.Number, testCase.From, exceptValue)
		assert.Equalf(t, string(testCase.From), actualValue, "Test %d", testCase.Number)
	}
}
