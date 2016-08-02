// Ṕackage isoweek calculates the starting time of ISO 8601 week.
//
// The Go standard library "time" package has ISOWeek() method for getting
// ISO week number of given time.Time, but there is no reverse functionality
// for getting the date from week number. This package implements that.
package isoweek

import "time"

// isoWeekday returns the ISO weekday number of given day.
// (1 = Mon, 2 = Tue,.. 7 = Sun)
func isoWeekday(year, month, day int) (weekday int) {
	weekday = int(
		time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).Weekday())

	if weekday == 0 {
		weekday = 7
	}
	return weekday
}

// startOrdinal returns the ordinal day number (counting from 0) of Monday
// of the given week. Return value may be negative or bigger than number of
// days in a year.
func startOrdinal(year, week int) (ordinal int) {
	return week*7 - isoWeekday(year, 1, 4) - 3
}

// StartTime returns the starting time (Monday 00:00) of the given ISO week.
func StartTime(wyear, week int, loc *time.Location) (start time.Time) {
	return time.Date(wyear, 1, 1, 0, 0, 0, 0, loc).AddDate(0, 0,
		startOrdinal(wyear, week))
}

// StartDate returns the starting date (Monday) of the given ISO week.
func StartDate(wyear, week int) (year int, month time.Month, day int) {
	return StartTime(wyear, week, time.UTC).Date()
}
