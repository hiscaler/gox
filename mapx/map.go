package mapx

import (
	"net/url"
	"reflect"
	"sort"
	"strconv"
)

// Keys 获取 map 键值（默认按照升序排列）
func Keys(m interface{}) []string {
	var keys []string
	vo := reflect.ValueOf(m)
	if vo.Kind() == reflect.Map {
		mapKeys := vo.MapKeys()
		keys = make([]string, len(mapKeys))
		for k, v := range mapKeys {
			var vString string
			switch v.Type().Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				vString = strconv.FormatInt(v.Int(), 10)
			case reflect.Float32, reflect.Float64:
				vString = strconv.FormatFloat(v.Float(), 'f', -1, 64)
			case reflect.Bool:
				if v.Bool() {
					vString = "1"
				} else {
					vString = "0"
				}
			default:
				vString = v.String()
			}
			keys[k] = vString
		}
		if len(keys) > 0 {
			sort.Strings(keys)
		}
	}
	return keys
}

func StringMapStringEncode(params map[string]string) string {
	if len(params) == 0 {
		return ""
	}

	values := url.Values{}
	for k, v := range params {
		values.Add(k, v)
	}
	return values.Encode()
}
