package slicex

import (
	"strings"
)

// StringToInterface Change string slice to interface slice
func StringToInterface(values []string) []interface{} {
	if values == nil {
		return make([]interface{}, 0)
	}
	is := make([]interface{}, len(values))
	for i, value := range values {
		is[i] = value
	}
	return is
}

// IntToInterface Change int slice to interface slice
func IntToInterface(values []int) []interface{} {
	is := make([]interface{}, len(values))
	for i, value := range values {
		is[i] = value
	}
	return is
}

// StringSliceEqual Check a, b is equal
func StringSliceEqual(a, b []string, ignoreCase, ignoreEmpty, trim bool) bool {
	if ignoreCase || ignoreEmpty || trim {
		fixFunc := func(ss []string) []string {
			if len(ss) == 0 {
				return ss
			}
			values := make([]string, 0)
			for _, s := range ss {
				if trim {
					s = strings.TrimSpace(s)
				}
				if s == "" && ignoreEmpty {
					continue
				}
				if ignoreCase {
					s = strings.ToUpper(s)
				}
				values = append(values, s)
			}
			return values
		}
		a = fixFunc(a)
		b = fixFunc(b)
	}
	if len(a) != len(b) {
		return false
	}

	for _, av := range a {
		exists := false
		for _, bv := range b {
			if av == bv {
				exists = true
				break
			}
		}
		if !exists {
			return false
		}
	}
	return true
}

// IntSliceEqual Check a, b is equal
func IntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for _, av := range a {
		exists := false
		for _, bv := range b {
			if av == bv {
				exists = true
				break
			}
		}
		if !exists {
			return false
		}
	}
	return true
}

func StringSliceReverse(ss []string) []string {
	n := len(ss)
	if n > 1 {
		for k1 := 0; k1 < n/2; k1++ {
			k2 := n - k1 - 1
			ss[k1], ss[k2] = ss[k2], ss[k1]
		}
	}
	return ss
}

func IntSliceReverse(ss []int) []int {
	n := len(ss)
	if n > 1 {
		for k1 := 0; k1 < n/2; k1++ {
			k2 := n - k1 - 1
			ss[k1], ss[k2] = ss[k2], ss[k1]
		}
	}
	return ss
}

func StringSliceDiff(ss ...[]string) []string {
	diffValues := make([]string, 0)
	if len(ss) == 1 {
		diffValues = ss[0]
	} else if len(ss) > 1 {
		for _, v1 := range ss[0] {
			exists := false
			for _, items := range ss[1:] {
				for _, v2 := range items {
					if strings.EqualFold(v1, v2) {
						exists = true
						break
					}
				}
				if exists {
					break
				}
			}
			if !exists {
				diffValues = append(diffValues, v1)
			}
		}
	}
	return diffValues
}

func IntSliceDiff(ss ...[]int) []int {
	diffValues := make([]int, 0)
	if len(ss) == 1 {
		diffValues = ss[0]
	} else if len(ss) > 1 {
		for _, v1 := range ss[0] {
			exists := false
			for _, items := range ss[1:] {
				for _, v2 := range items {
					if v1 == v2 {
						exists = true
						break
					}
				}
				if exists {
					break
				}
			}
			if !exists {
				diffValues = append(diffValues, v1)
			}
		}
	}
	return diffValues
}
