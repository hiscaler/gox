package keyx

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// Generate 生成 Key
func Generate(values ...interface{}) string {
	var sb strings.Builder
	for _, value := range values {
		v := reflect.ValueOf(value)
		switch v.Kind() {
		case reflect.String:
			if v.Len() != 0 {
				sb.WriteString(v.String())
			}
		case reflect.Bool:
			sb.WriteString(strconv.FormatBool(v.Bool()))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if b, ok := value.(rune); ok {
				sb.WriteRune(b)
			} else {
				sb.WriteString(strconv.FormatInt(v.Int(), 10))
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if b, ok := value.(byte); ok {
				sb.WriteByte(b)
			} else {
				sb.WriteString(strconv.FormatUint(v.Uint(), 10))
			}
		case reflect.Float32, reflect.Float64:
			sb.WriteString(strconv.FormatFloat(v.Float(), 'f', 2, 64))
		case reflect.Map:
			keys := make([]string, len(v.MapKeys()))
			i := 0
			for _, mv := range v.MapKeys() {
				keys[i] = mv.String()
				i++
			}
			sort.Strings(keys)
			interfaces := make([]interface{}, 0)
			for k := range keys {
				interfaces = append(interfaces, keys[k], v.MapIndex(reflect.ValueOf(keys[k])).Interface())
			}
			sb.WriteString(Generate(interfaces...))
		case reflect.Slice, reflect.Array:
			interfaces := make([]interface{}, 0)
			for i := 0; i < v.Len(); i++ {
				interfaces = append(interfaces, v.Index(i).Interface())
			}
			sb.WriteString(Generate(interfaces...))
		case reflect.Struct:
			kv := map[string]interface{}{}
			t := reflect.TypeOf(value)
			if t.Name() != "" {
				sb.WriteString(t.Name() + ":")
			}
			for k := 0; k < t.NumField(); k++ {
				kv[t.Field(k).Name] = v.Field(k).Interface()
			}
			sb.WriteString(Generate(kv))
		default:
			sb.WriteString(v.String())
		}
	}
	return sb.String()
}
