package timex

import (
	"time"
)

// IsAmericaSummerTime 是否为美国夏令时间
// 夏令时开始于每年3月的第二个周日凌晨，人们需要将时间调早 (顺时针) 1个小时；
// 夏令时结束于每年11月的第一个周日凌晨，人们需要将时间调晚 (逆时针) 1个小时。
func IsAmericaSummerTime(t time.Time) (yes bool) {
	if t.IsZero() {
		return
	}

	month := t.Month()
	switch month {
	case 4, 5, 6, 7, 8, 9, 10:
		yes = true
	case 3, 11:
		day := t.Day()
		t1 := t.AddDate(0, 0, -day+1)
		weekday := int(t1.Weekday())
		if (month == 3 && day >= t1.AddDate(0, 0, 14-weekday).Day()) ||
			(month == 11 && day < t1.AddDate(0, 0, 7-weekday).Day()) {
			yes = true
		}
	}
	return
}

// ChineseTimeLocation Return chinese time location
func ChineseTimeLocation() *time.Location {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		loc = time.FixedZone("CST", 8*3600)
	}
	return loc
}

func Between(t, begin, end time.Time) bool {
	return (t.After(begin) && t.Before(end)) || t.Equal(begin) || t.Equal(end)
}

func DayStart(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

func DayEnd(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, time.Local)
}

func MonthStart(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
}

func MonthEnd(t time.Time) time.Time {
	return DayEnd(MonthStart(t).AddDate(0, 1, -1))
}

// IsAM Check is AM
func IsAM(t time.Time) bool {
	return t.Hour() <= 11
}

// IsPM Check is PM
func IsPM(t time.Time) bool {
	return t.Hour() >= 12
}

func WeekStart(yearWeek int) time.Time {
	year := yearWeek / 100
	week := yearWeek % year
	// Start from the middle of the year:
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)

	// Roll back to Monday:
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	// Difference in weeks:
	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)

	return t
}

func WeekEnd(yearWeek int) time.Time {
	t := WeekStart(yearWeek).AddDate(0, 0, 6)
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.UTC)
}

func YearWeeksByWeek(startYearWeek, endYearWeek int) []int {
	weeks := make([]int, 0)
	weekStart := WeekStart(startYearWeek)
	weekEnd := WeekStart(endYearWeek)
	for {
		if weekStart.After(weekEnd) {
			break
		}
		y, w := weekStart.ISOWeek()
		weeks = append(weeks, y*100+w)
		weekStart = weekStart.AddDate(0, 0, 7)
	}
	return weeks
}

func YearWeeksByTime(startDate, endDate time.Time) []int {
	y1, w1 := startDate.ISOWeek()
	y2, w2 := endDate.ISOWeek()
	return YearWeeksByWeek(y1*100+w1, y2*100+w2)
}
