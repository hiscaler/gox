package nullx

import (
	"gopkg.in/guregu/null.v4"
	"strings"
)

func StringFrom(s string) null.String {
	if s != "" {
		s = strings.TrimSpace(s)
	}
	if s == "" {
		return NullString()
	} else {
		return null.NewString(s, true)
	}
}

func NullString() null.String {
	return null.NewString("", false)
}
