package timex

import "time"

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
		days := int(t1.Weekday())
		if (month == 3 && day >= t1.AddDate(0, 0, 14-days).Day()) ||
			(month == 11 && day < t1.AddDate(0, 0, 7-days).Day()) {
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
