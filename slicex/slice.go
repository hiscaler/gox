package slicex

// StringSliceToInterfaceSlice Change string slice to interface slice
func StringSliceToInterfaceSlice(values []string) []interface{} {
	if values == nil {
		return make([]interface{}, 0)
	}
	is := make([]interface{}, len(values))
	for i, value := range values {
		is[i] = value
	}

	return is
}

// IntSliceToInterfaceSlice Change int slice to interface slice
func IntSliceToInterfaceSlice(values []int) []interface{} {
	is := make([]interface{}, len(values))
	for i, value := range values {
		is[i] = value
	}

	return is
}
