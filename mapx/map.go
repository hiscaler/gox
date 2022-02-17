package mapx

import "sort"

// StringKeys 获取排序后的 map 键值
func StringKeys(m map[string]interface{}) []string {
	keys := make([]string, 0)
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}
