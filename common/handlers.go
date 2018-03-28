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
	if err == ParseErr {
		WriteResponse(w, 400, err)
	} else if err == NotFoundErr {
		WriteResponse(w, 404, err)
	} else if IsValidationError(err) {
		WriteResponse(w, 422, err)
	} else if IsError(err) {
		WriteResponse(w, 500, err)
	} else {
		log.Println("Unexpected error: " + err.Error())
		WriteResponse(w, 500, map[string]string{"message": "Sorry, something went wrong on our end. Try again later!"})
	}
}
