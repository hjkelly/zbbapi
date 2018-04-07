package common

import (
	"encoding/json"
	"log"
	"net/http"
)

// WriteResponse preps some data to a response with the given status code.
func WriteResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			// TODO: Log error so we can avoid this in the future.
			panic(err)
		}
	}
}

// WriteErrorResponse preps the response by trying to guess the type of the error.
func WriteErrorResponse(w http.ResponseWriter, err error) {
	if result, ok := GetError(err); ok {
		// If it's our generic error, respond accordingly.
		if result == ParseErr {
			WriteResponse(w, 400, err)
		} else if result == NotFoundErr {
			WriteResponse(w, 404, err)
		} else {
			WriteResponse(w, 500, err)
		}
	} else if _, ok := GetValidationError(err); ok {
		// If it's a validation error, handle that.
		WriteResponse(w, 422, err)
	} else {
		// If it's just a generic error, respond and log an error.
		log.Println("Unexpected error: " + err.Error())
		WriteResponse(w, 500, map[string]string{"message": "Sorry, something went wrong on our end. Try again later!"})
	}
}
