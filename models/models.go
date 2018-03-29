package models

import (
	"time"

	"github.com/GoogleCloudPlatform/google-cloud-go/civil"
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

// Bills may occur weekly, biweekly, semimonthly, monthly, or annually, but that frequency determines when each transaction should occur.
type BudgetBill struct {
	CategoryRefAndAmount
	Schedule
	IsAmountPredictable bool `json:"isAmountPredictable"`
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
	StartDate          civil.Date
	EndDate            civil.Date
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

type Schedule struct {
	YearlyStartDate   *civil.Date `json:"yearlyStartDate,omitempty"`
	MonthlyOnDays     []uint      `json:"monthlyOnDays,omitempty"`
	BiweeklyStartDate *civil.Date `json:"biweeklyStartDate,omitempty"`
	WeeklyStartDate   *civil.Date `json:"weeklyStartDate,omitempty"`
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
