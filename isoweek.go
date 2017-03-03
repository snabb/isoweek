// Ṕackage isoweek calculates the starting date and time of ISO 8601 week.
//
// The Go standard library "time" package has ISOWeek() method for getting
// ISO week number of given time.Time, but there is no reverse functionality
// for getting the date from week number. This package implements that.
//
// Invalid week numbers are silently accepted. There is a separate Validate()
// function if validation is needed.
package isoweek

import "time"

// Tomohiko Sakamoto, modified for weeks starting on Monday
var tomohiko = [...]int{6, 2, 1, 4, 6, 2, 4, 0, 3, 5, 1, 3}

// ISOWeekday returns the ISO weekday number of given day.
// (1 = Mon, 2 = Tue,.. 7 = Sun)
//
// This is different from Go's standard time.Weekday, which you should use
// normally. It is exposed because it may be useful for some calculations.
func ISOWeekday(year int, month time.Month, day int) (weekday int) {
	if month < 1 || month > 12 {
		// prevent panic with array lookup
		return 0
	}
	// Tomohiko Sakamoto
	if month < 3 {
		year--
	}
	return (year+year/4-year/100+year/400+tomohiko[month-1]+day)%7 + 1
}

// startOffset returns the offset (in days) from the start of a year to
// Monday of the given week. Offset may be negative.
func startOffset(year, week int) (offset int) {
	return week*7 - ISOWeekday(year, 1, 4) - 3
}

// StartTime returns the starting time (Monday 00:00) of the given ISO week.
func StartTime(wyear, week int, loc *time.Location) (start time.Time) {
	return time.Date(wyear, 1, 1, 0, 0, 0, 0, loc).
		AddDate(0, 0, startOffset(wyear, week))
}

// StartDate returns the starting date (Monday) of the given ISO week.
func StartDate(wyear, week int) (year int, month time.Month, day int) {
	return StartTime(wyear, week, time.UTC).Date()
}

// Validate checks if a week number is valid. Returns true if it is valid.
func Validate(wyear, week int) (ok bool) {
	if week < 1 || week > 53 {
		return false
	}
	wyear2, week2 := StartTime(wyear, week, time.UTC).ISOWeek()

	if wyear == wyear2 && week == week2 {
		return true
	}
	return false
}
