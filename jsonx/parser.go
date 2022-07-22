package jsonx

import (
	"encoding/json"
	"github.com/hiscaler/gox/bytex"
	"github.com/hiscaler/gox/stringx"
	"reflect"
	"strconv"
	"strings"
)

// Parser is a json string parse helper and not required define struct.
// You can use Find() method get the path value, and convert to string, int, int64, float32, float64, bool value.
// And you can use Exists() method check path is exists
// Usage:
// parser := jsonx.NewParser("[0,1,2]")
// parser.Find("1").Int() // Return 1, founded
// parser.Find("10", 0).Int() // Return 0 because not found, you give a default value 0

type Parser struct {
	data  reflect.Value
	value reflect.Value
}

type ParseFinder Parser

func (pf ParseFinder) Interface() interface{} {
	return pf.value.Interface()
}

func (pf ParseFinder) String() string {
	switch pf.value.Kind() {
	case reflect.Invalid:
		return ""
	default:
		return stringx.String(pf.value.Interface())
	}
}

func (pf ParseFinder) Float32() float32 {
	switch pf.value.Kind() {
	case reflect.Invalid:
		return 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float32(pf.value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float32(pf.value.Uint())
	case reflect.Float32, reflect.Float64:
		return float32(pf.value.Float())
	case reflect.Bool:
		if pf.value.Bool() {
			return 1
		}
		return 0
	case reflect.String:
		d, err := strconv.ParseFloat(pf.value.String(), 32)
		if err != nil {
			return 0
		}
		return float32(d)
	default:
		return 0
	}
}

func (pf ParseFinder) Float64() float64 {
	switch pf.value.Kind() {
	case reflect.Invalid:
		return 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(pf.value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float64(pf.value.Uint())
	case reflect.Float32, reflect.Float64:
		return pf.value.Float()
	case reflect.Bool:
		if pf.value.Bool() {
			return 1
		}
		return 0
	case reflect.String:
		d, _ := strconv.ParseFloat(pf.value.String(), 64)
		return d
	default:
		return 0
	}
}

func (pf ParseFinder) Int() int {
	switch pf.value.Kind() {
	case reflect.Invalid:
		return 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(pf.value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return int(pf.value.Uint())
	case reflect.Float32, reflect.Float64:
		return int(pf.value.Float())
	case reflect.Bool:
		if pf.value.Bool() {
			return 1
		}
		return 0
	case reflect.String:
		d, _ := strconv.Atoi(pf.value.String())
		return d
	default:
		return 0
	}
}

func (pf ParseFinder) Int64() int64 {
	switch pf.value.Kind() {
	case reflect.Invalid:
		return 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return pf.value.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return int64(pf.value.Uint())
	case reflect.Float32, reflect.Float64:
		return int64(pf.value.Float())
	case reflect.Bool:
		if pf.value.Bool() {
			return 1
		}
		return 0
	case reflect.String:
		d, err := strconv.ParseInt(pf.value.String(), 10, 64)
		if err != nil {
			return 0
		}
		return d
	default:
		return 0
	}
}

func (pf ParseFinder) Bool() bool {
	switch pf.value.Kind() {
	case reflect.Invalid:
		return false
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return pf.value.Int() > 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return pf.value.Uint() > 0
	case reflect.Float32, reflect.Float64:
		return pf.value.Float() > 0
	case reflect.Bool:
		return pf.value.Bool()
	case reflect.String:
		v, _ := strconv.ParseBool(pf.value.String())
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
	if !p.data.IsValid() || path == "" {
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

func (p *Parser) Find(path string, defaultValue ...interface{}) *ParseFinder {
	if len(defaultValue) > 0 {
		p.value = reflect.ValueOf(defaultValue[0])
	}
	if !p.data.IsValid() || path == "" {
		return (*ParseFinder)(p)
	}

	data := p.data
	// find the value corresponding to the path
	// if any part of path cannot be located, return the default value
	parts := strings.Split(path, ".")
	n := len(parts)
	for i := 0; i < n; i++ {
		if data = getElement(data, parts[i]); !data.IsValid() {
			return (*ParseFinder)(p)
		}
		if i == n-1 {
			// is last path
			p.value = data
		}
	}
	return (*ParseFinder)(p)
}
