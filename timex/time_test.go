package timex

import (
	"testing"
	"time"
)

func TestIsAmericaSummerTime(t *testing.T) {
	testCases := []struct {
		Date       string
		SummerTime bool
	}{
		{"2021-11-10", false},
		{"2021-12-10", false},
		{"2021-03-10", false},
		{"2021-03-14", true},
		{"2021-11-01", true},
		{"2021-10-10", true},
		{"2021-10-11", true},
		{"2021-10-12", true},
	}
	for _, testCase := range testCases {
		d, _ := time.Parse("2006-01-02", testCase.Date)
		v := IsAmericaSummerTime(d)
		if v != testCase.SummerTime {
			t.Errorf("%s except %v, actual %v", testCase.Date, testCase.SummerTime, v)
		}
	}
}
