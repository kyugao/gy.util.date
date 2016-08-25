package udate

import "time"

func DayStart(t time.Time, offset int, timezone *time.Location) (start time.Time) {
	if timezone == nil {
		timezone = time.UTC
	}
	year, month, day := t.Date()
	start = time.Date(year, month, day + offset, 0, 0, 0, 0, timezone)
	return
}

