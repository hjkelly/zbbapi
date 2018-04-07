package common

import "testing"

func TestAddValidationContext(t *testing.T) {
	inputErr := NewValidationError("foo", "MISSING", "lakjsdf")
	actualErr := AddValidationContext(inputErr, "thisIsContext")
	actualValidationErr, ok := actualErr.(*ValidationError)
	if !ok {
		t.Logf("The result wasn't a *ValidationError")
		t.FailNow()
	}
	if len(actualValidationErr.Fields) != 1 {
		t.Logf("ValidationError didn't have exactly 1 field; had %d", len(actualValidationErr.Fields))
		t.FailNow()
	}
	if actualValidationErr.Fields[0].FieldName != "thisIsContext.foo" {
		t.Logf("ValidationError field didn't have context and field name properly; had %s", actualValidationErr.Fields[0].FieldName)
		t.FailNow()
	}
}
