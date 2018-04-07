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
