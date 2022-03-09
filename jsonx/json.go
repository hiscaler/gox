package jsonx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func ToRawMessage(i interface{}, defaultValue string) (json.RawMessage, error) {
	m := json.RawMessage{}
	var b []byte
	var err error
	if b, err = json.Marshal(&i); err == nil {
		b = bytes.TrimSpace(b)
		if len(b) == 0 || bytes.EqualFold(b, []byte("null")) {
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
	if json.Valid(buf.Bytes()) {
		return buf.String()
	} else {
		return defaultValue
	}
}

func ToPrettyJson(i interface{}, defaultValue string) string {
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
			if strings.Index(s, " ") != -1 {
				s = strings.Replace(s, " ", "", -1)
			}
			return s == "[]" || s == "{}"
		}
	} else {
		return true
	}
}
