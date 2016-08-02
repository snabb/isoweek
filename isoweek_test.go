package isoweek_test

import (
	"fmt"
	"github.com/snabb/isoweek"
	"testing"
	"time"
)

func TestStartTime(t *testing.T) {
	for y := 1; y < 4000; y++ {
		for w := 1; w < 53; w++ {
			st := isoweek.StartTime(y, w, time.UTC)
			ty, tw := st.ISOWeek()
			if ty != y || tw != w {
				t.Errorf("mismatch: "+
					"y = %d, w = %d "+
					"-> %s "+
					"-> ty = %d, tw = %d",
					y, w,
					st.Format("2006-01-02"),
					ty, tw)
			}
		}
	}
}

func ExampleStartTime() {
	st := isoweek.StartTime(1985, 1, time.UTC)
	fmt.Println(st)
	// Output:
	// 1984-12-31 00:00:00 +0000 UTC
}

func ExampleStartDate() {
	y, m, d := isoweek.StartDate(2000, 1)
	fmt.Println(d, m, y)
	// Output:
	// 3 January 2000
}
