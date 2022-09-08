package slicex

import (
	"github.com/hiscaler/gox"
	"strings"
)

// Map values all value execute f function, and return a new slice
//
// Example
//	Map([]string{"A", "B", "C"}, func(v string) string {
//		return strings.ToLower(v)
//	})
//	// Output: ["a", "b", "c"]
func Map[T comparable](values []T, f func(v T) T) []T {
	items := make([]T, len(values))
	for i, v := range values {
		items[i] = f(v)
	}
	return items
}

// Filter return matched f function condition value
//
// Example:
//	Filter([]int{0, 1, 2, 3}, func(v int) bool {
//		return v > 0
//	})
//	// Output: [1, 2, 3]
func Filter[T comparable](values []T, f func(v T) bool) []T {
	items := make([]T, 0)
	for _, v := range values {
		if ok := f(v); ok {
			items = append(items, v)
		}
	}
	return items
}

func ToInterface[T gox.Number | ~string](values []T) []interface{} {
	if values == nil || len(values) == 0 {
		return []interface{}{}
	}

	ifs := make([]interface{}, len(values))
	for i, value := range values {
		ifs[i] = value
	}
	return ifs
}

// StringToInterface Change string slice to interface slice
func StringToInterface(values []string) []interface{} {
	return ToInterface(values)
}

// IntToInterface Change int slice to interface slice
func IntToInterface(values []int) []interface{} {
	return ToInterface(values)
}

// StringSliceEqual Check a, b is equal
func StringSliceEqual(a, b []string, caseSensitive, ignoreEmpty, trim bool) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}

	if !caseSensitive || ignoreEmpty || trim {
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
				if !caseSensitive {
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
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil || len(a) != len(b) {
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
	if n <= 1 {
		return ss
	}

	vv := make([]string, len(ss))
	copy(vv, ss)
	for k1 := 0; k1 < n/2; k1++ {
		k2 := n - k1 - 1
		vv[k1], vv[k2] = vv[k2], vv[k1]
	}
	return vv
}

func IntSliceReverse(ss []int) []int {
	n := len(ss)
	if n <= 1 {
		return ss
	}

	vv := make([]int, len(ss))
	copy(vv, ss)
	for k1 := 0; k1 < n/2; k1++ {
		k2 := n - k1 - 1
		vv[k1], vv[k2] = vv[k2], vv[k1]
	}
	return vv
}

// Diff return a slice in ss[0] and not in ss[1:]
func Diff[T comparable](values ...[]T) []T {
	diffValues := make([]T, 0)
	n := len(values)
	if n == 0 || values[0] == nil {
		return diffValues
	} else if n == 1 {
		return values[0]
	} else {
		items := make(map[T]struct{}, 0)
		for _, vs := range values[1:] {
			for _, v := range vs {
				items[v] = struct{}{}
			}
		}
		for _, v := range values[0] {
			if _, ok := items[v]; !ok {
				diffValues = append(diffValues, v)
			}
		}
	}
	return diffValues
}

func StringSliceDiff(ss ...[]string) []string {
	diffValues := make([]string, 0)
	if len(ss) == 0 || ss[0] == nil {
		return diffValues
	} else if len(ss) == 1 {
		return ss[0]
	} else {
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
	return Diff(ss...)
}

// Chunk chunks a slice by size
func Chunk[T comparable](items []T, size int) [][]T {
	chunkItems := make([][]T, 0)
	n := len(items)
	if items == nil || n == 0 {
		return chunkItems
	} else if size <= 0 {
		return [][]T{items}
	}
	for i := 0; i < n; i += size {
		end := i + size
		if end > n {
			end = n
		}
		chunkItems = append(chunkItems, items[i:end])
	}
	return chunkItems
}
