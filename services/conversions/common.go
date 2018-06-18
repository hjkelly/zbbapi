package conversions

import (
	"strings"

	"github.com/hjkelly/zbbapi/common"
	"github.com/hjkelly/zbbapi/models"
	"github.com/hjkelly/zbbapi/services/plans"
)

// Make sure this Conversion has input sufficient enough to be saved.
func getValidated(input models.Conversion) (models.Conversion, models.Plan, error) {
	input, err := input.GetValidated()
	if err != nil {
		return models.Conversion{}, models.Plan{}, err
	}
	// fetch plan (validating it exists)
	plan, err := plans.Retrieve(string(input.PlanID))
	if err != nil {
		return models.Conversion{}, models.Plan{}, err
	}
	// which of the plan's bills need exact amounts?
	inconsistentBillNames := []string{}
	for _, bill := range plan.Bills {
		if !bill.IsAmountExact {
			inconsistentBillNames = append(inconsistentBillNames, bill.Name)
		}
	}
	// does the conversion provide all inconsistent values?
	for _, exactBill := range input.ExactBills {
		for index, billName := range inconsistentBillNames {
			if exactBill.Name == billName {
				// remove the element since its exact amount is accounted for
				inconsistentBillNames = append(
					inconsistentBillNames[:index],
					inconsistentBillNames[index+1:]...,
				)
				break
			}
		}
	}
	if len(inconsistentBillNames) > 0 {
		return input, plan, common.NewValidationError(common.MissingCode, "exactBills", "You're missing exact amounts for the following bills: "+strings.Join(inconsistentBillNames, ", "))
	}
	return input, plan, nil
}

// Returns the updated Conversion, which is the current Conversion updated with the input data for the update.
func getUpdated(current, input models.Conversion) models.Conversion {
	return current
}
