package conversions

import (
	"github.com/hjkelly/zbbapi/models"
	"github.com/hjkelly/zbbapi/services/budgets"
)

func Create(input models.Conversion) (models.Conversion, error) {
	// Did they give us enough to save?
	input, plan, err := getValidated(input)
	if err != nil {
		return input, err
	}

	budget := input.MakeBudget(plan)
	budget, err = budgets.Create(budget)
	if err != nil {
		return models.Conversion{}, err
	}

	// prepare the rest of the resource
	input.ID = models.NewSafeUUID()
	input.BudgetID = budget.ID
	input.SetCreationTimestamp()

	// save
	ds := newDatastore()
	err = ds.C().Insert(input)
	if err != nil {
		return input, err
	}
	return input, nil
}
