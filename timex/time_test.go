package timex

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIsAmericaSummerTime(t *testing.T) {
	testCases := []struct {
		Date       string
		SummerTime bool
	}{
		{"0001-01-01", false},
		{"2021-11-10", false},
		{"2021-12-10", false},
		{"2021-03-10", false},
		{"2021-03-14", true},
		{"2021-11-01", true},
		{"2021-10-10", true},
		{"2021-10-11", true},
		{"2021-10-12", true},
		{"2021-12-12", false},
	}
	for _, testCase := range testCases {
		d, _ := time.Parse("2006-01-02", testCase.Date)
		v := IsAmericaSummerTime(d)
		if v != testCase.SummerTime {
			t.Errorf("%s except %v, actual %v", testCase.Date, testCase.SummerTime, v)
		}
	}
}

func TestBetween(t *testing.T) {
	testCases := []struct {
		tag      string
		t        string
		begin    string
		end      string
		expected bool
	}{
		{"t1", "2022-01-01", "2022-01-01", "2022-01-01", true},
		{"t2", "2022-01-02", "2022-01-01", "2022-01-01", false},
		{"t2", "2022-01-02", "2022-01-01", "2022-01-02", true},
	}
	format := "2006-01-02"
	for _, testCase := range testCases {
		tv, _ := time.Parse(format, testCase.t)
		begin, _ := time.Parse(format, testCase.begin)
		end, _ := time.Parse(format, testCase.end)
		v := Between(tv, begin, end)
		assert.Equal(t, testCase.expected, v, testCase.tag)
	}
}
