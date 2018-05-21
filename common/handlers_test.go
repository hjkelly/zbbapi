package common

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteErrorResponse(t *testing.T) {
	for _, testCase := range []struct {
		desc         string
		inputErr     error
		expectedCode int
		expectedBody map[string]interface{}
	}{
		{
			desc:         "common.BasicError",
			inputErr:     &BasicError{Code: "CODE", Message: "MESSAGE"},
			expectedCode: 500,
			expectedBody: map[string]interface{}{"message": "MESSAGE", "code": "CODE"},
		},
		{
			desc:         "common.ValidationError",
			inputErr:     NewValidationError("FIELDNAME", "FIELDCODE", "FIELDMESSAGE"),
			expectedCode: 422,
			expectedBody: map[string]interface{}{
				"message": "One or more fields was either missing or invalid.",
				"code":    "INVALID_DATA",
				"fields": []interface{}{
					map[string]interface{}{
						"fieldName": "FIELDNAME",
						"code":      "FIELDCODE",
						"message":   "FIELDMESSAGE",
					},
				},
			},
		},
		{
			desc:         "errorString (builtin)",
			inputErr:     errors.New("laksjd"),
			expectedCode: 500,
			expectedBody: map[string]interface{}{"message": "Sorry, something went wrong on our end. Try again later!"},
		},
		{
			desc:         "*json.UnmarshalTypeError",
			inputErr:     &json.UnmarshalTypeError{Field: "FIELDNAME", Value: "ACTUAL", Type: reflect.TypeOf("EXPECTED")},
			expectedCode: 422,
			expectedBody: map[string]interface{}{
				"message": "Got value of wrong type for FIELDNAME. Expected string, but got ACTUAL.",
				"code":    "WRONG_TYPE",
			},
		},
	} {

		recorder := httptest.NewRecorder()
		WriteErrorResponse(recorder, testCase.inputErr)
		actualCode, actualBody := getCodeAndData(recorder)
		assert.Equal(t, testCase.expectedCode, actualCode, "CASE: %s, status code mismatch", testCase.desc)
		assert.Equal(t, testCase.expectedBody, actualBody, "CASE: %s, response body mismatch", testCase.desc)
	}
}

func getCodeAndData(recorder *httptest.ResponseRecorder) (int, map[string]interface{}) {
	result := recorder.Result()
	data := map[string]interface{}{}
	err := json.NewDecoder(result.Body).Decode(&data)
	if err != nil {
		panic("Failed to read response body while testing")
	}
	return result.StatusCode, data
}
