package models

import (
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

func (b Budget) Validate() *common.ValidationError {
	possibleErrors := []*common.ValidationError{}
	for _, income := range b.Income {
		possibleErrors = append(possibleErrors, income.Validate())
	}
	for _, bill := range b.Bills {
		possibleErrors = append(possibleErrors, bill.Validate())
	}
	for _, expense := range b.Expenses {
		possibleErrors = append(possibleErrors, expense.Validate())
	}
	return common.CombineErrors(possibleErrors...)
}

type BudgetIncome struct {
	CategoryRefAndAmount
	Schedule `json:"schedule"`
}

func (i BudgetIncome) Validate() *common.ValidationError {
	return common.CombineErrors(
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

func (i BudgetBill) Validate() *common.ValidationError {
	return nil
}

// Expenses set aside money to cover costs of things. Perhaps you don't use any of it, or perhaps you go over.
type BudgetExpense struct {
	CategoryRefAndAmount
}

func (i BudgetExpense) Validate() *common.ValidationError {
	return nil
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
