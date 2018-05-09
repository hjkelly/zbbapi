package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hjkelly/zbbapi/common"
	"github.com/hjkelly/zbbapi/models"
	"github.com/hjkelly/zbbapi/services/plans"
	"github.com/julienschmidt/httprouter"
)

func listPlans(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	results, err := plans.List()
	if err != nil {
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 200, results)
}

func createPlan(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Parse the request body.
	var plan models.Plan
	err := json.NewDecoder(r.Body).Decode(&plan)
	if err != nil {
		log.Print(err.Error())
		common.WriteErrorResponse(w, common.ParseErr)
		return
	}
	// Save it.
	result, err := plans.Create(plan)
	if err != nil {
		// TODO: Handle validation vs. DB error...
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 201, result)
}

func retrievePlan(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	result, err := plans.Retrieve(params.ByName("id"))
	if err != nil {
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 200, result)
}

func updatePlan(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Parse the request body.
	var plan models.Plan
	err := json.NewDecoder(r.Body).Decode(&plan)
	if err != nil {
		log.Print(err.Error())
		common.WriteErrorResponse(w, common.ParseErr)
		return
	}
	// Update according to the URL.
	result, err := plans.UpdateID(params.ByName("id"), plan)
	if err != nil {
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 200, result)
}

func deletePlan(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := plans.Delete(params.ByName("id"))
	if err != nil {
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 204, nil)
}
