package models

import (
	"strings"
	"time"

	"github.com/hjkelly/zbbapi/common"
	uuid "github.com/satori/go.uuid"
)

// -----------------------------------------------------------------------------
// Category types: Obviously a name, but also an ID for specificity and name changes.
// -----------------------------------------------------------------------------

type Category struct {
	ID   uuid.UUID `json:"id" bson:"_id"`
	Name string    `json:"name"`
	Timestamped
}

// -----------------------------------------------------------------------------
// Budget types: Used to plan for your typical/expected month.
// -----------------------------------------------------------------------------

type Budget struct {
	ID       uuid.UUID       `json:"id" bson:"_id"`
	Income   []BudgetIncome  `json:"income"`
	Bills    []BudgetBill    `json:"bills"`
	Expenses []BudgetExpense `json:"expenses"`
	//Goal         []BudgetGoal
	//GoalStrategy string
	Timestamped
}

type BudgetIncome struct {
	CategoryRefAndAmount
	Schedule `json:"schedule"`
}

func (i BudgetIncome) Validate() *common.ValidationError {
	return CombineErrors(
		i.CategoryRefAndAmount.Validate(),
		i.Schedule.Validate(),
	)
}

// Bills may occur weekly, biweekly, semimonthly, monthly, or annually, but that frequency determines when each transaction should occur.
type BudgetBill struct {
	CategoryRefAndAmount
	Schedule
	IsAmountExact       bool `json:"isAmountExact"`
	IsPaidAutomatically bool `json:"isPaidAutomatically"`
}

// Expenses set aside money to cover costs of things. Perhaps you don't use any of it, or perhaps you go over.
type BudgetExpense struct {
	CategoryRefAndAmount
}

// -----------------------------------------------------------------------------
// Pay period types: Handling actual income/expenses amounts for each pay period.
// -----------------------------------------------------------------------------

type PayPeriod struct {
	ID                 uuid.UUID
	BudgetID           uuid.UUID
	StartDate          common.Date
	EndDate            common.Date
	ExactIncomes       []CategoryRefAndAmount
	ExactBills         []CategoryRefAndAmount
	ExactExpenses      []CategoryRefAndAmount
	AdditionalExpenses []CategoryRefAndAmount
	Checklist          []ChecklistItem
	Calculations       []PayPeriodCalculations
}

type ChecklistItem struct {
	CategoryID uuid.UUID // TODO: include name, or only ID?
	Done       bool
}

type PayPeriodCalculations struct {
	Total      Amount
	Categories []CategoryRefAndAmount
}

// -----------------------------------------------------------------------------
// Abstract types useful for composition, but without DB stores of their own.
// -----------------------------------------------------------------------------

type Amount struct {
	AmountCents uint64 `json:"amount"`
}

type CategoryRefAndAmount struct {
	CategoryID uuid.UUID `json:"categoryID"`
	Amount
}

const SCHEDULE_TYPES = []string{"yearly", "quarterly", "monthly", "biweekly", "weekly"}

type Schedule struct {
	Type        string       `json:"type"`
	DaysOfMonth []int        `json:"daysOfMonth,omitempty"`
	StartDate   *common.Date `json:"startDate,omitempty"`
}

func (s Schedule) Validate() *common.ValidationError {
	if _, ok := SCHEDULE_TYPE_MAP[s.Type]; ok == False {
		return common.NewValidationError("type", common.FIELD_BAD_ENUM_CHOICE, "You must choose one of the following schedule types: "+strings.Join(SCHEDULE_TYPES, ", "))
	}
	if s.Type == "monthly" {
		if len(s.DaysOfMonth) == 0 {
			return common.NewValidationError("daysOfMonth", common.FIELD_MISSING, "With a monthly schedule, you must provide one or more days of the month.")
		}
		for _, day := range s.DaysOfMonth {
			if day < 1 || day > 31 {
				return common.NewValidationError("daysOfMonth", common.FIELD_OUT_OF_RANGE, "Days of the month must be between 1 and 31 (inclusive).")
			}
		}
	} else {
		if s.StartDate == nil || s.StartDate.IsZero() {
			return common.NewValidationError("daysOfMonth", common.FIELD_MISSING, "Unless the schedule is monthly, you must provide a start date.")
		}
		if s.StartDate.IsValid() == false {
			return common.NewValidationError("daysOfMonth", common.FIELD_OUT_OF_RANGE, "This doesn't appear to be a valid date. Perhaps there aren't that many days in this month?")
		}
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
