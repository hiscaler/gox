package extractx

import (
	"regexp"
	"strconv"
	"strings"
)

var rxNumber = regexp.MustCompile(`[\-]?\d+[\d.,]*\d*`)

// 提取的内容默认为 1,234.56 格式的数字，未实现根据国家标准实现提取
// https://zhuanlan.zhihu.com/p/157980325
func clean(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	s = strings.ReplaceAll(s, ",", "")
	n := len(s)
	if s[n-1:] == "." {
		s = s[n-2 : n-1]
	}
	return s
}

func Number(s string) string {
	if s == "" {
		return ""
	}
	return clean(rxNumber.FindString(s))
}

func Numbers(s string) []string {
	if s == "" {
		return []string{}
	}

	matches := rxNumber.FindAllString(s, -1)
	if matches == nil {
		return []string{}
	}

	for i, v := range matches {
		matches[i] = clean(v)
	}
	return matches
}

func Float64(s string) float64 {
	var v float64
	s = Number(s)
	if s != "" {
		if d, err := strconv.ParseFloat(s, 64); err == nil {
			v = d
		}
	}
	return v
}

func Float32(s string) float32 {
	return float32(Float64(s))
}

func Int64(s string) int64 {
	var v int64
	s = Number(s)
	if s != "" {
		if d, err := strconv.ParseInt(s, 10, 64); err == nil {
			v = d
		}
	}
	return v
}

func Int32(s string) int32 {
	return int32(Int64(s))
}

func Int16(s string) int16 {
	return int16(Int64(s))
}

func Int8(s string) int8 {
	return int8(Int64(s))
}

func Int(s string) int {
	return int(Int64(s))
}
