package slicex

// ToInterfaceSlice Change values to interface slice
func ToInterfaceSlice(values ...interface{}) []interface{} {
	is := make([]interface{}, len(values))
	for i, value := range values {
		is[i] = value
	}

	return is
}
