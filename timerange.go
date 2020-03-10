package timerange

import (
	"time"
)

type Time struct {
	time.Time
}

type Duration int64

type TimeDuration struct {
	time time.Time
	Duration
}

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

func (t Time) InMonth() TimeDuration {
	return TimeDuration{
		time:     t.Time,
		Duration: Duration(t.DaysInMonth()) * Day,
	}
}

func (t Time) InYear() TimeDuration {
	t1 := time.Date(t.Year(), 13, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	duration := t1.Sub(t2)
	return TimeDuration{
		time:     t.Time,
		Duration: Duration(duration),
	}
}

func (t Time) InQuarter() TimeDuration {
	offset := (int(t.Month()) - 1) % 3
	t1 := t.AddDate(0, 3-offset,0)
	t2 := t.AddDate(0, -offset,0)
	t1 = time.Date(t1.Year(), t1.Month(), 1, 0, 0, 0, 0, time.UTC)
	t2 = time.Date(t2.Year(), t2.Month(), 1, 0, 0, 0, 0, time.UTC)
	duration := t1.Sub(t2)
	return TimeDuration{
		time:     t.Time,
		Duration: Duration(duration),
	}
}

func (d TimeDuration) ToDays() int {
	return int(d.Duration / Day)
}

func (d TimeDuration) ToSeconds() int {
	return int(d.Duration / Second)
}
