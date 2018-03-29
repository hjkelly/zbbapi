package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hjkelly/zbbapi/common"
	"github.com/hjkelly/zbbapi/controllers/budgets"
	"github.com/hjkelly/zbbapi/models"
	"github.com/julienschmidt/httprouter"
)

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
