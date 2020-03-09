package timerange

import (
	"testing"
	"time"
)

type testTime struct {
	time time.Time
	daysInMounth int
}

var samples = []testTime{
	{
		time: time.Date(2020, 2, 3, 2, 1, 10, 9, time.UTC),
		daysInMounth: 29,
	},
	{
		time: time.Date(2020, 3, 3, 2, 1, 10, 9, time.UTC),
		daysInMounth: 31,
	},
}

func TestDuration_ToDays(t *testing.T) {
	for _, sample := range samples{
		daysInMonth := FromTime(sample.time).InMonth().ToDays()
		if daysInMonth != FromTime(sample.time).DaysInMonth() {
			t.Error("Invalid ToDays conversion")
		}
		if daysInMonth != sample.daysInMounth {
			t.Error("Invalid days count")
		}
	}
}