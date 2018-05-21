package models

import (
	"testing"

	"github.com/hjkelly/zbbapi/common"
	"github.com/stretchr/testify/assert"
)

var tooLowDay = 0
var minDay = 1
var maxDay = 31
var tooHighDay = 32

var noDays = []int{}
var notEnoughDays = []int{minDay}
var lowDays = []int{tooLowDay, minDay}
var highDays = []int{maxDay, tooHighDay}
var goodDays = []int{minDay, maxDay}

var dayOfWeek = "Tuesday"
var badDayOfWeek = "asdf"

func TestScheduleGetValidated(t *testing.T) {
	for _, testCase := range []struct {
		desc   string
		input  Schedule
		result Schedule
		err    error
	}{
		// valid: ----------
		{
			desc:   "yearly",
			input:  Schedule{Year: &common.Date{Year: 2018, Month: 5, Day: 19}},
			result: Schedule{Year: &common.Date{Year: 2018, Month: 5, Day: 19}},
			err:    nil,
		},
		{
			desc:   "monthly (min)",
			input:  Schedule{Month: &minDay},
			result: Schedule{Month: &minDay},
			err:    nil,
		},
		{
			desc:   "monthly (max)",
			input:  Schedule{Month: &maxDay},
			result: Schedule{Month: &maxDay},
			err:    nil,
		},
		{
			desc:   "semimonthly",
			input:  Schedule{HalfMonth: &goodDays},
			result: Schedule{HalfMonth: &goodDays},
			err:    nil,
		},
		{
			desc:   "biweekly",
			input:  Schedule{TwoWeeks: &common.Date{Year: 2018, Month: 5, Day: 19}},
			result: Schedule{TwoWeeks: &common.Date{Year: 2018, Month: 5, Day: 19}},
			err:    nil,
		},
		{
			desc:   "weekly",
			input:  Schedule{Week: &dayOfWeek},
			result: Schedule{Week: &dayOfWeek},
			err:    nil,
		},
		// invalid: values ----------
		{
			desc:   "missing yearly",
			input:  Schedule{Year: &common.Date{}},
			result: Schedule{},
			err:    common.NewValidationError("yearStarting", common.MissingCode, "You must provide a date."),
		},
		{
			desc:   "invalid yearly",
			input:  Schedule{Year: &common.Date{Year: 2018, Month: 13, Day: 4}},
			result: Schedule{},
			err:    common.NewValidationError("yearStarting", common.BadDateCode, "Date must be in the format YYYY-MM-DD, and it must be a valid date."),
		},
		{
			desc:   "invalid monthly (too low)",
			input:  Schedule{Month: &tooLowDay},
			result: Schedule{},
			err:    common.NewValidationError("monthOnDay", common.NumOutOfRangeCode, "Must be between 1 and 31 (inclusive)."),
		},
		{
			desc:   "invalid monthly (too high)",
			input:  Schedule{Month: &tooHighDay},
			result: Schedule{},
			err:    common.NewValidationError("monthOnDay", common.NumOutOfRangeCode, "Must be between 1 and 31 (inclusive)."),
		},
		{
			desc:   "invalid semimonthly (no days)",
			input:  Schedule{HalfMonth: &noDays},
			result: Schedule{},
			err:    common.NewValidationError("halfMonthOnDays", common.MissingCode, "You must provide at least two days of the month."),
		},
		{
			desc:   "invalid semimonthly (not enough days)",
			input:  Schedule{HalfMonth: &notEnoughDays},
			result: Schedule{},
			err:    common.NewValidationError("halfMonthOnDays", common.MissingCode, "You must provide at least two days of the month."),
		},
		{
			desc:   "invalid semimonthly (days too low)",
			input:  Schedule{HalfMonth: &lowDays},
			result: Schedule{},
			err:    common.NewValidationError("halfMonthOnDays.0", common.NumOutOfRangeCode, "Must be between 1 and 31 (inclusive)."),
		},
		{
			desc:   "invalid semimonthly (days too high)",
			input:  Schedule{HalfMonth: &highDays},
			result: Schedule{},
			err:    common.NewValidationError("halfMonthOnDays.1", common.NumOutOfRangeCode, "Must be between 1 and 31 (inclusive)."),
		},
		{
			desc:   "missing biweekly",
			input:  Schedule{TwoWeeks: &common.Date{}},
			result: Schedule{},
			err:    common.NewValidationError("twoWeeksStarting", common.MissingCode, "You must provide a date."),
		},
		{
			desc:   "invalid biweekly",
			input:  Schedule{TwoWeeks: &common.Date{Year: 2018, Month: 13, Day: 4}},
			result: Schedule{},
			err:    common.NewValidationError("twoWeeksStarting", common.BadDateCode, "Date must be in the format YYYY-MM-DD, and it must be a valid date."),
		},
		{
			desc:   "invalid weekly",
			input:  Schedule{Week: &badDayOfWeek},
			result: Schedule{},
			err:    common.NewValidationError("weekOn", common.BadEnumChoiceCode, "You must specify the day of the week: Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday"),
		},
		{
			desc:   "no schedule defined",
			input:  Schedule{},
			result: Schedule{},
			err:    common.NewValidationError("", common.MissingCode, "You must specify exactly one schedule: yearStarting, monthOnDay, halfMonthOnDays, twoWeeksStarting, weekOn"),
		},
		{
			desc:   "too many schedules defined",
			input:  Schedule{Year: &common.Date{Year: 2018, Month: 5, Day: 19}, Month: &minDay},
			result: Schedule{},
			err:    common.NewValidationError("", common.TooManyCode, "You must specify exactly one schedule: yearStarting, monthOnDay, halfMonthOnDays, twoWeeksStarting, weekOn"),
		},
		// overlapping specs ----------
		// ""
	} {
		result, err := testCase.input.GetValidated()
		assert.Equal(t, testCase.result, result, "CASE %s, didn't get expected result", testCase.desc)
		assert.Equal(t, testCase.err, err, "CASE %s, didn't get expected error", testCase.desc)
	}
}
