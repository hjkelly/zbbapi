// -----------------------------------------------------------------------------
// Abstract types useful for composition, but without DB stores of their own.
// -----------------------------------------------------------------------------
package models

import (
	"log"
	"strings"
	"time"

	"github.com/hjkelly/zbbapi/common"
	uuid "github.com/satori/go.uuid"
)

type Amount struct {
	AmountCents int64 `json:"amount"`
}

func (a Amount) Validate() *common.ValidationError {
	if a.AmountCents < 0 {
		return common.NewValidationError("amount", common.FIELD_OUT_OF_RANGE, "The amount cannot be negative.")
	}
	return nil
}

type CategoryRefAndAmount struct {
	CategoryID uuid.UUID `json:"categoryID"`
	Amount
}

func (cram CategoryRefAndAmount) Validate() *common.ValidationError {
	idErr := new(common.ValidationError)
	if cram.CategoryID == uuid.FromBytesOrNil([]byte{}) {
		idErr = common.NewValidationError("categoryID", common.FIELD_MISSING, "You must reference a category's ID, which is a UUID.")
	}
	return common.CombineErrors(
		cram.Amount.Validate(),
		idErr,
	)
}

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

func (s Schedule) Validate() *common.ValidationError {
	if IsScheduleType(s.Type) == false {
		return common.NewValidationError("type", common.FIELD_BAD_ENUM_CHOICE, "You must choose one of the following schedule types: "+strings.Join(SCHEDULE_TYPES, ", "))
	}
	// If it's monthly, validate the days.
	if s.Type == "monthly" {
		log.Printf("checking days of week\n")
		if len(s.DaysOfMonth) == 0 {
			return common.NewValidationError("daysOfMonth", common.FIELD_MISSING, "With a monthly schedule, you must provide one or more days of the month.")
		}
		log.Printf("they provided 1+\n")
		for _, day := range s.DaysOfMonth {
			if day < 1 || day > 31 {
				return common.NewValidationError("daysOfMonth", common.FIELD_OUT_OF_RANGE, "Days of the month must be between 1 and 31 (inclusive).")
			}
		}
		log.Printf("they were all valid\n")
	} else {
		log.Printf("checking start date\n")
		if s.StartDate == nil || s.StartDate.IsZero() {
			return common.NewValidationError("startDate", common.FIELD_MISSING, "Unless the schedule is monthly, you must provide a start date.")
		}
		log.Printf("date isn't zero: %s\n", s.StartDate)
		if s.StartDate.IsValid() == false {
			return common.NewValidationError("startDate", common.FIELD_OUT_OF_RANGE, "This doesn't appear to be a valid date. Perhaps there aren't that many days in this month?")
		}
		log.Printf("date is valid: %s\n", s.StartDate)
	}
	return nil
}

type Timestamped struct {
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func (t *Timestamped) SetCreationTimestamp() {
	t.Created = time.Now()
	t.Modified = time.Now()
}

func (t *Timestamped) SetModificationTimestamp() {
	t.Modified = time.Now()
}
