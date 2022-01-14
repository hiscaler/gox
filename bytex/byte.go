package bytex

import (
	"bytes"
)

// IsEmpty Check byte is empty
func IsEmpty(b []byte) bool {
	if len(b) == 0 || len(bytes.TrimSpace(b)) == 0 {
		return true
	}
	return false
}
