package mapx

import "sort"

// StringKeys 获取排序后的 map 键值
func StringKeys(m map[string]interface{}) []string {
	if m == nil || len(m) == 0 {
		return []string{}
	}

	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	if i > 1 {
		sort.Strings(keys)
	}
	return keys
}
