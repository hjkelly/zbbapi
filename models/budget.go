package models

import (
	"strconv"

	"github.com/hjkelly/zbbapi/common"
	uuid "github.com/satori/go.uuid"
)

// -----------------------------------------------------------------------------
// Budget types: Used to plan for your typical/expected month.
// -----------------------------------------------------------------------------

type Budget struct {
	ID       uuid.UUID          `json:"id" bson:"_id"`
	Incomes  ManyBudgetIncomes  `json:"incomes"`
	Bills    ManyBudgetBills    `json:"bills"`
	Expenses ManyBudgetExpenses `json:"expenses"`
	Timestamped
}

func (b Budget) Validate() error {
	return common.CombineErrors(
		b.Incomes.Validate(),
		b.Bills.Validate(),
		b.Expenses.Validate(),
	)
}

type BudgetIncome struct {
	CategoryRefAndAmount
	Schedule `json:"schedule"`
}

func (i BudgetIncome) Validate() error {
	return common.CombineErrors(
		i.CategoryRefAndAmount.Validate(),
		common.AddValidationContext(i.Schedule.Validate(), "schedule"),
	)
}

type ManyBudgetIncomes []BudgetIncome

func (incomes ManyBudgetIncomes) Validate() error {
	errs := make([]error, 0)
	for idx, income := range incomes {
		errs = append(errs, common.AddValidationContext(income.Validate(), strconv.Itoa(idx)))
	}
	return common.AddValidationContext(common.CombineErrors(errs...), "incomes")
}

// Bills may occur weekly, biweekly, semimonthly, monthly, or annually, but that frequency determines when each transaction should occur.
type BudgetBill struct {
	CategoryRefAndAmount
	Schedule            `json:"schedule"`
	IsAmountExact       bool `json:"isAmountExact"`
	IsPaidAutomatically bool `json:"isPaidAutomatically"`
}

func (i BudgetBill) Validate() error {
	return common.CombineErrors(
		i.CategoryRefAndAmount.Validate(),
		common.AddValidationContext(i.Schedule.Validate(), "schedule"),
	)
}

type ManyBudgetBills []BudgetBill

func (bills ManyBudgetBills) Validate() error {
	errs := make([]error, 0)
	for idx, bill := range bills {
		errs = append(errs, common.AddValidationContext(bill.Validate(), strconv.Itoa(idx)))
	}
	return common.AddValidationContext(common.CombineErrors(errs...), "bills")
}

// Expenses set aside money to cover costs of things. Perhaps you don't use any of it, or perhaps you go over.
type BudgetExpense struct {
	CategoryRefAndAmount
}

func (i BudgetExpense) Validate() error {
	return i.CategoryRefAndAmount.Validate()
}

type ManyBudgetExpenses []BudgetExpense

func (expenses ManyBudgetExpenses) Validate() error {
	errs := make([]error, 0)
	for idx, expense := range expenses {
		errs = append(errs, common.AddValidationContext(expense.Validate(), strconv.Itoa(idx)))
	}
	return common.AddValidationContext(common.CombineErrors(errs...), "expenses")
}
