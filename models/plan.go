package models

import (
	"strconv"
	"strings"

	"github.com/hjkelly/zbbapi/common"
)

// -----------------------------------------------------------------------------
// Plan types: Used to plan for your typical/expected month.
// -----------------------------------------------------------------------------

// Plan holds incomes, bills, expenses, and goals that define your expectation/plan for a typical month.
type Plan struct {
	ID              SafeUUID            `json:"id" bson:"_id"`
	Incomes         ManyPlannedIncomes  `json:"incomes"`
	Bills           ManyPlannedBills    `json:"bills"`
	Expenses        ManyPlannedExpenses `json:"expenses"`
	Savings         ManyPlannedSavings  `json:"savings"`
	SavingsStrategy string              `json:"savingsStrategy"`
	Timestamped
}

// GetValidated returns a sanitized copy if all incomes, bills, expenses, and goals are properly defined; otherwise, it returns an error.
func (plan Plan) GetValidated() (Plan, error) {
	cleanIncomes, incomesErr := plan.Incomes.GetValidated()
	cleanBills, billsErr := plan.Bills.GetValidated()
	cleanExpenses, expensesErr := plan.Expenses.GetValidated()
	cleanSavings, savingsErr := plan.Savings.GetValidated()
	var savingsStrategyErr error
	if !IsSavingsStrategy(plan.SavingsStrategy) {
		savingsStrategyErr = common.NewValidationError("savingsStrategy", common.BadEnumChoiceCode, "You must provide a valid strategy: %s", strings.Join(SavingsStrategies, ", "))
	}

	err := common.CombineErrors(
		common.AddValidationContext(incomesErr, "incomes"),
		common.AddValidationContext(billsErr, "bills"),
		common.AddValidationContext(expensesErr, "expenses"),
		common.AddValidationContext(savingsErr, "savings"),
		savingsStrategyErr,
	)
	if err != nil {
		return Plan{}, err
	}

	plan.Incomes = cleanIncomes
	plan.Bills = cleanBills
	plan.Expenses = cleanExpenses
	plan.Savings = cleanSavings
	return plan, nil
}

// INCOMES ----------

// PlannedIncome stores a single category reference, its amount within this budget, and the payday schedule.
type PlannedIncome struct {
	NameAndAmount
	Schedule `json:"every"`
}

// GetValidated returns a sanitized copy if the category, amount, and schedule are properly defined; otherwise, it returns an error.
func (income PlannedIncome) GetValidated() (PlannedIncome, error) {
	cleanCRAM, cramErr := income.NameAndAmount.GetValidated()
	cleanSchedule, scheduleErr := income.Schedule.GetValidated()

	err := common.CombineErrors(
		cramErr,
		common.AddValidationContext(scheduleErr, "every"),
	)
	if err != nil {
		return PlannedIncome{}, err
	}

	income.NameAndAmount = cleanCRAM
	income.Schedule = cleanSchedule
	return income, nil
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
	Schedule            `json:"every"`
	IsAmountExact       bool `json:"isAmountExact"`
	IsPaidAutomatically bool `json:"isPaidAutomatically"`
}

// GetValidated returns a sanitized copy if the category, amount, and schedule are properly defined; otherwise, it returns an error.
func (bill PlannedBill) GetValidated() (PlannedBill, error) {
	cleanCRAM, cramErr := bill.NameAndAmount.GetValidated()
	cleanSchedule, scheduleErr := bill.Schedule.GetValidated()

	err := common.CombineErrors(
		cramErr,
		common.AddValidationContext(scheduleErr, "every"),
	)
	if err != nil {
		return PlannedBill{}, err
	}

	bill.NameAndAmount = cleanCRAM
	bill.Schedule = cleanSchedule
	return bill, nil
}

// BILLS ----------

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

// EXPENSES ----------

// PlannedExpense set aside money to cover costs of expenses. Perhaps you don't use any of it, or perhaps you go over.
type PlannedExpense struct {
	NameAndAmount
}

// GetValidated returns a sanitized copy if the category and amount are properly defined; otherwise, it returns an error.
func (expense PlannedExpense) GetValidated() (PlannedExpense, error) {
	cleanCRAM, cramErr := expense.NameAndAmount.GetValidated()
	if cramErr != nil {
		return PlannedExpense{}, cramErr
	}
	expense.NameAndAmount = cleanCRAM
	return expense, nil
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

// SAVINGS ----------

// PlannedSaving set aside money to cover costs of expenses. Perhaps you don't use any of it, or perhaps you go over.
type PlannedSaving struct {
	NameAndAmount
}

// GetValidated returns a sanitized copy if the category and amount are properly defined; otherwise, it returns an error.
func (saving PlannedSaving) GetValidated() (PlannedSaving, error) {
	cleanCRAM, cramErr := saving.NameAndAmount.GetValidated()
	if cramErr != nil {
		return PlannedSaving{}, cramErr
	}
	saving.NameAndAmount = cleanCRAM
	return saving, nil
}

// ManyPlannedSavings groups several savings together in a way that makes validation easier.
type ManyPlannedSavings []PlannedSaving

// GetValidated returns a sanitized copy if all savings are properly defined; otherwise, it returns an error.
func (savings ManyPlannedSavings) GetValidated() (ManyPlannedSavings, error) {
	errs := make([]error, 0)
	for idx, saving := range savings {
		var savingErr error
		savings[idx], savingErr = saving.GetValidated()
		errs = append(errs, common.AddValidationContext(savingErr, strconv.Itoa(idx)))
	}

	err := common.CombineErrors(errs...)
	if err != nil {
		return ManyPlannedSavings{}, err
	}

	return savings, nil
}

var SavingsStrategies = []string{"shared", "prioritized"}

func IsSavingsStrategy(input string) bool {
	for _, ss := range SavingsStrategies {
		if input == ss {
			return true
		}
	}
	return false
}
