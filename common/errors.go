package common

import (
	"strings"
)

// ParseErr is a reusable error for any time we can't parse a request body, or perhaps the path vars, query params, or header values too.
var ParseErr = &Error{
	Code:    "CANNOT_PARSE",
	Message: "Couldn't parse the data you provided.",
}

// NotFoundErr is a reusable error for any time when we can't find something.
var NotFoundErr = &Error{
	Code:    "NOT_FOUND",
	Message: "One or more resources you referenced couldn't be found.",
}

// Define our own internal types of errors. ----------

// Error is our custom format for passing helpful information around. The main reason for this is so our responses can guess the appropriate status code and also provide helpful info to the client.
type Error struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Fields  []ErrorField `json:"fields,omitempty"`
}

// ErrorField is used within Error to denote the problem with a specific field.
type ErrorField struct {
	FieldName string `json:"fieldName"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

// IsError returns true when the item of interface type `error` is actually one of our errors.
func IsError(err error) bool {
	if _, match := err.(Error); match {
		return true
	}
	return false
}

// Satisfies the `error` interface; returns the message, or list of affected fields if provided.
func (e Error) Error() string {
	if len(e.Fields) > 0 {
		fieldNames := []string{}
		for _, errErr := range e.Fields {
			fieldNames = append(fieldNames, errErr.FieldName)
		}
		return "Errors involving these fields: " + strings.Join(fieldNames, ", ")
	}
	return e.Message
}

// Implement some extra stuff specifically for validation errors. ----------

var validationErrorCode = "INVALID_DATA"

// IsValidationError returns true when the item of interface type `error` is actually one of our errors AND has code type INVALID_DATA. This is useful for knowing if we should respond with 422.
func IsValidationError(err error) bool {
	if ourError, match := err.(Error); match {
		if ourError.Code == validationErrorCode {
			return true
		}
	}
	return false
}

// NewValidationError creates a new Error with code INVALID_DATA and instantiates a single field error.
func NewValidationError(fieldName, code, message string) Error {
	return Error{
		Code:    validationErrorCode,
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
