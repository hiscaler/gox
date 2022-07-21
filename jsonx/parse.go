package jsonx

import (
	"encoding/json"
	"github.com/hiscaler/gox/bytex"
	"github.com/hiscaler/gox/stringx"
	"reflect"
	"strconv"
	"strings"
)

type Parser struct {
	data  reflect.Value
	value reflect.Value
}

func (p Parser) ToString() string {
	switch p.value.Kind() {
	case reflect.Invalid:
		return ""
	default:
		return stringx.String(p.value.Interface())
	}
}

func (p Parser) ToFloat32() float32 {
	switch p.value.Kind() {
	case reflect.Invalid:
		return 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float32(p.value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float32(p.value.Uint())
	case reflect.Float32, reflect.Float64:
		return float32(p.value.Float())
	case reflect.Bool:
		if p.value.Bool() {
			return 1
		}
		return 0
	case reflect.String:
		d, err := strconv.ParseFloat(p.value.String(), 32)
		if err != nil {
			return 0
		}
		return float32(d)
	default:
		return 0
	}
}

func (p Parser) ToFloat64() float64 {
	switch p.value.Kind() {
	case reflect.Invalid:
		return 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(p.value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float64(p.value.Uint())
	case reflect.Float32, reflect.Float64:
		return p.value.Float()
	case reflect.Bool:
		if p.value.Bool() {
			return 1
		}
		return 0
	case reflect.String:
		d, _ := strconv.ParseFloat(p.value.String(), 64)
		return d
	default:
		return 0
	}
}

func (p Parser) ToInt() int {
	switch p.value.Kind() {
	case reflect.Invalid:
		return 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(p.value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return int(p.value.Uint())
	case reflect.Float32, reflect.Float64:
		return int(p.value.Float())
	case reflect.Bool:
		if p.value.Bool() {
			return 1
		}
		return 0
	case reflect.String:
		d, _ := strconv.Atoi(p.value.String())
		return d
	default:
		return 0
	}
}

func (p Parser) ToInt64() int64 {
	switch p.value.Kind() {
	case reflect.Invalid:
		return 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return p.value.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return int64(p.value.Uint())
	case reflect.Float32, reflect.Float64:
		return int64(p.value.Float())
	case reflect.Bool:
		if p.value.Bool() {
			return 1
		}
		return 0
	case reflect.String:
		d, err := strconv.Atoi(p.value.String())
		if err != nil {
			return 0
		}
		return int64(d)
	default:
		return 0
	}
}

func (p Parser) ToBool() bool {
	switch p.value.Kind() {
	case reflect.Invalid:
		return false
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return p.value.Int() > 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return p.value.Uint() > 0
	case reflect.Float32, reflect.Float64:
		return p.value.Float() > 0
	case reflect.Bool:
		return p.value.Bool()
	case reflect.String:
		v, _ := strconv.ParseBool(p.value.String())
		return v
	default:
		return false
	}
}

func getElement(v reflect.Value, p string) reflect.Value {
	switch v.Kind() {
	case reflect.Map:
		vv := v.MapIndex(reflect.ValueOf(p))
		if vv.Kind() == reflect.Interface {
			vv = vv.Elem()
		}
		return vv
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

func NewParser(s string) *Parser {
	p := &Parser{}
	return p.LoadString(s)
}

func (p *Parser) LoadString(s string) *Parser {
	if s == "" {
		return p
	}
	return p.LoadBytes(stringx.ToBytes(s))
}

func (p *Parser) LoadBytes(bytes []byte) *Parser {
	if bytex.IsBlank(bytes) {
		return p
	}

	var sd interface{}
	if err := json.Unmarshal(bytes, &sd); err != nil {
		return p
	}
	p.data = reflect.ValueOf(sd)
	return p
}

func (p Parser) Exists(path string) bool {
	if !p.data.IsValid() {
		return false
	}

	data := p.data
	parts := strings.Split(path, ".")
	n := len(parts)
	for i := 0; i < n; i++ {
		if data = getElement(data, parts[i]); !data.IsValid() {
			return false
		}
		if i == n-1 {
			// is last path
			return true
		}
	}
	return false
}

func (p *Parser) Find(path string, defaultValue ...interface{}) *Parser {
	if len(defaultValue) > 0 {
		p.value = reflect.ValueOf(defaultValue[0])
	}
	if !p.data.IsValid() {
		return p
	}

	data := p.data
	// find the value corresponding to the path
	// if any part of path cannot be located, return the default value
	parts := strings.Split(path, ".")
	n := len(parts)
	for i := 0; i < n; i++ {
		if data = getElement(data, parts[i]); !data.IsValid() {
			return p
		}
		if i == n-1 {
			// is last path
			p.value = data
		}
	}
	return p
}
