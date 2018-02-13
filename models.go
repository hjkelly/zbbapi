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
	IsAmountExact       bool
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
	ID                 uuid.UUID
	BudgetID           uuid.UUID
	StartDate          string // TODO date
	EndDate            string // TODO date
	ExactIncomes       []AmountAssignment
	ExactBills         []AmountAssignment
	ExactExpenses      []AmountAssignment
	AdditionalExpenses []CategoryAndAmount
	Checklist          []ChecklistItem
	Calculations       []PayPeriodCalculations
}

type ChecklistItem struct {
	CategoryID uuid.UUID // TODO: include name, or only ID?
	Done       bool
}

type PayPeriodCalculations struct {
	Total      Amount
	Categories []CategoryAndAmount
}

// -----------------------------------------------------------------------------
// Abstract types useful for composition, but without DB stores of their own.
// -----------------------------------------------------------------------------

type Amount struct {
	AmountCents uint
}

type Category struct {
	ID   uuid.UUID
	Name string
}

type CategoryAndAmount struct {
	Category
	Amount
}

type AmountAssignment struct {
	Amount
	CategoryID uuid.UUID
}

type Schedule struct {
	YearlyStartDate   string // TODO: date
	MonthlyOnDays     []uint
	BiweeklyStartDate string // TODO: date
	WeeklyStartDate   string // TODO: date
}
