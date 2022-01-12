package isx

import "strings"

// Number Check string is a number
func Number(s string) bool {
	isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
	return strings.IndexFunc(s, isNotDigit) == -1
}
