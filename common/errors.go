package common

import (
	"strings"
)

var ParseErr = &Error{
	Code:    "CANNOT_PARSE",
	Message: "Couldn't parse the data you provided.",
}
var NotFoundErr = &Error{
	Code:    "NOT_FOUND",
	Message: "One or more resources you referenced couldn't be found.",
}

// Define our own internal types of errors.

// This is the top-level error, to be used for everything.
type Error struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Fields  []ErrorField `json:"fields,omitempty"`
}

// This is the field-specific error.
type ErrorField struct {
	FieldName string `json:"fieldName"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

func IsError(err error) bool {
	if _, match := err.(Error); match {
		return true
	}
	return false
}

func (e Error) Error() string {
	if len(e.Fields) > 0 {
		fieldNames := []string{}
		for _, errErr := range e.Fields {
			fieldNames = append(fieldNames, errErr.FieldName)
		}
		return "Errors involving these fields: " + strings.Join(fieldNames, ", ")
	} else {
		return e.Message
	}
}

// validation errors ----------

var VALIDATION_ERROR_CODE = "INVALID_DATA"

func IsValidationError(err error) bool {
	if ourError, match := err.(Error); match {
		if ourError.Code == VALIDATION_ERROR_CODE {
			return true
		}
	}
	return false
}

// Create a new validation error with a single field error
func NewValidationError(fieldName, code, message string) Error {
	return Error{
		Code:    VALIDATION_ERROR_CODE,
		Message: "One or more fields was either missing or invalid.",
		Fields: []ErrorField{
			{
				FieldName: fieldName,
				Code:      code,
				Message:   message,
			},
		},
	}
}
