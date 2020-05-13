package timerange

import "time"

func TimeZoneOffset(t time.Time) time.Duration {
	_, offset := FromTime(t).Zone()
	return time.Duration(-offset)*time.Second
}

func LocalTimeZoneOffset() time.Duration {
	return TimeZoneOffset(time.Now())
}
