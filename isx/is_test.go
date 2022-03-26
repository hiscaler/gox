package isx

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNumber(t *testing.T) {
	testCases := []struct {
		Value    string
		IsNumber bool
	}{
		{"a", false},
		{"111", true},
		{"1.23", true},
		{"1,234.5", true},
		{"1234.5,", false},
		{"12345.", false},
		{" 12345.6   ", true},
		{" 12345. 6   ", false},
		{"-1", true},
		{"+1", true},
	}
	for _, testCase := range testCases {
		v := Number(testCase.Value)
		if v != testCase.IsNumber {
			t.Errorf("%s except %v actual %v", testCase.Value, testCase.IsNumber, v)
		}
	}
}

func TestEmpty(t *testing.T) {
	var s1 string
	var s2 = "a"
	var s3 *string
	s4 := struct{}{}
	time1 := time.Now()
	var time2 time.Time
	tests := []struct {
		tag   string
		value interface{}
		empty bool
	}{
		// nil
		{"t0", nil, true},
		// string
		{"t1.1", "", true},
		{"t1.2", "1", false},
		// slice
		{"t2.1", []byte(""), true},
		{"t2.2", []byte("1"), false},
		// map
		{"t3.1", map[string]int{}, true},
		{"t3.2", map[string]int{"a": 1}, false},
		// bool
		{"t4.1", false, true},
		{"t4.2", true, false},
		// int
		{"t5.1", 0, true},
		{"t5.2", int8(0), true},
		{"t5.3", int16(0), true},
		{"t5.4", int32(0), true},
		{"t5.5", int64(0), true},
		{"t5.6", 1, false},
		{"t5.7", int8(1), false},
		{"t5.8", int16(1), false},
		{"t5.9", int32(1), false},
		{"t5.10", int64(1), false},
		// uint
		{"t6.1", uint(0), true},
		{"t6.2", uint8(0), true},
		{"t6.3", uint16(0), true},
		{"t6.4", uint32(0), true},
		{"t6.5", uint64(0), true},
		{"t6.6", uint(1), false},
		{"t6.7", uint8(1), false},
		{"t6.8", uint16(1), false},
		{"t6.9", uint32(1), false},
		{"t6.10", uint64(1), false},
		// float
		{"t7.1", float32(0), true},
		{"t7.2", float64(0), true},
		{"t7.3", float32(1), false},
		{"t7.4", float64(1), false},
		// interface, ptr
		{"t8.1", &s1, true},
		{"t8.2", &s2, false},
		{"t8.3", s3, true},
		// struct
		{"t9.1", s4, false},
		{"t9.2", &s4, false},
		// time.Time
		{"t10.1", time1, false},
		{"t10.2", &time1, false},
		{"t10.3", time2, true},
		{"t10.4", &time2, true},
		// rune
		{"t11.1", 'a', false},
		// byte
		{"t12.1", []byte(""), true},
		{"t12.2", []byte(" "), false},
	}

	for _, test := range tests {
		empty := Empty(test.value)
		assert.Equal(t, test.empty, empty, test.tag)
	}
}

func TestIsEqual(t *testing.T) {
	s1 := "hello"
	s2 := s1
	s3 := "hello"
	t1 := time.Now()
	t2 := time.Now().AddDate(0, 0, 1)
	type1 := []struct {
		username string
	}{
		{"john"},
	}
	type2 := []struct {
		username string
	}{
		{"john"},
	}
	tests := []struct {
		tag    string
		a      interface{}
		b      interface{}
		except bool
	}{
		{"t0", nil, nil, true},
		{"t1", nil, "", false},
		{"t2", "", "", true},
		{"t3", "", " ", false},
		{"t4", s1, s2, true},
		{"t5", s2, s3, true},
		{"t6", t1, t2, false},
		{"t7", type1, type2, true},
	}

	for _, test := range tests {
		equal := Equal(test.a, test.b)
		assert.Equal(t, test.except, equal, test.tag)
	}
}

func TestSafeCharacters(t *testing.T) {
	type testCast struct {
		String string
		Safe   bool
	}
	testCasts := []testCast{
		{"", false},
		{" ", false},
		{"a", true},
		{"111", true},
		{"ａ", false},
		{"A_B", true},
		{"A_中B", false},
		{"a.b-c_", true},
		{"_.a.b-c_", true},
		{`\.a.b-c_`, false},
	}
	for _, tc := range testCasts {
		safe := SafeCharacters(tc.String)
		if safe != tc.Safe {
			t.Errorf("%s except %v, actual：%v", tc.String, tc.Safe, safe)
		}
	}
}

func TestHttpURL(t *testing.T) {
	tests := []struct {
		tag    string
		url    string
		except bool
	}{
		{"t0", "www.example.com", true},
		{"t1", "http://www.example.com", true},
		{"t2", "https://www.example.com", true},
		{"t3", "https://www.com", true},
		{"t4", "https://a", true}, // is valid URL?
		{"t5", "https://127.0.0.1", true},
		{"t6", "https://", false},
		{"t7", "https://a", true},
		{"t8", "", false},
		{"t9", "aaa", false},
		{"t10", "https://www.example.com:8080", true},
	}

	for _, test := range tests {
		equal := HttpURL(test.url)
		assert.Equal(t, test.except, equal, test.tag)
	}
}
