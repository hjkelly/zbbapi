package api

import uuid "github.com/satori/go.uuid"

// -----------------------------------------------------------------------------
// Category types: Obviously a name, but also an ID for specificity and name changes.
// -----------------------------------------------------------------------------

type Category struct {
	ID   uuid.UUID
	Name string
}

// -----------------------------------------------------------------------------
// Budget types: Used to plan for your typical/expected month.
// -----------------------------------------------------------------------------

type Budget struct {
	ID       uuid.UUID
	Paydays  Schedule // TODO: separate from yearly schedule
	Income   []BudgetIncome
	Bills    []BudgetBill
	Expenses []BudgetExpense
	//Goal         []BudgetGoal
	//GoalStrategy string
}

type BudgetIncome struct {
	CategoryRefAndAmount
}

// Bills may occur weekly, biweekly, semimonthly, monthly, or annually, but that frequency determines when each transaction will occur.
type BudgetBill struct {
	CategoryRefAndAmount
	Schedule
	IsAmountExact       bool
	IsPaidAutomatically bool
}

// Expenses set aside money to cover costs of things. Perhaps you don't use any of it, or perhaps you go over.
type BudgetExpense struct {
	CategoryRefAndAmount
	IsAmountExact bool
}

// -----------------------------------------------------------------------------
// Pay period types: Handling actual income/expenses amounts for each pay period.
// -----------------------------------------------------------------------------

type PayPeriod struct {
	ID                 uuid.UUID
	BudgetID           uuid.UUID
	StartDate          string // TODO date
	EndDate            string // TODO date
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
	AmountCents uint
}

type CategoryRefAndAmount struct {
	CategoryID uuid.UUID
	Amount
}

type Schedule struct {
	YearlyStartDate   string // TODO: date
	MonthlyOnDays     []uint
	BiweeklyStartDate string // TODO: date
	WeeklyStartDate   string // TODO: date
}
