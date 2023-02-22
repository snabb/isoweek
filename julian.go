package isoweek

import "time"

// DateToJulian converts a date to a Julian day number.
func DateToJulian(year int, month time.Month, day int) (jdn int) {
	// Claus Tøndering's Calendar FAQ
	// http://www.tondering.dk/claus/cal/julperiod.php#formula

	if month < 3 {
		year = year - 1
		month = month + 12
	}
	year = year + 4800

	return day + (153*(int(month)-3)+2)/5 + 365*year +
		year/4 - year/100 + year/400 - 32045
}

// JulianToDate converts a Julian day number to a date.
func JulianToDate(jdn int) (year int, month time.Month, day int) {
	// Richards, E. G. (2013) pp. 585–624

	e := 4*(jdn+1401+(4*jdn+274277)/146097*3/4-38) + 3
	h := e%1461/4*5 + 2

	day = h%153/5 + 1
	month = time.Month((h/153+2)%12 + 1)
	year = e/1461 - 4716 + (14-int(month))/12

	return year, month, day
}
