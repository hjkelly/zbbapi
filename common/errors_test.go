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

func TestCombineErrors(t *testing.T) {
	actualErr := CombineErrors(
		NewValidationError("one", "code1", "message1"),
		NewValidationError("two", "code2", "message2"),
	)
	actualValidationErr, ok := actualErr.(*ValidationError)
	if !ok {
		t.Logf("The result wasn't a *ValidationError")
		t.FailNow()
	}
	if len(actualValidationErr.Fields) != 2 {
		t.Logf("ValidationError didn't have exactly 2 field; had %d", len(actualValidationErr.Fields))
		t.FailNow()
	}
	fieldsFound := map[string]bool{}
	for _, field := range actualValidationErr.Fields {
		if field.FieldName == "one" {
			if field.Code != "code1" {
				t.Logf("Got wrong code for field 1: %s", field.Code)
				t.Fail()
			}
			if field.Message != "message1" {
				t.Logf("Got wrong message for field 1: %s", field.Message)
				t.Fail()
			}
		} else if field.FieldName == "two" {
			if field.Code != "code2" {
				t.Logf("Got wrong code for field 2: %s", field.Code)
				t.Fail()
			}
			if field.Message != "message2" {
				t.Logf("Got wrong message for field 2: %s", field.Message)
				t.Fail()
			}
		} else {
			t.Logf("Got unexpected field in combined error: %s", field.FieldName)
			t.FailNow()
		}
		fieldsFound[field.FieldName] = true
	}
	if len(fieldsFound) != 2 {
		t.Logf("Combined error didn't have two distinct fields on it.")
		t.FailNow()
	}
}
