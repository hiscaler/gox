package isx

import (
	"bytes"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// Number Check string is a number
func Number(s string) bool {
	s = strings.TrimSpace(s)
	n := len(s)
	if n == 0 {
		return false
	}

	isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
	if strings.IndexFunc(s[0:1], isNotDigit) != -1 ||
		(n > 1 && strings.IndexFunc(s[n-1:], isNotDigit) != -1) {
		return false
	}

	return regexp.MustCompile(`^\d+$|^\d+[.]\d+$`).MatchString(strings.ReplaceAll(s, ",", ""))
}

// Empty 判断是否为空
func Empty(value interface{}) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Map, reflect.Slice:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Invalid:
		return true
	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			return true
		}
		return Empty(v.Elem().Interface())
	case reflect.Struct:
		v, ok := value.(time.Time)
		if ok && v.IsZero() {
			return true
		}
	}

	return false
}

func Equal(expected interface{}, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	if exp, ok := expected.([]byte); ok {
		act, ok := actual.([]byte)
		if !ok {
			return false
		}

		if exp == nil || act == nil {
			return true
		}

		return bytes.Equal(exp, act)
	}

	return reflect.DeepEqual(expected, actual)

}
