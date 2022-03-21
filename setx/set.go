package setx

import (
	"github.com/hiscaler/gox/inx"
	"strings"
)

func ToStringSet(values []string, ignoreCase bool) []string {
	if len(values) <= 1 {
		return values
	}

	sets := make([]string, 0)
	m := make(map[string]struct{}, 0)
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value != "" {
			fixedValue := value
			if ignoreCase {
				fixedValue = strings.ToLower(fixedValue)
			}
			if _, ok := m[fixedValue]; !ok {
				m[fixedValue] = struct{}{}
				sets = append(sets, value)
			}
		}
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
