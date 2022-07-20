package jsonx

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

type Parse struct {
	data       string
	foundValue reflect.Value
}

func (p Parse) ToString() string {
	switch p.foundValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(p.foundValue.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(p.foundValue.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(p.foundValue.Float(), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(p.foundValue.Float(), 'f', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(p.foundValue.Bool())
	case reflect.Invalid:
		return ""
	default:
		return p.foundValue.String()
	}
}

func (p Parse) ToInt() int {
	switch p.foundValue.Kind() {
	case reflect.Invalid:
		return 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(p.foundValue.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return int(p.foundValue.Uint())
	case reflect.Bool:
		if p.foundValue.Bool() {
			return 1
		}
		return 0
	case reflect.String:
		d, _ := strconv.Atoi(p.foundValue.String())
		return d
	default:
		return 0
	}
}

func mapIndex(mp reflect.Value, index reflect.Value) reflect.Value {
	v := mp.MapIndex(index)
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v
}

func getElement(v reflect.Value, p string) reflect.Value {
	switch v.Kind() {
	case reflect.Map:
		return mapIndex(v, reflect.ValueOf(p))
	case reflect.Array, reflect.Slice:
		if i, err := strconv.Atoi(p); err == nil {
			if i >= 0 && i < v.Len() {
				v = v.Index(i)
				for v.Kind() == reflect.Interface {
					v = v.Elem()
				}
				return v
			}
		}
	}
	return reflect.Value{}
}

func Find(s string, path string, defaultValue ...interface{}) Parse {
	p := Parse{data: s}
	var d reflect.Value
	if len(defaultValue) > 0 {
		d = reflect.ValueOf(defaultValue[0])
	}

	var sd interface{}
	if err := json.Unmarshal([]byte(s), &sd); err != nil {
		return p
	}

	data := reflect.ValueOf(sd)
	// find the value corresponding to the path
	// if any part of path cannot be located, return the default value
	parts := strings.Split(path, ".")
	n := len(parts)
	for i := 0; i < n-1; i++ {
		if data = getElement(data, parts[i]); !data.IsValid() {
			p.foundValue = d
			return p
		}
	}
	v := getElement(data, parts[n-1])
	if !v.IsValid() {
		p.foundValue = d
		return p
	} else {
		p.foundValue = v
	}
	return p
}
