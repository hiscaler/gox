package extractx

import (
	"regexp"
	"strconv"
	"strings"
)

var rxNumber = regexp.MustCompile(`\-?\d+[\d.,]*\d*`)

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
	if s = Number(s); s != "" {
		if v, err := strconv.ParseFloat(s, 64); err == nil {
			return v
		}
	}
	return 0
}

func Float32(s string) float32 {
	if s = Number(s); s != "" {
		if v, err := strconv.ParseFloat(s, 64); err == nil {
			return float32(v)
		}
	}
	return 0
}

func Int64(s string) int64 {
	if s = Number(s); s != "" {
		if v, err := strconv.ParseInt(s, 10, 64); err == nil {
			return v
		}
	}
	return 0
}

func Int32(s string) int32 {
	if s = Number(s); s != "" {
		if v, err := strconv.ParseInt(s, 10, 32); err == nil {
			return int32(v)
		}
	}
	return 0
}

func Int16(s string) int16 {
	if s = Number(s); s != "" {
		if v, err := strconv.ParseInt(s, 10, 16); err == nil {
			return int16(v)
		}
	}
	return 0
}

func Int8(s string) int8 {
	if s = Number(s); s != "" {
		if v, err := strconv.ParseInt(s, 10, 16); err == nil {
			return int8(v)
		}
	}
	return 0
}

func Int(s string) int {
	if s = Number(s); s != "" {
		if v, err := strconv.ParseInt(s, 10, 16); err == nil {
			return int(v)
		}
	}
	return 0
}
