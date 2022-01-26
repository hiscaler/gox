package urlx

import (
	"net/url"
	"strings"
)

type URL struct {
	Path    string     // URL path
	URL     *url.URL   // A url.URL represents
	Invalid bool       // Path is a valid url
	values  url.Values // Query values
}

func NewURL(path string) *URL {
	u := &URL{
		Path:    path,
		Invalid: false,
		values:  url.Values{},
	}
	if v, err := url.Parse(u.Path); err == nil {
		u.URL = v
		u.Invalid = true
		if values, err := url.ParseQuery(v.RawQuery); err == nil {
			u.values = values
		}
	}

	return u
}

func (u URL) GetValue(key, defaultValue string) string {
	v := u.values.Get(key)
	if v == "" {
		v = defaultValue
	}
	return v
}

func (u URL) SetValue(key, value string) URL {
	u.values.Set(key, value)
	return u
}

func (u URL) AddValue(key, value string) URL {
	u.values.Add(key, value)
	return u
}

func (u URL) DelKey(key string) URL {
	u.values.Del(key)
	return u
}

func (u URL) HasKey(key string) bool {
	return u.values.Has(key)
}

func (u URL) String() string {
	s := u.URL.String()
	rawQuery := u.URL.RawQuery
	if rawQuery == "" {
		if len(u.values) > 0 {
			s += "?" + u.values.Encode()
		}
	} else {
		s = strings.Replace(s, rawQuery, u.values.Encode(), 1)
	}
	return s
}
