isoweek
=======

The Go package isoweek calculates the starting time and date of ISO 8601
week.

The Go standard library "time" package has ISOWeek() method for getting
ISO week number of given time.Time, but there is no reverse functionality
for getting the date from week number. This package implements that.

Documentation:

https://godoc.org/github.com/snabb/isoweek

Simple example:
```Go
	st := isoweek.StartTime(1985, 1, time.UTC)
	fmt.Println(st)
	// Output: 1984-12-31 00:00:00 +0000 UTC
```

If you want to iterate through weeks or to get the ending time of the week
you can use the standard AddDate() method:
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
