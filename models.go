package api

// -----------------------------------------------------------------------------
// Budget types: Used to plan for your typical/expected month.
// -----------------------------------------------------------------------------

type Budget struct {
	ID           uuid.UUID
	Income       []BudgetIncome
	Bills        []BudgetBill
	Expenses     []BudgetExpense
	Goal         []BudgetGoal
	GoalStrategy string
	Paydays      Schedule // TODO: separate from yearly schedule
}

type BudgetIncome struct {
	Amount
	Category
}

// Bills may occur weekly, biweekly, semimonthly, monthly, or annually, but that frequency determines when each transaction will occur.
type BudgetBill struct {
	Amount
	Category
	Schedule
	IsPredictable       bool
	IsPaidAutomatically bool
}

// Expenses set aside money to cover costs of things. Perhaps you don't use any of it, or perhaps you go over.
type BudgetExpenses struct {
	Amount
	Category
	IsAmountExact bool
}

// -----------------------------------------------------------------------------
// Pay period types: Handling actual income/expenses amounts for each pay period.
// -----------------------------------------------------------------------------

type PayPeriod struct {
	ID        uuid.UUID
	BudgetID  uuid.UUID
	StartDate string // TODO date
	EndDate   string // TODO date
}

// -----------------------------------------------------------------------------
// Abstract types useful for composition, but without DB stores of their own.
// -----------------------------------------------------------------------------

type Amount struct {
	AmountCents uint
}

type Category struct {
	Name string
}

type Schedule struct {
	YearlyStartDate   string // TODO: date
	MonthlyOnDays     []uint
	BiweeklyStartDate string // TODO: date
	WeeklyStartDate   string // TODO: date
}
