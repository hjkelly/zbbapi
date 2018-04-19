package models

import (
	"strings"

	"github.com/hjkelly/zbbapi/common"
)

var SCHEDULE_TYPES = []string{"yearly", "quarterly", "monthly", "biweekly", "weekly"}

type Schedule struct {
	Type        string       `json:"type"`
	DaysOfMonth []int        `json:"daysOfMonth,omitempty"`
	StartDate   *common.Date `json:"startDate,omitempty"`
}

func IsScheduleType(input string) bool {
	for _, schedType := range SCHEDULE_TYPES {
		if input == schedType {
			return true
		}
	}
	return false
}

func (s Schedule) GetValidated() (Schedule, error) {
	if IsScheduleType(s.Type) == false {
		return s, common.NewValidationError("type", common.BadEnumChoiceCode, "You must choose one of the following schedule types: "+strings.Join(SCHEDULE_TYPES, ", "))
	}
	// If it's monthly, validate the days.
	if s.Type == "monthly" {
		if len(s.DaysOfMonth) == 0 {
			return s, common.NewValidationError("daysOfMonth", common.MissingCode, "With a monthly schedule, you must provide one or more days of the month.")
		}
		for _, day := range s.DaysOfMonth {
			if day < 1 || day > 31 {
				return s, common.NewValidationError("daysOfMonth", common.NumOutOfRangeCode, "Days of the month must be between 1 and 31 (inclusive).")
			}
		}
	} else {
		if s.StartDate == nil || s.StartDate.IsZero() {
			return s, common.NewValidationError("startDate", common.MissingCode, "Unless the schedule is monthly, you must provide a start date.")
		}
		if s.StartDate.IsValid() == false {
			return s, common.NewValidationError("startDate", common.NonexistentDateCode, "This doesn't appear to be a valid date. Perhaps there aren't that many days in this month?")
		}
	}
	return s, nil
}
