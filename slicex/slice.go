package slicex

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
