isoweek
=======

[![Go Reference](https://pkg.go.dev/badge/github.com/snabb/isoweek.svg)](https://pkg.go.dev/github.com/snabb/isoweek)
[![Build Status](https://github.com/snabb/isoweek/actions/workflows/go.yml/badge.svg)](https://github.com/snabb/isoweek/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/snabb/isoweek/branch/master/graph/badge.svg)](https://codecov.io/gh/snabb/isoweek)
[![Go Report Card](https://goreportcard.com/badge/github.com/snabb/isoweek)](https://goreportcard.com/report/github.com/snabb/isoweek)

The Go package isoweek calculates a starting date and time of [ISO 8601] week.

ISO 8601 standard defines the common [week number] system used in Europe
and many other countries. Monday is the first day of a week.

The Go standard library `time` package has `ISOWeek` function
for getting ISO 8601 week number of a given `time.Time`, but there is no
reverse functionality for getting a date from a week number. This package
implements that.

Invalid input is silently accepted. There is a separate `Validate`
function if week number validation is needed.

There are also functions for working with [Julian day numbers]. Using Julian
day numbers is often the easiest and fastest way to do date calculations.

This package does not work with the "traditional" week system used in
US/Canada/Japan/etc. (weeks starting on Sundays). However the Julian day
number functions may be still useful.

[ISO 8601]: https://en.wikipedia.org/wiki/ISO_8601
[week number]: https://en.wikipedia.org/wiki/ISO_week_date
[Julian day numbers]: https://en.wikipedia.org/wiki/Julian_day


Documentation
-------------

https://pkg.go.dev/github.com/snabb/isoweek


Examples
--------

### Using weeks with Go standard time.Time

A simple example which gets the starting time of the 1st week of 1985:
```Go
	t := isoweek.StartTime(1985, 1, time.UTC)
	fmt.Println(t)
	// Output: 1984-12-31 00:00:00 +0000 UTC
```
The returned time may be within a previous year as can be seen above.

The `AddDate` function in Go standard library `time` package can be used
for getting the `Time` at the end of the week or for iterating through weeks:
```Go
	t = t.AddDate(0, 0, 7)
	fmt.Println(t)
	// Output: 1985-01-07 00:00:00 +0000 UTC

	wyear, week := t.ISOWeek()
	fmt.Println(wyear, week)
	// Output: 1985 2
```

### Using weeks, dates and Julian day numbers

The same as above without using Go standard library `time` package:

```Go
	y, m, d := isoweek.StartDate(1985, 1)
	fmt.Println(y, m, d)
	// Output: 1984 December 31

	jdn := isoweek.DateToJulian(y, m, d)
	jdn = jdn + 7 // next week
	y, m, d = isoweek.JulianToDate(jdn)
	fmt.Println(y, m, d)
	// Output: 1985 January 7

	wyear, week := isoweek.FromDate(y, m, d)
	fmt.Println(wyear, week)
	// Output: 1985 2
```


Repository
----------

https://github.com/snabb/isoweek


License
-------

MIT
