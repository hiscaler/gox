package setx

import (
	"github.com/hiscaler/gox"
	"github.com/hiscaler/gox/inx"
	"strings"
)

// ToSet change slice to unique values
func ToSet[T gox.Number | string | bool | byte | rune](values []T) []T {
	if len(values) <= 1 {
		return values
	}

	uniqueValues := make([]T, 0)
	kv := make(map[T]struct{}, len(values))
	for _, value := range values {
		if _, ok := kv[value]; !ok {
			kv[value] = struct{}{}
			uniqueValues = append(uniqueValues, value)
		}
	}
	return uniqueValues
}

func ToStringSet(values []string, caseSensitive bool) []string {
	if len(values) <= 1 {
		return values
	}

	m := make(map[string]string, 0)
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value != "" {
			fixedValue := value
			if !caseSensitive {
				fixedValue = strings.ToLower(fixedValue)
			}
			if _, ok := m[fixedValue]; !ok {
				m[fixedValue] = value
			}
		}
	}
	if len(m) == 0 {
		return nil
	}

	sets := make([]string, len(m))
	i := 0
	for _, v := range m {
		sets[i] = v
		i++
	}
	return sets
}

func ToIntSet(values []int) []int {
	if len(values) <= 1 {
		return values
	}

	sets := make([]int, 0)
	for _, value := range values {
		if !inx.IntIn(value, sets...) {
			sets = append(sets, value)
		}
	}
	return sets
}
