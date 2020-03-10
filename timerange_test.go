package timerange

import (
	"testing"
	"time"
)

type testTime struct {
	time          time.Time
	daysInMounth  int
	daysInYear    int
	daysInQuarter int
}

var samples = []testTime{
	{
		time:          time.Date(2020, 2, 3, 2, 1, 10, 9, time.UTC),
		daysInMounth:  29,
		daysInYear:    366,
		daysInQuarter: 91,
	},
	{
		time:          time.Date(2019, 3, 3, 2, 1, 10, 9, time.UTC),
		daysInMounth:  31,
		daysInYear:    365,
		daysInQuarter: 90,
	},
	{
		time:          time.Date(2019, 12, 3, 2, 1, 10, 9, time.UTC),
		daysInMounth:  31,
		daysInYear:    365,
		daysInQuarter: 92,
	},
	{
		time:          time.Date(2018, 11, 3, 2, 1, 10, 9, time.UTC),
		daysInMounth:  30,
		daysInYear:    365,
		daysInQuarter: 92,
	},
}

func TestDurations_InMonth(t *testing.T) {
	for _, sample := range samples {
		daysInMonth := FromTime(sample.time).InMonth().ToDays()
		if daysInMonth != FromTime(sample.time).DaysInMonth() {
			t.Error("Invalid ToDays conversion")
		}
		if daysInMonth != sample.daysInMounth {
			t.Error("Invalid days in month count")
		}

		secondsInMonth := FromTime(sample.time).InMonth().ToSeconds()
		if secondsInMonth != sample.daysInMounth*24*60*60 {
			t.Error("Invalid seconds in month count")
		}
	}
}

func TestDurations_InYear(t *testing.T) {
	for _, sample := range samples {
		daysInYear := FromTime(sample.time).InYear().ToDays()
		if daysInYear != sample.daysInYear {
			t.Error("Invalid days in year count")
		}

		secondsInYear := FromTime(sample.time).InYear().ToSeconds()
		if secondsInYear != sample.daysInYear*24*60*60 {
			t.Error("Invalid seconds in year count")
		}
	}
}

func TestDurations_InQuarter(t *testing.T) {
	for _, sample := range samples {
		daysInQuarter := FromTime(sample.time).InQuarter().ToDays()
		if daysInQuarter != sample.daysInQuarter {
			t.Error("Invalid days in quarter count", sample.time, FromTime(sample.time).InQuarter().ToDays())
		}

		secondsInQuarter := FromTime(sample.time).InQuarter().ToSeconds()
		if secondsInQuarter != sample.daysInQuarter*24*60*60 {
			t.Error("Invalid seconds in quarter count")
		}
	}
}
