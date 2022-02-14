package isx

import (
	"regexp"
	"strings"
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
