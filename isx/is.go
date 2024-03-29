package isx

import (
	"bytes"
	"github.com/hiscaler/gox/stringx"
	"net/url"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	rxSafeCharacters = regexp.MustCompile("^[a-zA-Z0-9\\.\\-_][a-zA-Z0-9\\.\\-_]*$")
	rxNumber         = regexp.MustCompile("^[+-]?\\d+$|^\\d+[.]\\d+$")
	rxColorHex       = regexp.MustCompile("^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$")
)

// OS type
const (
	IsAix       = "aix"
	IsAndroid   = "android"
	IsDarwin    = "darwin"
	IsDragonfly = "dragonfly"
	IsFreebsd   = "freebsd"
	IsHurd      = "hurd"
	IsIllumos   = "illumos"
	IsIos       = "ios"
	IsJs        = "js"
	IsLinux     = "linux"
	IsNacl      = "nacl"
	IsNetbsd    = "netbsd"
	IsOpenbsd   = "openbsd"
	IsPlan9     = "plan9"
	IsSolaris   = "solaris"
	IsWindows   = "windows"
	IsZos       = "zos"
)

// Number Check any value is a number
func Number(i interface{}) bool {
	switch i.(type) {
	case string:
		s := stringx.TrimAny(strings.TrimSpace(i.(string)), "+", "-")
		n := len(s)
		if n == 0 {
			return false
		}

		if strings.IndexFunc(s[n-1:], func(c rune) bool {
			return c < '0' || c > '9'
		}) != -1 {
			return false
		}
		return rxNumber.MatchString(strings.ReplaceAll(s, ",", ""))
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64,
		complex64, complex128:
		return true
	default:
		return false
	}
}

// Empty 判断是否为空
func Empty(value interface{}) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Map, reflect.Slice:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Invalid:
		return true
	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			return true
		}
		return Empty(v.Elem().Interface())
	case reflect.Struct:
		v, ok := value.(time.Time)
		if ok && v.IsZero() {
			return true
		}
	}
	return false
}

func Equal(expected interface{}, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	if exp, ok := expected.([]byte); ok {
		act, ok := actual.([]byte)
		if !ok {
			return false
		}

		if exp == nil || act == nil {
			return true
		}

		return bytes.Equal(exp, act)
	}
	return reflect.DeepEqual(expected, actual)
}

// SafeCharacters Only include a-zA-Z0-9.-_
// Reference https://www.quora.com/What-are-valid-file-names
func SafeCharacters(str string) bool {
	if str == "" {
		return false
	}
	return rxSafeCharacters.MatchString(str)
}

// HttpURL checks if the string is a HTTP URL.
// govalidator/IsURL
func HttpURL(str string) bool {
	const (
		URLSchema    string = `((https?):\/\/)`
		URLPath      string = `((\/|\?|#)[^\s]*)`
		URLPort      string = `(:(\d{1,5}))`
		URLIP        string = `([1-9]\d?|1\d\d|2[01]\d|22[0-3]|24\d|25[0-5])(\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-5]))`
		URLSubdomain string = `((www\.)|([a-zA-Z0-9]+([-_\.]?[a-zA-Z0-9])*[a-zA-Z0-9]\.[a-zA-Z0-9]+))`
		URL                 = `^` + URLSchema + `?` + `((` + URLIP + `|(\[` + `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))` + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + URLSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + URLPort + `?` + URLPath + `?$`
	)

	if str == "" || utf8.RuneCountInString(str) >= 2083 || len(str) <= 3 || strings.HasPrefix(str, ".") {
		return false
	}
	if strings.HasPrefix(str, "//") {
		str = "http:" + str
	}
	strTemp := str
	if strings.Contains(str, ":") && !strings.Contains(str, "://") {
		// support no indicated urlscheme but with colon for port number
		// http:// is appended so url.Parse will succeed, strTemp used so it does not impact rxURL.MatchString
		strTemp = "http://" + str
	}
	u, err := url.Parse(strTemp)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}
	return regexp.MustCompile(URL).MatchString(str)
}

// OS check typ is a valid OS type
// Usage: isx.OS(isx.IsLinux)
func OS(typ string) bool {
	return runtime.GOOS == typ
}

func ColorHex(s string) bool {
	return rxColorHex.MatchString(s)
}
