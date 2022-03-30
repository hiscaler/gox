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
