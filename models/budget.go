package models

import (
	"strconv"

	"github.com/hjkelly/zbbapi/common"
)

// -----------------------------------------------------------------------------
// Budget types: Used to plan for your typical/expected month.
// -----------------------------------------------------------------------------

// Budget holds incomes, bills, expenses, and goals that define your expectation/plan for a typical month.
type Budget struct {
	ID       SafeUUID           `json:"id" bson:"_id"`
	Incomes  ManyBudgetIncomes  `json:"incomes"`
	Bills    ManyBudgetBills    `json:"bills"`
	Expenses ManyBudgetExpenses `json:"expenses"`
	Timestamped
}

// GetValidated returns a sanitized copy if all incomes, bills, expenses, and goals are properly defined; otherwise, it returns an error.
func (b Budget) GetValidated() (Budget, error) {
	cleanIncomes, incomesErr := b.Incomes.GetValidated()
	cleanBills, billsErr := b.Bills.GetValidated()
	cleanExpenses, expensesErr := b.Expenses.GetValidated()

	err := common.CombineErrors(
		common.AddValidationContext(incomesErr, "incomes"),
		common.AddValidationContext(billsErr, "bills"),
		common.AddValidationContext(expensesErr, "expenses"),
	)
	if err != nil {
		return Budget{}, err
	}

	b.Incomes = cleanIncomes
	b.Bills = cleanBills
	b.Expenses = cleanExpenses
	return b, nil
}

// BudgetIncome stores a single category reference, its amount within this budget, and the payday schedule.
type BudgetIncome struct {
	CategoryRefAndAmount
	Schedule `json:"schedule"`
}

// GetValidated returns a sanitized copy if the category, amount, and schedule are properly defined; otherwise, it returns an error.
func (i BudgetIncome) GetValidated() (BudgetIncome, error) {
	cleanCRAM, cramErr := i.CategoryRefAndAmount.GetValidated()
	cleanSchedule, scheduleErr := i.Schedule.GetValidated()

	err := common.CombineErrors(
		cramErr,
		common.AddValidationContext(scheduleErr, "schedule"),
	)
	if err != nil {
		return BudgetIncome{}, err
	}

	i.CategoryRefAndAmount = cleanCRAM
	i.Schedule = cleanSchedule
	return i, nil
}

// ManyBudgetIncomes groups several incomes together in a way that makes validation easier.
type ManyBudgetIncomes []BudgetIncome

// GetValidated returns a sanitized copy if all incomes are properly defined; otherwise, it returns an error.
func (incomes ManyBudgetIncomes) GetValidated() (ManyBudgetIncomes, error) {
	errs := make([]error, 0)
	for idx, income := range incomes {
		var incomeErr error
		incomes[idx], incomeErr = income.GetValidated()
		errs = append(errs, common.AddValidationContext(incomeErr, strconv.Itoa(idx)))
	}

	err := common.CombineErrors(errs...)
	if err != nil {
		return ManyBudgetIncomes{}, err
	}

	return incomes, nil
}

// BudgetBill marks a single bill that will need to be paid: the category, expected/average amount, and schedule (along with some other details). Unlike expenses, bills are expected to be a single transaction rather than a 'fund' for many transactions.
type BudgetBill struct {
	CategoryRefAndAmount
	Schedule            `json:"schedule"`
	IsAmountExact       bool `json:"isAmountExact"`
	IsPaidAutomatically bool `json:"isPaidAutomatically"`
}

// GetValidated returns a sanitized copy if the category, amount, and schedule are properly defined; otherwise, it returns an error.
func (b BudgetBill) GetValidated() (BudgetBill, error) {
	cleanCRAM, cramErr := b.CategoryRefAndAmount.GetValidated()
	cleanSchedule, scheduleErr := b.Schedule.GetValidated()

	err := common.CombineErrors(
		cramErr,
		common.AddValidationContext(scheduleErr, "schedule"),
	)
	if err != nil {
		return BudgetBill{}, err
	}

	b.CategoryRefAndAmount = cleanCRAM
	b.Schedule = cleanSchedule
	return b, nil
}

// ManyBudgetBills groups several bills together in a way that makes validation easier.
type ManyBudgetBills []BudgetBill

// GetValidated returns a sanitized copy if all bills are properly defined; otherwise, it returns an error.
func (bills ManyBudgetBills) GetValidated() (ManyBudgetBills, error) {
	errs := make([]error, 0)
	for idx, bill := range bills {
		var billErr error
		bills[idx], billErr = bill.GetValidated()
		errs = append(errs, common.AddValidationContext(billErr, strconv.Itoa(idx)))
	}

	err := common.CombineErrors(errs...)
	if err != nil {
		return ManyBudgetBills{}, err
	}

	return bills, nil
}

// BudgetExpense set aside money to cover costs of expenses. Perhaps you don't use any of it, or perhaps you go over.
type BudgetExpense struct {
	CategoryRefAndAmount
}

// GetValidated returns a sanitized copy if the category and amount are properly defined; otherwise, it returns an error.
func (e BudgetExpense) GetValidated() (BudgetExpense, error) {
	cleanCRAM, cramErr := e.CategoryRefAndAmount.GetValidated()
	if cramErr != nil {
		return BudgetExpense{}, cramErr
	}
	e.CategoryRefAndAmount = cleanCRAM
	return e, nil
}

// ManyBudgetExpenses groups several expenses together in a way that makes validation easier.
type ManyBudgetExpenses []BudgetExpense

// GetValidated returns a sanitized copy if all expenses are properly defined; otherwise, it returns an error.
func (expenses ManyBudgetExpenses) GetValidated() (ManyBudgetExpenses, error) {
	errs := make([]error, 0)
	for idx, expense := range expenses {
		var expenseErr error
		expenses[idx], expenseErr = expense.GetValidated()
		errs = append(errs, common.AddValidationContext(expenseErr, strconv.Itoa(idx)))
	}

	err := common.CombineErrors(errs...)
	if err != nil {
		return ManyBudgetExpenses{}, err
	}

	return expenses, nil
}
