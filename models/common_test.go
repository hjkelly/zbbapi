package models

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

var negativeAmount = Amount{AmountCents: -1}
var zeroAmount = Amount{AmountCents: 0}
var positiveAmount = Amount{AmountCents: 2000}

func TestAmountValidate(t *testing.T) {
	if negativeAmount.Validate() == nil {
		t.Logf("Didn't get an error for a negative amount.")
		t.Fail()
	}
	if zeroAmount.Validate() != nil {
		t.Logf("Got an error for a zero amount.")
		t.Fail()
	}
	if positiveAmount.Validate() != nil {
		t.Logf("Got an error for a positive amount.")
		t.Fail()
	}
}

func TestCategoryRefAndAmountValidate(t *testing.T) {
	empty := CategoryRefAndAmount{}
	if empty.Validate() == nil {
		t.Logf("Didn't get an error for an empty value.")
		t.Fail()
	}
	missingID := CategoryRefAndAmount{Amount: positiveAmount}
	if missingID.Validate() == nil {
		t.Logf("Didn't get an error with a missing ID.")
		t.Fail()
	}
	badAmount := CategoryRefAndAmount{CategoryID: uuid.NewV4(), Amount: negativeAmount}
	if badAmount.Validate() == nil {
		t.Logf("Didn't get an error with a bad amount.")
		t.Fail()
	}
	valid := CategoryRefAndAmount{CategoryID: uuid.NewV4(), Amount: positiveAmount}
	if valid.Validate() != nil {
		t.Logf("Got an error for what looks like a valid value.")
		t.Fail()
	}
}
