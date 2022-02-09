package nullx

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

func TimeFrom(t time.Time) null.Time {
	if t.IsZero() {
		return NullTime()
	}
	return null.TimeFrom(t)
}

func NullTime() null.Time {
	return null.NewTime(time.Time{}, false)
}
