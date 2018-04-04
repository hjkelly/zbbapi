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

// GetError returns a type-asserted Error if it can.
func GetError(err error) (*Error, bool) {
	if result, ok := err.(Error); ok {
		return &result, ok
	} else if ptrResult, ptrOk := err.(*Error); ptrOk {
		return ptrResult, ptrOk
	} else {
		return nil, false
	}
}

// Satisfies the `error` interface; returns the message.
func (e Error) Error() string {
	return e.Message
}

// Validation Errors
type ValidationError struct {
	BaseError Error        `json:",inline"`
	Fields    []ErrorField `json:"fields,omitempty"`
}

// GetValidationError returns a type-asserted Error if it can.
func GetValidationError(err error) (*ValidationError, bool) {
	if result, ok := err.(ValidationError); ok {
		return &result, ok
	} else if ptrResult, ptrOk := err.(*ValidationError); ptrOk {
		return ptrResult, ptrOk
	} else {
		return nil, false
	}
}

// Satisfies the `error` interface; returns the list of affected fields.
func (e ValidationError) Error() string {
	fieldNames := []string{}
	for _, errErr := range e.Fields {
		fieldNames = append(fieldNames, errErr.FieldName)
	}
	return "Errors involving these fields: " + strings.Join(fieldNames, ", ")
}

// ErrorField is used within Error to denote the problem with a specific field.
type ErrorField struct {
	FieldName string `json:"fieldName"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

// These are codes for errors on fields (`ErrorField`).
const (
	FIELD_MISSING         string = "MISSING"
	FIELD_BAD_ENUM_CHOICE string = "BAD_ENUM_CHOICE"
	FIELD_OUT_OF_RANGE    string = "OUT_OF_RANGE"
)

const INVALID_DATA_MESSAGE = "One or more fields was either missing or invalid."

// NewValidationError creates a new Error with code INVALID_DATA and instantiates a single field error.
func NewValidationError(fieldName, code, message string) *ValidationError {
	return &ValidationError{
		BaseError: Error{
			Code:    "INVALID_DATA",
			Message: INVALID_DATA_MESSAGE,
		},
		Fields: []ErrorField{
			{
				FieldName: fieldName,
				Code:      code,
				Message:   message,
			},
		},
	}
}

// AddValidationContext prefixes any problematic field's names with some parent field.
func (e ValidationError) AddValidationContext(fieldName string) ValidationError {
	for _, field := range e.Fields {
		field.FieldName = fieldName + "." + field.FieldName
	}
	return e
}

// CombineErrors takes several errors and returns a single validation error.
func CombineErrors(errors ...*ValidationError) *ValidationError {
	fields := make([]ErrorField, 0, 5)
	for _, err := range errors {
		if err == nil {
			continue
		}
		for _, field := range err.Fields {
			fields = append(fields, field)
		}
	}
	return &ValidationError{
		BaseError: Error{
			Code:    "INVALID_DATA",
			Message: INVALID_DATA_MESSAGE,
		},
		Fields: fields,
	}
}
