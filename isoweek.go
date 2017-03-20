// Package isoweek calculates the starting date and time of ISO 8601 week.
//
// The Go standard library "time" package has ISOWeek() method for getting
// ISO week number of given time.Time, but there is no reverse functionality
// for getting the date from week number. This package implements that.
//
// Invalid input is silently accepted. There is a separate Validate()
// function if week number validation is needed.
//
// There are also functions for working with Julian day numbers. Using Julian
// day numbers is often the easiest and fastest way to do date calculations.
//
package isoweek

import "time"

// ISOWeekday returns the ISO weekday number of given day.
// (1 = Mon, 2 = Tue,.. 7 = Sun)
//
// This is different from Go's standard time.Weekday.
func ISOWeekday(year int, month time.Month, day int) (weekday int) {
	// Richards, E. G. (2013) pp. 592, 618

	return DateToJulian(year, month, day)%7 + 1
}

// startOffset returns the offset (in days) from the start of a year to
// Monday of the given week. Offset may be negative.
func startOffset(y, week int) (offset int) {
	// This is optimized version of the following:
	//
	// return week*7 - ISOWeekday(y, 1, 4) - 3
	//
	// Uses Tomohiko Sakamoto's algorithm for calculating the weekday.

	y = y - 1
	return week*7 - (y+y/4-y/100+y/400+3)%7 - 4
}

// StartTime returns the starting time (Monday 00:00) of the given ISO week.
func StartTime(wyear, week int, loc *time.Location) (start time.Time) {
	y, m, d := StartDate(wyear, week)
	return time.Date(y, m, d, 0, 0, 0, 0, loc)
}

// StartDate returns the starting date (Monday) of the given ISO week.
func StartDate(wyear, week int) (year int, month time.Month, day int) {
	return JulianToDate(
		DateToJulian(wyear, 1, 1) + startOffset(wyear, week))
}

// ordinalInYear returns the ordinal (within a year) day number.
func ordinalInYear(year int, month time.Month, day int) (dayNo int) {
	return DateToJulian(year, month, day) - DateToJulian(year, 1, 1) + 1
}

// FromDate returns ISO week number of a date.
func FromDate(year int, month time.Month, day int) (wyear, week int) {
	week = (ordinalInYear(year, month, day) - ISOWeekday(year, month, day) + 10) / 7
	if week < 1 {
		return FromDate(year-1, 12, 31) // last week of preceding year
	}
	if week == 53 &&
		DateToJulian(StartDate(year+1, 1)) <= DateToJulian(year, month, day) {
		return year + 1, 1 // first week of following year
	}
	return year, week
}

// Validate checks if a week number is valid. Returns true if it is valid.
func Validate(wyear, week int) (ok bool) {
	if week < 1 || week > 53 {
		return false
	}
	wyear2, week2 := FromDate(StartDate(wyear, week))

	if wyear == wyear2 && week == week2 {
		return true
	}
	return false
}
