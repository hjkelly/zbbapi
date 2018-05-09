package common

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

/*
This code was taken from / inspired by Google Cloud's `civil` project and modified so it would unmarshal valid dates properly, rather than failing.
https://github.com/GoogleCloudPlatform/google-cloud-go/blob/master/civil/civil.go
*/

// A Date represents a date (year, month, day).
// This type does not include location information, and therefore does not
// describe a unique 24-hour timespan.
type Date struct {
	Year  int        // Year (e.g., 2014).
	Month time.Month // Month of the year (January = 1, ...).
	Day   int        // Day of the month, starting at 1.
}

// DateOf returns the Date in which a time occurs in that time's location.
func DateOf(t time.Time) Date {
	var d Date
	d.Year, d.Month, d.Day = t.Date()
	return d
}

// DatePattern is the RFC3339 full-date format and should be used for all start dates in schedules.
var DatePattern = regexp.MustCompile("^([0-9]{4})-([0-9]{2})-([0-9]{2})$")

// ParseDate parses a valid-seeming RFC339 date string without regard for whether or not the numbers make sense.
func ParseDate(s string) (Date, error) {
	match := DatePattern.FindStringSubmatch(s)
	if len(match) != 4 {
		return Date{}, fmt.Errorf("Dates must be formatted: YYYY-MM-DD")
	}
	// Normally I'm not for throwing away errors, but we already validated that these are integers above.
	year, _ := strconv.Atoi(match[1])
	month, _ := strconv.Atoi(match[2])
	day, _ := strconv.Atoi(match[3])
	return Date{Year: year, Month: time.Month(month), Day: day}, nil
}

// String returns the date in RFC3339 full-date format.
func (d Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

// IsValid reports whether the date is valid.
func (d Date) IsValid() bool {
	return DateOf(d.In(time.UTC)) == d
}

// IsZero checks to see if the date has zero values.
func (d Date) IsZero() bool {
	return d == Date{}
}

func (d Date) ValidateNonZero() error {
	if d.IsZero() {
		return NewValidationError("", MissingCode, "You must provide a start date.")
	} else if !d.IsValid() {
		return NewValidationError("", BadDateCode, "Start date must be in the format YYYY-MM-DD, and it must be a valid date.")
	}
	return nil
}

// In returns the time corresponding to time 00:00:00 of the date in the location.
//
// In is always consistent with time.Date, even when time.Date returns a time
// on a different day. For example, if loc is America/Indiana/Vincennes, then both
//     time.Date(1955, time.May, 1, 0, 0, 0, 0, loc)
// and
//     civil.Date{Year: 1955, Month: time.May, Day: 1}.In(loc)
// return 23:00:00 on April 30, 1955.
//
// In panics if loc is nil.
func (d Date) In(loc *time.Location) time.Time {
	return time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, loc)
}

// AddDays returns the date that is n days in the future.
// n can also be negative to go into the past.
func (d Date) AddDays(n int) Date {
	return DateOf(d.In(time.UTC).AddDate(0, 0, n))
}

// DaysSince returns the signed number of days between the date and s, not including the end day.
// This is the inverse operation to AddDays.
func (d Date) DaysSince(s Date) (days int) {
	// We convert to Unix time so we do not have to worry about leap seconds:
	// Unix time increases by exactly 86400 seconds per day.
	deltaUnix := d.In(time.UTC).Unix() - s.In(time.UTC).Unix()
	return int(deltaUnix / 86400)
}

// Before reports whether d occurs before d2.
func (d Date) Before(d2 Date) bool {
	if d.Year != d2.Year {
		return d.Year < d2.Year
	}
	if d.Month != d2.Month {
		return d.Month < d2.Month
	}
	return d.Day < d2.Day
}

// After reports whether d occurs after d2.
func (d Date) After(d2 Date) bool {
	return d2.Before(d)
}

// MarshalText implements the encoding.TextMarshaler interface.
// The output is the result of d.String().
func (d Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The date is expected to be a string in a format accepted by ParseDate.
func (d *Date) UnmarshalText(data []byte) error {
	var err error
	*d, err = ParseDate(string(data))
	return err
}
