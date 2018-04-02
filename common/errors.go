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

// Error is our custom format for passing helpful information around. The main reason for this is so our responses can guess the appropriate status code and also provide helpful info to the client.
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Validation Errors
type ValidationError struct {
	Error
	Fields []ErrorField `json:"fields,omitempty"`
}

// ErrorField is used within Error to denote the problem with a specific field.
type ErrorField struct {
	FieldName string `json:"fieldName"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

// Satisfies the `error` interface; returns the message, or list of affected fields if provided.
func (e ValidationError) Error() string {
	fieldNames := []string{}
	for _, errErr := range e.Fields {
		fieldNames = append(fieldNames, errErr.FieldName)
	}
	return "Errors involving these fields: " + strings.Join(fieldNames, ", ")
}

// Implement some extra stuff specifically for validation errors. ----------

// IsValidationError returns true when the item of interface type `error` is actually one of our errors AND has code type INVALID_DATA. This is useful for knowing if we should respond with 422.

const (
	FILED_MISSING         string = "MISSING"
	FIELD_BAD_ENUM_CHOICE string = "BAD_ENUM_CHOICE"
	FIELD_OUT_OF_RANGE    string = "OUT_OF_RANGE"
)

// NewValidationError creates a new Error with code INVALID_DATA and instantiates a single field error.
func NewValidationError(fieldName, code, message string) Error {
	return Error{
		Code:    "INVALID_DATA",
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

// AddValidationContext prefixes any validation errors with fieldnames.
func (e ValidationError) AddValidationContext(fieldName string) ValidationError {
	// TODO
}

// CombineErrors takes several errors and returns a single validation error.
func CombineErrors(errors ...[]Error) *ValidationError {
	for _, err := range errors {
	}
}
