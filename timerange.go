package timerange

import "time"

type Time struct {
	time.Time
}

type Duration int64

const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
	Day                  = 24 * Hour
)

func Now() Time {
	return Time{time.Now()}
}

func FromTime(t time.Time) Time {
	return Time{t}
}

func (t Time) DaysInMonth() int {
	nextMonth := time.Date(t.Year(), t.Month(), 32, 0, 0, 0, 0, time.UTC)
	return 32 - nextMonth.Day()
}

func (t Time) InMonth() Duration {
	return Duration(t.DaysInMonth()) * Day
}

func (d Duration) ToDays() int {
	return int(d / Day)
}
