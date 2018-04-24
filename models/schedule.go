package models

import (
	"strings"

	"github.com/hjkelly/zbbapi/common"
)

// These are the valid schedule types.
var SCHEDULE_TYPES = []string{"yearly", "quarterly", "monthly", "biweekly", "weekly"}

// Schedule defines when an event is expected to occur. Depending on the type, you'll have to provide at least one other piece of info so the event can be predicted.
type Schedule struct {
	Type        string       `json:"type"`
	DaysOfMonth []int        `json:"daysOfMonth,omitempty"`
	StartDate   *common.Date `json:"startDate,omitempty"`
}

// IsValidScheduleType ensures the string is a valid type.
func IsValidScheduleType(input string) bool {
	for _, schedType := range SCHEDULE_TYPES {
		if input == schedType {
			return true
		}
	}
	return false
}

// GetValidated returns a sanitized schedule if the type and other sometimes-required attributes are in order; otherwise, returns an error.
func (s Schedule) GetValidated() (Schedule, error) {
	if IsValidScheduleType(s.Type) == false {
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
