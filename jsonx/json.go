package jsonx

import (
	"bytes"
	"encoding/json"
	"strings"
)

func ToRawMessage(i interface{}, defaultValue string) (json.RawMessage, error) {
	m := json.RawMessage{}
	var b []byte
	var err error
	if b, err = json.Marshal(&i); err == nil {
		sb := strings.TrimSpace(string(b))
		if sb == "null" || sb == "" {
			b = []byte(defaultValue)
		}
		if err = m.UnmarshalJSON(b); err != nil {
			m = json.RawMessage{}
		}
	}
	return m, err
}

// ToJson Change interface to json string
func ToJson(i interface{}, defaultValue string) string {
	b, err := json.Marshal(i)
	if err != nil {
		return defaultValue
	}
	var buf bytes.Buffer
	err = json.Compact(&buf, b)
	if err != nil {
		return defaultValue
	}

	return buf.String()
}

// EmptyObjectRawMessage 空对象
func EmptyObjectRawMessage() json.RawMessage {
	v := json.RawMessage{}
	v.UnmarshalJSON([]byte("{}"))
	return v
}

// EmptyArrayRawMessage 空数组
func EmptyArrayRawMessage() json.RawMessage {
	v := json.RawMessage{}
	v.UnmarshalJSON([]byte("[]"))
	return v
}

// IsEmptyRawMessage 验证数据是否为空
func IsEmptyRawMessage(data json.RawMessage) bool {
	if data == nil {
		return true
	}
	b, err := data.MarshalJSON()
	if err == nil {
		s := string(bytes.TrimSpace(b))
		if s == "" || s == "[]" || s == "{}" || strings.EqualFold(s, "null") {
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}
