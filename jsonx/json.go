package jsonx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func ToRawMessage(i interface{}, defaultValue string) (json.RawMessage, error) {
	m := json.RawMessage{}
	var b []byte
	var err error
	b, err = json.Marshal(&i)
	if err != nil {
		return m, err
	}

	b = bytes.TrimSpace(b)
	if len(b) == 0 || bytes.EqualFold(b, []byte("null")) {
		b = []byte(defaultValue)
	}
	err = m.UnmarshalJSON(b)
	return m, err
}

// ToJson Change interface to json string
func ToJson(i interface{}, defaultValue string) string {
	if i == nil {
		return defaultValue
	}
	vo := reflect.ValueOf(i)
	switch vo.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		if vo.IsNil() {
			return defaultValue
		}
	default:
	}

	b, err := json.Marshal(i)
	if err != nil {
		return defaultValue
	}
	var buf bytes.Buffer
	err = json.Compact(&buf, b)
	if err != nil {
		return defaultValue
	}
	if json.Valid(buf.Bytes()) {
		return buf.String()
	}
	return defaultValue
}

func ToPrettyJson(i interface{}) string {
	if i == nil {
		return "null"
	}
	vo := reflect.ValueOf(i)
	switch vo.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		if vo.IsNil() {
			return "null"
		}
	default:
	}

	b, err := json.Marshal(i)
	if err != nil {
		return fmt.Sprintf("%+v", i)
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", i)
	}
	return buf.String()
}

// EmptyObjectRawMessage 空对象
func EmptyObjectRawMessage() json.RawMessage {
	v := json.RawMessage{}
	_ = v.UnmarshalJSON([]byte("{}"))

	return v
}

// EmptyArrayRawMessage 空数组
func EmptyArrayRawMessage() json.RawMessage {
	v := json.RawMessage{}
	_ = v.UnmarshalJSON([]byte("[]"))
	return v
}

// IsEmptyRawMessage 验证数据是否为空
func IsEmptyRawMessage(data json.RawMessage) bool {
	if data == nil {
		return true
	}

	b, err := data.MarshalJSON()
	if err != nil {
		return true
	}

	s := string(bytes.TrimSpace(b))
	if s == "" || s == "[]" || s == "{}" || strings.EqualFold(s, "null") {
		return true
	}

	if strings.Index(s, " ") != -1 {
		s = strings.ReplaceAll(s, " ", "")
	}
	return s == "[]" || s == "{}"
}

func Convert(from json.RawMessage, to any) error {
	if from == nil {
		return nil
	}

	var b []byte
	b, err := from.MarshalJSON()
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &to)
}
