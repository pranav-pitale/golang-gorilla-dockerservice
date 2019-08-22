package utils

import (

	"time"
)

// DateTimeRanges is type for storing min, max and range of datetime between +/-3 years
type DateTimeRanges struct {
	MinDate   time.Time
	MaxDate   time.Time
	YearRange []int
}

// CreateYearRangeBetweenTimeStamp generates year ran between +/- 3years from timestamp and also has min max of it
func CreateYearRangeBetweenTimeStamp(timestamp string) *DateTimeRanges {

	t, _ := time.Parse(time.RFC3339, timestamp)
	mindate := t.AddDate(-3, 0, 0)
	maxdate := t.AddDate(3, 0, 0)
	yearrange := make([]int, 0)
	datetimerange := new(DateTimeRanges)
	datetimerange.MinDate = mindate
	datetimerange.MaxDate = maxdate

	// Generating year range between +/- 3 years
	for i := -3; i <= 3; i++ {

		yearrange = append(yearrange, int(t.AddDate(i, 0, 0).Year()))
	}
	datetimerange.YearRange = yearrange
	return datetimerange
}

// InTimeSpan checks whether given date time is between start and end time
func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
