isoweek
=======

[![GoDoc](https://godoc.org/github.com/snabb/isoweek?status.svg)](https://godoc.org/github.com/snabb/isoweek)

The Go package isoweek calculates the starting date and time of ISO 8601
week.

The Go standard library "time" package has ISOWeek() method for getting
ISO week number of given time.Time, but there is no reverse functionality
for getting the date from week number. This package implements that.

Invalid week numbers are silently accepted. There is a separate Validate()
function if validation is needed.

It also has functions for working with Julian day numbers (which is
often the easiest way to do date calculations -- and much faster than
using Go's standard "time" package) as well as a function for getting
the international standard weekday number from a date. The Go standard
library uses strange US weekday numbering.


Documentation:

https://godoc.org/github.com/snabb/isoweek

Simple example which gets the starting time of the 1st week of 1985:
```Go
	st := isoweek.StartTime(1985, 1, time.UTC)
	fmt.Println(st)
	// Output: 1984-12-31 00:00:00 +0000 UTC
```
The returned time may be within previous year as can be seen above.

The AddDate() method in Go standard library can be used for getting the
time at the end of the week and for iterating through weeks:
```Go
	et := st.AddDate(0, 0, 7)
	wyear, week := et.ISOWeek()
	fmt.Println(wyear, week)
	// Output: 1985 2
```

The Git repository is located at: https://github.com/snabb/isoweek


License
-------

MIT
