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
