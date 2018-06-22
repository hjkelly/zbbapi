package common

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDaysUntilWeekday(t *testing.T) {
	for idx, tc := range []struct {
		desc          string
		currentDate   Date
		targetWeekday time.Weekday
		expected      int
	}{
		{
			desc:          "Tuesday to Sunday",
			currentDate:   Date{Year: 2018, Month: 6, Day: 19}, // Tuesday
			targetWeekday: 0,
			expected:      5,
		},
		{
			desc:          "Tuesday to Saturday",
			currentDate:   Date{Year: 2018, Month: 6, Day: 19}, // Tuesday
			targetWeekday: 6,
			expected:      4,
		},
		{
			desc:          "Tuesday to Monday",
			currentDate:   Date{Year: 2018, Month: 6, Day: 19}, // Tuesday
			targetWeekday: 1,
			expected:      6,
		},
		{
			desc:          "Sunday to Sunday",
			currentDate:   Date{Year: 2018, Month: 6, Day: 24}, // Sunday
			targetWeekday: 0,
			expected:      0,
		},
		{
			desc:          "Sunday to Saturday",
			currentDate:   Date{Year: 2018, Month: 6, Day: 24}, // Sunday
			targetWeekday: 6,
			expected:      6,
		},
		{
			desc:          "Saturday to Sunday",
			currentDate:   Date{Year: 2018, Month: 6, Day: 23}, // Saturday
			targetWeekday: 0,
			expected:      1,
		},
	} {
		actual := tc.currentDate.DaysUntilWeekday(tc.targetWeekday)
		assert.Equal(t, actual, tc.expected,
			"In case %d (%s), didn't get expected days from %s to %s; got %d instead of %d)",
			idx, tc.desc,
			tc.currentDate.NoonUTC().Weekday().String(),
			tc.targetWeekday.String(),
			actual, tc.expected)
	}
}
