package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hjkelly/zbbapi/common"
	"github.com/hjkelly/zbbapi/controllers/categories"
	"github.com/hjkelly/zbbapi/models"
	"github.com/julienschmidt/httprouter"
)

func listCategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	results, err := categories.List()
	if err != nil {
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 200, results)
}

func createCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Parse the request body.
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Printf(err.Error())
		common.WriteErrorResponse(w, common.ParseErr)
		return
	}
	// Save it.
	result, err := categories.Create(category)
	if err != nil {
		// TODO: Handle validation vs. DB error...
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 201, result)
}

func retrieveCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	result, err := categories.Retrieve(params.ByName("id"))
	if err != nil {
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 200, result)
}

func deleteCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := categories.Delete(params.ByName("id"))
	if err != nil {
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 204, nil)
}

func updateCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Parse the request body.
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Printf(err.Error())
		common.WriteErrorResponse(w, common.ParseErr)
		return
	}
	// Update according to the URL.
	result, err := categories.UpdateID(params.ByName("id"), category)
	if err != nil {
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 200, result)
}
