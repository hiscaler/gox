package pathx

import (
	"path"
	"strings"
)

func Filename(s string) string {
	if s == "" {
		return ""
	}

	filename := path.Base(s)
	if ext := path.Ext(s); ext != "" {
		filename = strings.TrimSuffix(filename, ext)
	}
	return filename
}
