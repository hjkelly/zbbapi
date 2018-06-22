package models

import (
	"testing"

	"github.com/hjkelly/zbbapi/common"
	"github.com/stretchr/testify/assert"
)

var tuesdayTheNineteenth = common.Date{Year: 2018, Month: 6, Day: 19}
var lastWednesday = common.Date{Year: 2018, Month: 6, Day: 13}
var aYearAgo = common.Date{tuesdayTheNineteenth.Year - 1, tuesdayTheNineteenth.Month, tuesdayTheNineteenth.Day}
var aDayOfMonth = 15
var anotherDayOfMonth = 30
var monday = "Monday"

func TestOccurrences(t *testing.T) {
	for idx, tc := range []struct {
		desc   string
		sched  Schedule
		start  common.Date
		end    common.Date
		expect int
	}{
		// year ----------
		{
			desc:   "annual bill miss",
			sched:  Schedule{Year: &aYearAgo},
			start:  tuesdayTheNineteenth.AddDays(5),
			end:    tuesdayTheNineteenth.AddDays(10),
			expect: 0,
		},
		{
			desc:   "annual bill hit",
			sched:  Schedule{Year: &aYearAgo},
			start:  tuesdayTheNineteenth,
			end:    tuesdayTheNineteenth.AddDays(1),
			expect: 1,
		},
		{
			desc:   "annual bill multiple (silly, I know)",
			sched:  Schedule{Year: &aYearAgo},
			start:  tuesdayTheNineteenth,
			end:    tuesdayTheNineteenth.AddYears(1),
			expect: 2,
		},
		// month ----------
		{
			desc:   "monthly bill miss",
			sched:  Schedule{Month: &aDayOfMonth},
			start:  tuesdayTheNineteenth,
			end:    tuesdayTheNineteenth.AddDays(14),
			expect: 0,
		},
		{
			desc:   "monthly bill hit",
			sched:  Schedule{Month: &anotherDayOfMonth},
			start:  tuesdayTheNineteenth,
			end:    tuesdayTheNineteenth.AddDays(14),
			expect: 1,
		},
		{
			desc:   "monthly bill multiple",
			sched:  Schedule{Month: &aDayOfMonth},
			start:  tuesdayTheNineteenth,
			end:    tuesdayTheNineteenth.AddDays(60),
			expect: 2,
		},
		// halfmonth ----------
		{
			desc:   "semi-monthly bill miss",
			sched:  Schedule{HalfMonth: &[]int{aDayOfMonth, anotherDayOfMonth}},
			start:  tuesdayTheNineteenth,
			end:    tuesdayTheNineteenth.AddDays(5),
			expect: 0,
		},
		{
			desc:   "semi-monthly bill hit",
			sched:  Schedule{HalfMonth: &[]int{aDayOfMonth, anotherDayOfMonth}},
			start:  tuesdayTheNineteenth,
			end:    tuesdayTheNineteenth.AddDays(14),
			expect: 1,
		},
		{
			desc:   "semi-monthly bill multiple",
			sched:  Schedule{HalfMonth: &[]int{aDayOfMonth, anotherDayOfMonth}},
			start:  tuesdayTheNineteenth,
			end:    tuesdayTheNineteenth.AddDays(30),
			expect: 2,
		},
		// twoweeks ----------
		{
			desc:   "biweekly bill miss",
			sched:  Schedule{TwoWeeks: &lastWednesday},
			start:  tuesdayTheNineteenth,
			end:    tuesdayTheNineteenth.AddDays(5),
			expect: 0,
		},
		{
			desc:   "biweekly bill hit",
			sched:  Schedule{TwoWeeks: &lastWednesday},
			start:  tuesdayTheNineteenth,
			end:    tuesdayTheNineteenth.AddDays(13),
			expect: 1,
		},
		{
			desc:   "biweekly bill multiple",
			sched:  Schedule{TwoWeeks: &lastWednesday},
			start:  tuesdayTheNineteenth,
			end:    tuesdayTheNineteenth.AddDays(29),
			expect: 2,
		},
		// weekly ----------
		{
			desc:   "weekly bill miss",
			sched:  Schedule{Week: &monday},
			start:  tuesdayTheNineteenth,
			end:    tuesdayTheNineteenth.AddDays(5),
			expect: 0,
		},
		{
			desc:   "weekly bill hit",
			sched:  Schedule{Week: &monday},
			start:  tuesdayTheNineteenth,
			end:    tuesdayTheNineteenth.AddDays(6),
			expect: 1,
		},
		{
			desc:   "weekly bill multiple",
			sched:  Schedule{Week: &monday},
			start:  tuesdayTheNineteenth,
			end:    tuesdayTheNineteenth.AddDays(13),
			expect: 2,
		},
	} {
		actual := tc.sched.Occurrences(tc.start, tc.end)
		assert.Equal(t, tc.expect, actual,
			"expected %d, actual %d occurrences for case %d named: %s",
			tc.expect, actual, idx, tc.desc)
	}
}
