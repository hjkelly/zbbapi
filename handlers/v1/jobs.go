package v1

import (
	"encoding/json"
	"net/http"

	"github.com/hjkelly/zbbapi/common"
	"github.com/hjkelly/zbbapi/models"
	"github.com/hjkelly/zbbapi/services/jobs"
	"github.com/julienschmidt/httprouter"
)

func listJobs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	results, err := jobs.List()
	if err != nil {
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 200, results)
}

func createJob(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Parse the request body.
	var job models.Job
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		common.WriteErrorResponse(w, common.ParseErr)
		return
	}
	// Save it.
	result, err := jobs.Create(job)
	if err != nil {
		// TODO: Handle validation vs. DB error...
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 201, result)
}

func retrieveJob(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	result, err := jobs.Retrieve(params.ByName("id"))
	if err != nil {
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 200, result)
}

func updateJob(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Parse the request body.
	var job models.Job
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		common.WriteErrorResponse(w, common.ParseErr)
		return
	}
	// Update according to the URL.
	result, err := jobs.UpdateID(params.ByName("id"), job)
	if err != nil {
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 200, result)
}

func deleteJob(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := jobs.Delete(params.ByName("id"))
	if err != nil {
		common.WriteErrorResponse(w, err)
		return
	}
	common.WriteResponse(w, 204, nil)
}
