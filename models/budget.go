package models

import (
	"strconv"

	"github.com/hjkelly/zbbapi/common"
)

// -----------------------------------------------------------------------------
// Budget types: Used to plan for your typical/expected month.
// -----------------------------------------------------------------------------

type Budget struct {
	ID       SafeUUID           `json:"id" bson:"_id"`
	Incomes  ManyBudgetIncomes  `json:"incomes"`
	Bills    ManyBudgetBills    `json:"bills"`
	Expenses ManyBudgetExpenses `json:"expenses"`
	Timestamped
}

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

type BudgetIncome struct {
	CategoryRefAndAmount
	Schedule `json:"schedule"`
}

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

type ManyBudgetIncomes []BudgetIncome

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

// Bills may occur weekly, biweekly, semimonthly, monthly, or annually, but that frequency determines when each transaction should occur.
type BudgetBill struct {
	CategoryRefAndAmount
	Schedule            `json:"schedule"`
	IsAmountExact       bool `json:"isAmountExact"`
	IsPaidAutomatically bool `json:"isPaidAutomatically"`
}

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

type ManyBudgetBills []BudgetBill

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

// Expenses set aside money to cover costs of things. Perhaps you don't use any of it, or perhaps you go over.
type BudgetExpense struct {
	CategoryRefAndAmount
}

func (e BudgetExpense) GetValidated() (BudgetExpense, error) {
	cleanCRAM, cramErr := e.CategoryRefAndAmount.GetValidated()
	if cramErr != nil {
		return BudgetExpense{}, cramErr
	}
	e.CategoryRefAndAmount = cleanCRAM
	return e, nil
}

type ManyBudgetExpenses []BudgetExpense

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
