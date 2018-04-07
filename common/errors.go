package common

import (
	"log"
	"strings"
)

// ParseErr is a reusable error for any time we can't parse a request body, or perhaps the path vars, query params, or header values too.
var ParseErr = &BasicError{
	Code:    "CANNOT_PARSE",
	Message: "Couldn't parse the data you provided.",
}

// NotFoundErr is a reusable error for any time when we can't find something.
var NotFoundErr = &BasicError{
	Code:    "NOT_FOUND",
	Message: "One or more resources you referenced couldn't be found.",
}

// BasicError is our custom format for passing helpful information around. The main reason for this is so our responses can guess the appropriate status code and also provide helpful info to the client.
type BasicError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// GetError returns a type-asserted BasicError if it can.
func GetError(err error) (*BasicError, bool) {
	if result, ok := err.(BasicError); ok {
		return &result, ok
	} else if ptrResult, ptrOk := err.(*BasicError); ptrOk {
		return ptrResult, ptrOk
	} else {
		return nil, false
	}
}

// Satisfies the `error` interface; returns the message.
func (e BasicError) Error() string {
	return e.Message
}

// Validation Errors
type ValidationError struct {
	BasicError
	Fields []InvalidField `json:"fields,omitempty"`
}

// GetValidationError returns a type-asserted BasicError if it can.
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

// InvalidField is used within ValidationError to denote the problem with a specific field.
type InvalidField struct {
	FieldName string `json:"fieldName"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

// These are codes for errors on fields (`InvalidField`).
const (
	FIELD_MISSING         string = "MISSING"
	FIELD_BAD_ENUM_CHOICE string = "BAD_ENUM_CHOICE"
	FIELD_OUT_OF_RANGE    string = "OUT_OF_RANGE"
)

const INVALID_DATA_MESSAGE = "One or more fields was either missing or invalid."

// NewValidationError creates a new ValidationError with code INVALID_DATA and instantiates a single field error.
func NewValidationError(fieldName, code, message string) *ValidationError {
	return &ValidationError{
		BasicError: BasicError{
			Code:    "INVALID_DATA",
			Message: INVALID_DATA_MESSAGE,
		},
		Fields: []InvalidField{
			{
				FieldName: fieldName,
				Code:      code,
				Message:   message,
			},
		},
	}
}

// AddValidationContext prefixes any problematic field's names with some parent field.
func AddValidationContext(err error, fieldName string) error {
	if err == nil {
		return nil
	}
	validationErr, ok := err.(*ValidationError)
	if !ok {
		log.Printf("One of the errors passed to PrefixFieldNames was not a ValidationError like expected! %#v", err)
		return err
	}
	for idx, field := range validationErr.Fields {
		withoutContext := field.FieldName
		validationErr.Fields[idx].FieldName = fieldName + "." + withoutContext
	}
	return validationErr
}

// CombineErrors takes several errors and returns a single validation error.
func CombineErrors(errs ...error) error {
	fields := make([]InvalidField, 0, 5)
	for _, err := range errs {
		if err == nil {
			continue
		}
		validationErr, ok := err.(*ValidationError)
		if !ok {
			log.Printf("One of the errors passed to CombineErrors was not a ValidationError like expected! %#v", err)
			return err
		}
		for _, field := range validationErr.Fields {
			fields = append(fields, field)
		}
	}
	if len(fields) == 0 {
		return nil
	}
	return &ValidationError{
		BasicError: BasicError{
			Code:    "INVALID_DATA",
			Message: INVALID_DATA_MESSAGE,
		},
		Fields: fields,
	}
}
