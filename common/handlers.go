package common

import (
	"encoding/json"
	"fmt"
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
	switch err.(type) {
	case BasicError:
		be := err.(BasicError)
		WriteResponse(w, be.ResponseCode(), be)
	case *BasicError:
		be := err.(*BasicError)
		WriteResponse(w, be.ResponseCode(), be)
	case ValidationError, *ValidationError:
		WriteResponse(w, 422, err)
	case *json.UnmarshalTypeError:
		ute := err.(*json.UnmarshalTypeError)
		WriteResponse(w, 422, map[string]string{
			"message": fmt.Sprintf("Got value of wrong type for %s. Expected %s, but got %s.", ute.Field, ute.Type, ute.Value),
			"code":    "WRONG_TYPE",
		})
	default:
		log.Println("Unexpected error: " + err.Error())
		WriteResponse(w, 500, map[string]string{"message": "Sorry, something went wrong on our end. Try again later!"})
	}
}
