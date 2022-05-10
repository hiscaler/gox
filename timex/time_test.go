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
		{"2022-03-15", true},
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
		{"t3", "2022-01-02", "2022-01-01", "2022-01-02", true},
	}
	layout := "2006-01-02"
	for _, testCase := range testCases {
		tv, _ := time.Parse(layout, testCase.t)
		begin, _ := time.Parse(layout, testCase.begin)
		end, _ := time.Parse(layout, testCase.end)
		v := Between(tv, begin, end)
		assert.Equal(t, testCase.expected, v, testCase.tag)
	}
}

func TestDayStart(t *testing.T) {
	testCases := []struct {
		tag      string
		time     string
		layout   string
		expected string
	}{
		{"t1", "2022-01-01T12:12:00.924Z", "2006-01-02T15:04:05Z", "2022-01-01 00:00:00"},
		{"t2", "2022-01-01 00:00:00", "2006-01-02 15:04:05", "2022-01-01 00:00:00"},
	}

	for _, testCase := range testCases {
		tv, err := time.Parse(testCase.layout, testCase.time)
		if err != nil {
			t.Errorf(err.Error())
		}
		v := DayStart(tv).Format("2006-01-02 15:04:05")
		assert.Equal(t, testCase.expected, v, testCase.tag)
	}
}

func TestDayEnd(t *testing.T) {
	testCases := []struct {
		tag      string
		time     string
		layout   string
		expected string
	}{
		{"t1", "2022-01-01 12:12:00", "2006-01-02 15:04:05", "2022-01-01 23:59:59"},
		{"t2", "2022-01-01 00:00:00", "2006-01-02 15:04:05", "2022-01-01 23:59:59"},
	}
	for _, testCase := range testCases {
		tv, _ := time.Parse(testCase.layout, testCase.time)
		v := DayEnd(tv).Format("2006-01-02 15:04:05")
		assert.Equal(t, testCase.expected, v, testCase.tag)
	}
}

func TestMonthStart(t *testing.T) {
	testCases := []struct {
		tag      string
		time     string
		layout   string
		expected string
	}{
		{"t1", "2022-01-12 12:12:00", "2006-01-02 15:04:05", "2022-01-01 00:00:00"},
		{"t2", "2022-01-21 00:00:00", "2006-01-02 15:04:05", "2022-01-01 00:00:00"},
	}
	for _, testCase := range testCases {
		tv, _ := time.Parse(testCase.layout, testCase.time)
		v := MonthStart(tv).Format("2006-01-02 15:04:05")
		assert.Equal(t, testCase.expected, v, testCase.tag)
	}
}

func TestMonthEnd(t *testing.T) {
	testCases := []struct {
		tag      string
		time     string
		layout   string
		expected string
	}{
		{"t1", "2022-01-12 12:12:00", "2006-01-02 15:04:05", "2022-01-31 23:59:59"},
		{"t2", "2022-02-21 00:00:00", "2006-01-02 15:04:05", "2022-02-28 23:59:59"},
	}
	for _, testCase := range testCases {
		tv, _ := time.Parse(testCase.layout, testCase.time)
		v := MonthEnd(tv).Format("2006-01-02 15:04:05")
		assert.Equal(t, testCase.expected, v, testCase.tag)
	}
}

func TestWeekStart(t *testing.T) {
	testCases := []struct {
		tag      string
		yearWeek int
		expected string
	}{
		{"t1", 202201, "2022-01-03 00:00:00"},
		{"t2", 202202, "2022-01-10 00:00:00"},
	}
	for _, testCase := range testCases {
		v := WeekStart(testCase.yearWeek).Format("2006-01-02 15:04:05")
		assert.Equal(t, testCase.expected, v, testCase.tag)
	}
}

func TestWeekEnd(t *testing.T) {
	testCases := []struct {
		tag      string
		yearWeek int
		expected string
	}{
		{"t1", 202201, "2022-01-09 23:59:59"},
		{"t2", 202202, "2022-01-16 23:59:59"},
	}
	for _, testCase := range testCases {
		v := WeekEnd(testCase.yearWeek).Format("2006-01-02 15:04:05")
		assert.Equal(t, testCase.expected, v, testCase.tag)
	}
}

func TestYearWeeksByWeek(t *testing.T) {
	testCases := []struct {
		tag           string
		beginYearWeek int
		endYearWeek   int
		expected      []int
	}{
		{"t1", 202201, 202202, []int{202201, 202202}},
		{"t2", 202201, 202204, []int{202201, 202202, 202203, 202204}},
	}
	for _, testCase := range testCases {
		v := YearWeeksByWeek(testCase.beginYearWeek, testCase.endYearWeek)
		assert.Equal(t, testCase.expected, v, testCase.tag)
	}
}

func TestYearWeeksByTime(t *testing.T) {
	testCases := []struct {
		tag       string
		beginDate string
		endDate   string
		expected  []int
	}{
		{"t1", "2022-01-01", "2022-01-02", []int{202152}},
		{"t2", "2022-01-01", "2022-02-02", []int{202152, 202201, 202202, 202203, 202204, 202205}},
	}
	for _, testCase := range testCases {
		beginDate, _ := time.Parse("2006-01-02", testCase.beginDate)
		endDate, _ := time.Parse("2006-01-02", testCase.endDate)
		v := YearWeeksByTime(beginDate, endDate)
		assert.Equal(t, testCase.expected, v, testCase.tag)
	}
}
