package bytex

import (
	"bytes"
	"unsafe"
)

// IsEmpty Check byte is empty
func IsEmpty(b []byte) bool {
	if len(b) == 0 || len(bytes.TrimSpace(b)) == 0 {
		return true
	}
	return false
}

func ToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
