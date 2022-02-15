package fmtx

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func toJson(prefix string, data interface{}) string {
	s := ""
	if b, err := json.MarshalIndent(data, "", "    "); err == nil {
		s = string(b)
	} else {
		s = fmt.Sprintf("%#v", data)
	}
	if prefix != "" {
		s = fmt.Sprintf(`%s
%s`, prefix, s)
	}
	return s
}

func SprettyPrint(a ...interface{}) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	values := make([]string, n)
	for _, v := range a {
		values = append(values, toJson("", v))
	}
	return strings.Join(values, "\n")
}

func PrettyPrint(prefix string, a ...interface{}) {
	onlyOne := len(a) == 1
	for k, v := range a {
		p := prefix
		if p == "" {
			if !onlyOne {
				p = strconv.Itoa(k + 1)
			}
		} else {
			if !onlyOne {
				p = fmt.Sprintf("%s %d", prefix, k+1)
			}
		}
		fmt.Print(toJson(prefix, v))
	}
}

func PrettyPrintln(prefix string, a ...interface{}) {
	onlyOne := len(a) == 1
	for k, v := range a {
		p := prefix
		if p == "" {
			if !onlyOne {
				p = strconv.Itoa(k + 1)
			}
		} else {
			if !onlyOne {
				p = fmt.Sprintf("%s %d", prefix, k+1)
			}
		}
		fmt.Println(toJson(p, v))
	}
}
