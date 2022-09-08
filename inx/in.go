package inx

import (
	"strings"
)

// In Check value in values, return true if in values, otherwise return false.
// Value T is a generic value
func In[T comparable](value T, values []T) bool {
	if values == nil || len(values) == 0 {
		return false
	}
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

// StringIn 判断 s 是否在 ss 中（忽略大小写）
func StringIn(s string, ss ...string) bool {
	if len(ss) == 0 {
		return false
	}
	for _, s2 := range ss {
		if strings.EqualFold(s, s2) {
			return true
		}
	}
	return false
}

// IntIn 判断 i 是否在 ii 中
func IntIn(i int, ii ...int) bool {
	return In(i, ii)
}
