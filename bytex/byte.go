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

func StartsWith(s []byte, ss [][]byte, caseSensitive bool) bool {
	if ss == nil || len(ss) == 0 {
		return true
	}

	has := false
	if !caseSensitive {
		s = bytes.ToLower(s)
	}
	for _, prefix := range ss {
		if len(prefix) == 0 {
			has = true
		} else {
			if !caseSensitive {
				prefix = bytes.ToLower(prefix)
			}
			has = bytes.HasPrefix(s, prefix)
		}
		if has {
			break
		}
	}
	return has
}

func EndsWith(s []byte, ss [][]byte, caseSensitive bool) bool {
	if ss == nil || len(ss) == 0 {
		return true
	}

	has := false
	if !caseSensitive {
		s = bytes.ToLower(s)
	}
	for _, suffix := range ss {
		if len(suffix) == 0 {
			has = true
		} else {
			if !caseSensitive {
				suffix = bytes.ToLower(suffix)
			}
			has = bytes.HasSuffix(s, suffix)
		}
		if has {
			break
		}
	}
	return has
}
