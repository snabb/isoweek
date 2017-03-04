package isoweek_test

import (
	"fmt"
	"github.com/snabb/isoweek"
	"testing"
	"time"
)

func TestJulianToDate(test *testing.T) {
	j := isoweek.DateToJulian(1, time.January, 1)

	for {
		y, m, d := isoweek.JulianToDate(j)
		if y >= 4000 {
			break
		}
		if j != isoweek.DateToJulian(y, m, d) {
			test.Errorf("mismatch on %04d-%02d-%02d", y, m, d)
		}
		j++
	}
}

func TestDateToJulian(test *testing.T) {
	t := time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)

	for t.Year() < 4000 {
		j := isoweek.DateToJulian(t.Date())

		y, m, d := isoweek.JulianToDate(j)

		if y != t.Year() || m != t.Month() || d != t.Day() {
			test.Errorf("mismatch on %s", t.Format("2006-01-02"))
		}
		if j+1 != isoweek.DateToJulian(y, m, d+1) {
			test.Errorf("mismatch 2 on %s", t.Format("2006-01-02"))
		}
		t = t.AddDate(0, 0, 1)
	}
}

func ExampleDateToJulian() {
	fmt.Println(isoweek.DateToJulian(2006, 1, 2))
	// Output: 2453738
}

func ExampleJulianToDate() {
	fmt.Println(isoweek.JulianToDate(2453738))
	// Output: 2006 January 2
}
