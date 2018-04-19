package models

import "testing"

var negativeAmount = Amount{AmountCents: -1}
var zeroAmount = Amount{AmountCents: 0}
var positiveAmount = Amount{AmountCents: 2000}

func TestAmountGetValidated(t *testing.T) {
	var err error
	_, err = negativeAmount.GetValidated()
	if err == nil {
		t.Logf("Didn't get an error for a negative amount.")
		t.Fail()
	}
	_, err = zeroAmount.GetValidated()
	if err != nil {
		t.Logf("Got an error for a zero amount.")
		t.Fail()
	}
	_, err = positiveAmount.GetValidated()
	if err != nil {
		t.Logf("Got an error for a positive amount.")
		t.Fail()
	}
}
