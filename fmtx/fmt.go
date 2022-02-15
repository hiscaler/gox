package fmtx

import (
	"encoding/json"
	"fmt"
)

func toJson(prefix string, data interface{}) string {
	s := ""
	if b, err := json.MarshalIndent(data, "", "    "); err == nil {
		s = string(b)
	} else {
		s = fmt.Sprintf("%#v", data)
	}
	if prefix != "" {
		s = fmt.Sprintf("%s: %s", prefix, s)
	}
	return s
}

func SprettyPrint(data interface{}) string {
	return toJson("", data)
}

func PrettyPrint(prefix string, data interface{}) {
	fmt.Print(toJson(prefix, data))
}

func PrettyPrintln(prefix string, data interface{}) {
	fmt.Println(toJson(prefix, data))
}
