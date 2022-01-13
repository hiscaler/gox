package slicex

import "strings"

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
