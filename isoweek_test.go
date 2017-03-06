package isoweek_test

import (
	"fmt"
	"github.com/snabb/isoweek"
	"testing"
	"time"
)

// TestISOWeekday tests all days from year 1 until year 4000.
// Ensures that behaviour matches the Go standard library Weekday().
func TestISOWeekday(test *testing.T) {
	t := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	var wd1, wd2 int
	for t.Year() < 4000 {
		wd1 = int(t.Weekday())
		wd2 = isoweek.ISOWeekday(t.Date())

		if wd2 == 7 {
			wd2 = 0
		}
		if wd1 != wd2 {
			test.Errorf("mismatch on %s", t.Format("2006-01-02"))
		}
		t = t.AddDate(0, 0, 1)
	}
}

func ExampleISOWeekday() {
	fmt.Println(isoweek.ISOWeekday(1984, 1, 1))
	// Output: 7
}

// TestStartTime tests all weeks from year 1 until year 4000.
// Ensures that behaviour matches the Go standard library ISOWeek().
func TestStartTime(test *testing.T) {
	t := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	for t.Weekday() != time.Monday {
		t = t.AddDate(0, 0, 1)
	}
	for t.Year() < 4000 {
		wy, ww := t.ISOWeek()
		wst := isoweek.StartTime(wy, ww, time.UTC)
		if !wst.Equal(t) {
			test.Errorf("mismatch: %v != %v (wy = %d, ww = %d)",
				t, wst, wy, ww)
		}
		t = t.AddDate(0, 0, 7)
	}
}

func ExampleStartTime() {
	st := isoweek.StartTime(1985, 1, time.UTC)
	fmt.Println(st)
	// Output: 1984-12-31 00:00:00 +0000 UTC
}

func ExampleStartDate() {
	y, m, d := isoweek.StartDate(2000, 1)
	fmt.Println(d, m, y)
	// Output: 3 January 2000
}

// TestFromDate  tests all days from year 1 until year 4000.
// Ensures that behaviour matches the Go standard library ISOWeek().
func TestFromDate(test *testing.T) {
	t := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	for t.Year() < 4000 {
		wy, ww := t.ISOWeek()
		wy2, ww2 := isoweek.FromDate(t.Date())
		if wy != wy2 || ww != ww2 {
			test.Errorf("mismatch on %s", t.Format("2006-01-02"))
		}
		t = t.AddDate(0, 0, 1)
	}
}

func ExampleFromDate() {
	fmt.Println(isoweek.FromDate(1984, 1, 1))
	// Output: 1983 52
}

func ExampleValidate() {
	fmt.Println(isoweek.Validate(2016, 52),
		isoweek.Validate(2016, 53),
		isoweek.Validate(2017, 0))
	// Output: true false false
}
