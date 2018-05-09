package models

import (
	"strconv"

	"github.com/hjkelly/zbbapi/common"
)

// -----------------------------------------------------------------------------
// Plan types: Used to plan for your typical/expected month.
// -----------------------------------------------------------------------------

// Plan holds incomes, bills, expenses, and goals that define your expectation/plan for a typical month.
type Plan struct {
	ID       SafeUUID            `json:"id" bson:"_id"`
	Incomes  ManyPlannedIncomes  `json:"incomes"`
	Bills    ManyPlannedBills    `json:"bills"`
	Expenses ManyPlannedExpenses `json:"expenses"`
	Timestamped
}

// GetValidated returns a sanitized copy if all incomes, bills, expenses, and goals are properly defined; otherwise, it returns an error.
func (b Plan) GetValidated() (Plan, error) {
	cleanIncomes, incomesErr := b.Incomes.GetValidated()
	cleanBills, billsErr := b.Bills.GetValidated()
	cleanExpenses, expensesErr := b.Expenses.GetValidated()

	err := common.CombineErrors(
		common.AddValidationContext(incomesErr, "incomes"),
		common.AddValidationContext(billsErr, "bills"),
		common.AddValidationContext(expensesErr, "expenses"),
	)
	if err != nil {
		return Plan{}, err
	}

	b.Incomes = cleanIncomes
	b.Bills = cleanBills
	b.Expenses = cleanExpenses
	return b, nil
}

// PlannedIncome stores a single category reference, its amount within this budget, and the payday schedule.
type PlannedIncome struct {
	NameAndAmount
	Schedule `json:"schedule"`
}

// GetValidated returns a sanitized copy if the category, amount, and schedule are properly defined; otherwise, it returns an error.
func (i PlannedIncome) GetValidated() (PlannedIncome, error) {
	cleanCRAM, cramErr := i.NameAndAmount.GetValidated()
	cleanSchedule, scheduleErr := i.Schedule.GetValidated()

	err := common.CombineErrors(
		cramErr,
		common.AddValidationContext(scheduleErr, "schedule"),
	)
	if err != nil {
		return PlannedIncome{}, err
	}

	i.NameAndAmount = cleanCRAM
	i.Schedule = cleanSchedule
	return i, nil
}

// ManyPlannedIncomes groups several incomes together in a way that makes validation easier.
type ManyPlannedIncomes []PlannedIncome

// GetValidated returns a sanitized copy if all incomes are properly defined; otherwise, it returns an error.
func (incomes ManyPlannedIncomes) GetValidated() (ManyPlannedIncomes, error) {
	errs := make([]error, 0)
	for idx, income := range incomes {
		var incomeErr error
		incomes[idx], incomeErr = income.GetValidated()
		errs = append(errs, common.AddValidationContext(incomeErr, strconv.Itoa(idx)))
	}

	err := common.CombineErrors(errs...)
	if err != nil {
		return ManyPlannedIncomes{}, err
	}

	return incomes, nil
}

// PlannedBill marks a single bill that will need to be paid: the category, expected/average amount, and schedule (along with some other details). Unlike expenses, bills are expected to be a single transaction rather than a 'fund' for many transactions.
type PlannedBill struct {
	NameAndAmount
	Schedule            `json:"schedule"`
	IsAmountExact       bool `json:"isAmountExact"`
	IsPaidAutomatically bool `json:"isPaidAutomatically"`
}

// GetValidated returns a sanitized copy if the category, amount, and schedule are properly defined; otherwise, it returns an error.
func (b PlannedBill) GetValidated() (PlannedBill, error) {
	cleanCRAM, cramErr := b.NameAndAmount.GetValidated()
	cleanSchedule, scheduleErr := b.Schedule.GetValidated()

	err := common.CombineErrors(
		cramErr,
		common.AddValidationContext(scheduleErr, "schedule"),
	)
	if err != nil {
		return PlannedBill{}, err
	}

	b.NameAndAmount = cleanCRAM
	b.Schedule = cleanSchedule
	return b, nil
}

// ManyPlannedBills groups several bills together in a way that makes validation easier.
type ManyPlannedBills []PlannedBill

// GetValidated returns a sanitized copy if all bills are properly defined; otherwise, it returns an error.
func (bills ManyPlannedBills) GetValidated() (ManyPlannedBills, error) {
	errs := make([]error, 0)
	for idx, bill := range bills {
		var billErr error
		bills[idx], billErr = bill.GetValidated()
		errs = append(errs, common.AddValidationContext(billErr, strconv.Itoa(idx)))
	}

	err := common.CombineErrors(errs...)
	if err != nil {
		return ManyPlannedBills{}, err
	}

	return bills, nil
}

// PlannedExpense set aside money to cover costs of expenses. Perhaps you don't use any of it, or perhaps you go over.
type PlannedExpense struct {
	NameAndAmount
}

// GetValidated returns a sanitized copy if the category and amount are properly defined; otherwise, it returns an error.
func (e PlannedExpense) GetValidated() (PlannedExpense, error) {
	cleanCRAM, cramErr := e.NameAndAmount.GetValidated()
	if cramErr != nil {
		return PlannedExpense{}, cramErr
	}
	e.NameAndAmount = cleanCRAM
	return e, nil
}

// ManyPlannedExpenses groups several expenses together in a way that makes validation easier.
type ManyPlannedExpenses []PlannedExpense

// GetValidated returns a sanitized copy if all expenses are properly defined; otherwise, it returns an error.
func (expenses ManyPlannedExpenses) GetValidated() (ManyPlannedExpenses, error) {
	errs := make([]error, 0)
	for idx, expense := range expenses {
		var expenseErr error
		expenses[idx], expenseErr = expense.GetValidated()
		errs = append(errs, common.AddValidationContext(expenseErr, strconv.Itoa(idx)))
	}

	err := common.CombineErrors(errs...)
	if err != nil {
		return ManyPlannedExpenses{}, err
	}

	return expenses, nil
}
