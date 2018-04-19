package models

import (
	"testing"
)

func TestCategoryRefAndAmountGetValidated(t *testing.T) {
	empty := CategoryRefAndAmount{}
	_, err := empty.GetValidated()
	if err == nil {
		t.Logf("Didn't get an error for an empty value.")
		t.Fail()
	}
	missingID := CategoryRefAndAmount{Amount: positiveAmount}
	_, err = missingID.GetValidated()
	if err == nil {
		t.Logf("Didn't get an error with a missing ID.")
		t.Fail()
	}
	badAmount := CategoryRefAndAmount{CategoryID: NewSafeUUID(), Amount: negativeAmount}
	_, err = badAmount.GetValidated()
	if err == nil {
		t.Logf("Didn't get an error with a bad amount.")
		t.Fail()
	}
	valid := CategoryRefAndAmount{CategoryID: NewSafeUUID(), Amount: positiveAmount}
	_, err = valid.GetValidated()
	if err != nil {
		t.Logf("Got an error for what looks like a valid value.")
		t.Fail()
	}
}
