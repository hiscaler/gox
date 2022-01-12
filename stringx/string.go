package stringx

import "strings"

// IsEmpty 判断字符串是否为空
func IsEmpty(s string) bool {
	if s == "" || strings.TrimSpace(s) == "" {
		return true
	}
	return false
}

// ToNumber 字符串转换为唯一数字
// https://stackoverflow.com/questions/5459436/how-can-i-generate-a-unique-int-from-a-unique-string
func ToNumber(s string) int {
	number := 0
	runes := []rune(s)
	for i, r := range runes {
		x := 0
		if i != 0 {
			x = int(runes[i-1])
		}
		number += ((x << 16) | (x >> 16)) ^ int(r)
	}
	return number
}
