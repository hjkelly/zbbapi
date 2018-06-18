package models

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hjkelly/zbbapi/common"
)

// Schedule defines when an event is expected to occur. Depending on the type, you'll have to provide at least one other piece of info so the event can be predicted.
type Schedule struct {
	Year      *common.Date `json:"yearStarting,omitempty"`
	Month     *int         `json:"monthOnDay,omitempty"`
	HalfMonth *[]int       `json:"halfMonthOnDays,omitempty"`
	TwoWeeks  *common.Date `json:"twoWeeksStarting,omitempty"`
	Week      *string      `json:"weekOn,omitempty"`
}

func (s Schedule) Occurrences(startDate, endDate common.Date) int {
	occurrences := 0
	if s.Year != nil {
		testDate := *s.Year
		// Advance the year until we pass our end date. If at any point we find a date that lands between our start and end dates, it's that time of year. It technically handles multiple occurrences over multiple years, just to see if any common logic stands out.
		for testDate.BeforeEqual(endDate) {
			if testDate.Between(startDate, endDate) {
				occurrences = occurrences + 1
			}
			testDate = testDate.AddYears(1)
		}
	} else if s.Month != nil {
		thisYear, thisMonth, _ := time.Now().Date()
		testDate := common.Date{
			Year:  thisYear,
			Month: thisMonth,
			Day:   *s.Month,
		}
		for testDate.BeforeEqual(endDate) {
			if testDate.Between(startDate, endDate) {
				occurrences = occurrences + 1
			}
			testDate = testDate.AddMonths(1)
		}
	} else if s.HalfMonth != nil {
		thisYear, _, _ := DateOf(time.Now())
		testDate := common.Date{
			Year:  thisYear,
			Month: startDate.Month,
			Day:   1,
		}
		monthOffset := 0
		for testDate.BeforeEqual(testDate) {
			for _, dayOfMonth := range *s.HalfMonth {
				testDate.Day = dayOfMonth
				if testDate.Between(startDate, endDate) {
					occurrences = occurrences + 1
				}
			}
			testDate = testDate.AddMonths(1)
		}
	} else if s.TwoWeeks != nil {
		testDate := *s.TwoWeeks
		// Advance the date by two weeks until we pass the end date, counting each occurrence.
		for testDate.BeforeEqual(endDate) {
			if testDate.Between(startDate, endDate) {
				occurrences = occurrences + 1
			}
			testDate = testDate.AddDays(14)
		}
	} else if s.Week != nil {
		testDate := startDate
		if testDate.NoonGMT().Weekday() == *s.Week {

		}
		// TODO
	} else {
		log.Printf("The schedule didn't have any day defined.")
	}
	return occurrences
}

// GetValidated returns a sanitized schedule if the type and other sometimes-required attributes are in order; otherwise, returns an error.
func (s Schedule) GetValidated() (Schedule, error) {
	defsFound := []string{}

	if s.Year != nil {
		defsFound = append(defsFound, "yearStarting")
		yearErr := s.Year.ValidateNonZero()
		if yearErr != nil {
			return Schedule{}, common.AddValidationContext(yearErr, "yearStarting")
		}
	}

	if s.Month != nil {
		defsFound = append(defsFound, "monthOnDay")
		if *s.Month < 1 || *s.Month > 31 {
			return Schedule{}, common.NewValidationError("monthOnDay", common.NumOutOfRangeCode, "Must be between 1 and 31 (inclusive).")
		}
	}

	if s.HalfMonth != nil {
		defsFound = append(defsFound, "halfMonthOnDays")
		if len(*s.HalfMonth) < 2 {
			return Schedule{}, common.NewValidationError("halfMonthOnDays", common.MissingCode, "You must provide at least two days of the month.")
		}
		for idx, day := range *s.HalfMonth {
			if day < 1 || day > 31 {
				return Schedule{}, common.NewValidationError("halfMonthOnDays."+strconv.Itoa(idx), common.NumOutOfRangeCode, "Must be between 1 and 31 (inclusive).")
			}
		}
	}

	if s.TwoWeeks != nil {
		defsFound = append(defsFound, "twoWeeksStarting")
		biweeklyErr := s.TwoWeeks.ValidateNonZero()
		if biweeklyErr != nil {
			return Schedule{}, common.AddValidationContext(biweeklyErr, "twoWeeksStarting")
		}
	}

	if s.Week != nil {
		defsFound = append(defsFound, "weekOn")
		if !IsDayOfWeek(*s.Week) {
			return Schedule{}, common.NewValidationError("weekOn", common.BadEnumChoiceCode, fmt.Sprintf("You must specify the day of the week: %s", strings.Join(daysOfWeek, ", ")))
		}
	}

	if len(defsFound) != 1 {
		code := common.MissingCode
		if len(defsFound) > 1 {
			code = common.TooManyCode
		}
		return Schedule{}, common.NewValidationError("", code, "You must specify exactly one schedule: yearStarting, monthOnDay, halfMonthOnDays, twoWeeksStarting, weekOn")
	}

	return s, nil
}

var daysOfWeek = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

func IsDayOfWeek(input string) bool {
	for _, day := range daysOfWeek {
		if day == input {
			return true
		}
	}
	return false
}
