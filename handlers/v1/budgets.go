package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hjkelly/zbbapi/common"
	"github.com/hjkelly/zbbapi/models"
	"github.com/hjkelly/zbbapi/services/budgets"
	"github.com/julienschmidt/httprouter"
)

func listBudgets(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	results, err := budgets.List()
	if err != nil {
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 200, results)
}

func createBudget(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Parse the request body.
	var budget models.Budget
	err := json.NewDecoder(r.Body).Decode(&budget)
	if err != nil {
		log.Print(err.Error())
		common.WriteErrorResponse(w, common.ParseErr)
		return
	}
	// Save it.
	result, err := budgets.Create(budget)
	if err != nil {
		// TODO: Handle validation vs. DB error...
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 201, result)
}

func retrieveBudget(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	result, err := budgets.Retrieve(params.ByName("id"))
	if err != nil {
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 200, result)
}

func updateBudget(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Parse the request body.
	var budget models.Budget
	err := json.NewDecoder(r.Body).Decode(&budget)
	if err != nil {
		log.Print(err.Error())
		common.WriteErrorResponse(w, common.ParseErr)
		return
	}
	// Update according to the URL.
	result, err := budgets.UpdateID(params.ByName("id"), budget)
	if err != nil {
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 200, result)
}

func deleteBudget(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := budgets.Delete(params.ByName("id"))
	if err != nil {
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 204, nil)
}
